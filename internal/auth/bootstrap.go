// SPDX-License-Identifier: GPL-3.0-or-later

package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/fjcrux/mieru-web-ui/internal/store"
)

const passwordAlphabet = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// Bootstrap ensures a panel admin account exists. On first run it uses
// envUser/envPass when provided, otherwise generates credentials and
// prints them to stdout once (3x-ui style). Existing accounts are left
// untouched.
func Bootstrap(st store.Store, envUser, envPass string) error {
	_, _, err := st.Admin()
	if err == nil {
		return nil
	}
	if !errors.Is(err, store.ErrNotFound) {
		return err
	}

	user := envUser
	if user == "" {
		user = "admin"
	}
	pass := envPass
	generated := false
	if pass == "" {
		pass, err = randomPassword(16)
		if err != nil {
			return err
		}
		generated = true
	}
	hash, err := HashPassword(pass)
	if err != nil {
		return err
	}
	if err := st.SetAdmin(user, hash); err != nil {
		return err
	}
	if generated {
		fmt.Printf("\n==============================================\n")
		fmt.Printf(" Panel admin created\n")
		fmt.Printf("   username: %s\n", user)
		fmt.Printf("   password: %s\n", pass)
		fmt.Printf(" Change it in Settings after first login.\n")
		fmt.Printf("==============================================\n\n")
	}
	return nil
}

func randomPassword(n int) (string, error) {
	out := make([]byte, n)
	max := big.NewInt(int64(len(passwordAlphabet)))
	for i := range out {
		idx, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		out[i] = passwordAlphabet[idx.Int64()]
	}
	return string(out), nil
}
