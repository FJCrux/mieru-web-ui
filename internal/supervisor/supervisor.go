// SPDX-License-Identifier: GPL-3.0-or-later

// Package supervisor manages the mita daemon as a child process:
// spawn, restart with exponential backoff, log capture into a ring
// buffer, and SIGTERM propagation on shutdown.
package supervisor

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

const (
	backoffInitial = 1 * time.Second
	backoffMax     = 30 * time.Second
	// If the child stayed up this long, the next crash restarts the backoff.
	stableAfter = 60 * time.Second

	shutdownGrace = 10 * time.Second

	ringCapacity = 2000
)

// Supervisor runs "<binary> run" and keeps it alive.
type Supervisor struct {
	name   string
	binary string
	env    []string

	mu           sync.Mutex
	cmd          *exec.Cmd
	restartCount int
	startedAt    time.Time
	ring         *ringBuffer
	restartReq   chan struct{}
}

// New creates a supervisor that runs "<binary> run". name is used as the log
// prefix. extraEnv entries are appended to the current process environment.
func New(name, binary string, extraEnv []string) *Supervisor {
	return &Supervisor{
		name:       name,
		binary:     binary,
		env:        append(os.Environ(), extraEnv...),
		ring:       newRingBuffer(ringCapacity),
		restartReq: make(chan struct{}, 1),
	}
}

// Run supervises the child until ctx is cancelled. It returns after the
// child has been terminated.
func (s *Supervisor) Run(ctx context.Context) error {
	backoff := backoffInitial
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		start := time.Now()
		err := s.runOnce(ctx)
		if ctx.Err() != nil {
			return ctx.Err()
		}

		uptime := time.Since(start)
		if uptime >= stableAfter {
			backoff = backoffInitial
		}
		s.mu.Lock()
		s.restartCount++
		count := s.restartCount
		s.mu.Unlock()
		log.Printf("%s exited (%v) after %s; restart #%d in %s", s.name, err, uptime.Round(time.Second), count, backoff)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(backoff):
		}
		backoff = min(backoff*2, backoffMax)
	}
}

// runOnce starts the child and blocks until it exits or ctx is cancelled
// (in which case the child is terminated gracefully).
func (s *Supervisor) runOnce(ctx context.Context) error {
	cmd := exec.Command(s.binary, "run")
	cmd.Env = s.env

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start %s run: %w", s.binary, err)
	}
	s.mu.Lock()
	s.cmd = cmd
	s.startedAt = time.Now()
	s.mu.Unlock()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); s.capture(stdout) }()
	go func() { defer wg.Done(); s.capture(stderr) }()

	done := make(chan error, 1)
	go func() {
		wg.Wait()
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		s.terminate(cmd, done)
		return ctx.Err()
	case <-s.restartReq:
		s.terminate(cmd, done)
		return fmt.Errorf("restart requested")
	}
}

// terminate sends SIGTERM, waits up to shutdownGrace, then SIGKILLs.
func (s *Supervisor) terminate(cmd *exec.Cmd, done <-chan error) {
	if cmd.Process != nil {
		_ = cmd.Process.Signal(syscall.SIGTERM)
	}
	select {
	case <-done:
	case <-time.After(shutdownGrace):
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		<-done
	}
}

// Restart asks the supervisor to restart the child immediately
// (used by the panel's "restart mita" action). No-op if a restart
// is already pending.
func (s *Supervisor) Restart() {
	select {
	case s.restartReq <- struct{}{}:
	default:
	}
}

// capture copies child output lines into the ring buffer and mirrors
// them to the panel's stdout so they land in `docker logs`.
func (s *Supervisor) capture(r io.Reader) {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64*1024), 256*1024)
	for sc.Scan() {
		line := sc.Text()
		s.ring.append(line)
		fmt.Printf("[%s] %s\n", s.name, line)
	}
}

// Logs returns up to n most recent child log lines (all if n <= 0).
func (s *Supervisor) Logs(n int) []string {
	return s.ring.tail(n)
}

// Stats reports restart count and current child start time.
func (s *Supervisor) Stats() (restarts int, startedAt time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.restartCount, s.startedAt
}

// ringBuffer is a fixed-capacity line buffer.
type ringBuffer struct {
	mu    sync.Mutex
	lines []string
	next  int
	full  bool
}

func newRingBuffer(capacity int) *ringBuffer {
	return &ringBuffer{lines: make([]string, capacity)}
}

func (r *ringBuffer) append(line string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.lines[r.next] = line
	r.next = (r.next + 1) % len(r.lines)
	if r.next == 0 {
		r.full = true
	}
}

func (r *ringBuffer) tail(n int) []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	var out []string
	if r.full {
		out = append(out, r.lines[r.next:]...)
	}
	out = append(out, r.lines[:r.next]...)
	if n > 0 && len(out) > n {
		out = out[len(out)-n:]
	}
	return out
}
