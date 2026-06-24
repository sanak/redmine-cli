package auth

import (
	"testing"

	"github.com/aarondpn/redmine-cli/v2/internal/config"
	"github.com/aarondpn/redmine-cli/v2/internal/credstore"
	"github.com/zalando/go-keyring"
)

func TestPersistCredentialsKeyring(t *testing.T) {
	keyring.MockInit()
	cfg := &config.Config{Server: "https://x", AuthMethod: "apikey"}

	if err := persistCredentials("work", cfg, "plain-key", credstore.New(), true); err != nil {
		t.Fatal(err)
	}
	if cfg.CredentialStore != "keyring" {
		t.Fatalf("CredentialStore = %q, want keyring", cfg.CredentialStore)
	}
	if cfg.APIKey != "" {
		t.Fatalf("APIKey should be empty when stored in keyring, got %q", cfg.APIKey)
	}
	got, err := credstore.New().Get("work")
	if err != nil {
		t.Fatal(err)
	}
	if got != "plain-key" {
		t.Fatalf("keyring secret = %q, want plain-key", got)
	}
}

func TestPersistCredentialsFileFallback(t *testing.T) {
	cfg := &config.Config{Server: "https://x", AuthMethod: "apikey"}
	if err := persistCredentials("work", cfg, "plain-key", credstore.New(), false); err != nil {
		t.Fatal(err)
	}
	if cfg.CredentialStore != "file" {
		t.Fatalf("CredentialStore = %q, want file", cfg.CredentialStore)
	}
	if cfg.APIKey != "plain-key" {
		t.Fatalf("APIKey = %q, want plain-key in file mode", cfg.APIKey)
	}
}
