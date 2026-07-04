// SPDX-License-Identifier: GPL-3.0-or-later

package supervisor

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// fakeMita writes a shell script that prints a line and sleeps forever.
func fakeMita(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "fake-mita")
	script := "#!/bin/sh\necho \"fake mita started pid $$\"\ntrap 'exit 0' TERM\nwhile true; do sleep 0.1; done\n"
	if err := os.WriteFile(path, []byte(script), 0o755); err != nil {
		t.Fatal(err)
	}
	return path
}

func waitFor(t *testing.T, timeout time.Duration, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatal("condition not met in time")
}

func TestCapturesLogsAndRestarts(t *testing.T) {
	s := New("test", fakeMita(t), nil)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go s.Run(ctx)

	waitFor(t, 5*time.Second, func() bool {
		return len(s.Logs(0)) >= 1
	})
	if got := s.Logs(0)[0]; !strings.Contains(got, "fake mita started") {
		t.Fatalf("unexpected log line: %q", got)
	}

	// Trigger a restart and expect a second start line from a new pid.
	s.Restart()
	waitFor(t, 10*time.Second, func() bool {
		restarts, _ := s.Stats()
		return restarts >= 1 && len(s.Logs(0)) >= 2
	})
}

func TestGracefulShutdown(t *testing.T) {
	s := New("test", fakeMita(t), nil)
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- s.Run(ctx) }()

	waitFor(t, 5*time.Second, func() bool { return len(s.Logs(0)) >= 1 })
	cancel()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("supervisor did not stop after cancel")
	}
}

func TestRingBufferWraps(t *testing.T) {
	r := newRingBuffer(3)
	for _, l := range []string{"a", "b", "c", "d", "e"} {
		r.append(l)
	}
	got := r.tail(0)
	want := []string{"c", "d", "e"}
	if len(got) != len(want) {
		t.Fatalf("got %v want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v want %v", got, want)
		}
	}
	if tail := r.tail(2); len(tail) != 2 || tail[1] != "e" {
		t.Fatalf("tail(2) = %v", tail)
	}
}
