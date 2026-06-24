// Package credstore stores Redmine credentials in the OS keychain
// (macOS Keychain, Windows Credential Manager, Linux Secret Service)
// using the profile name as the account key.
package credstore

import "github.com/zalando/go-keyring"

// Service is the keyring service name under which secrets are stored.
const Service = "redmine-cli"

// ErrNotFound is returned by Get when no secret exists for the profile.
var ErrNotFound = keyring.ErrNotFound

// Store reads and writes per-profile secrets.
type Store interface {
	Get(profile string) (string, error)
	Set(profile, secret string) error
	Delete(profile string) error
}

type keyringStore struct{}

// New returns the OS keyring-backed store.
func New() Store { return keyringStore{} }

func (keyringStore) Get(profile string) (string, error) {
	return keyring.Get(Service, profile)
}

func (keyringStore) Set(profile, secret string) error {
	return keyring.Set(Service, profile, secret)
}

func (keyringStore) Delete(profile string) error {
	return keyring.Delete(Service, profile)
}

// Available reports whether the OS keyring can be used in this environment
// (false on headless systems without a Secret Service, etc.).
func Available() bool {
	const probe = "__redmine_cli_probe__"
	if err := keyring.Set(Service, probe, "1"); err != nil {
		return false
	}
	_ = keyring.Delete(Service, probe)
	return true
}
