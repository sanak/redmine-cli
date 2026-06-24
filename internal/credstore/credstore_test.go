package credstore

import (
	"errors"
	"testing"

	"github.com/zalando/go-keyring"
)

func TestStoreSetGetDelete(t *testing.T) {
	keyring.MockInit()

	s := New()
	if err := s.Set("work", "secret-key"); err != nil {
		t.Fatal(err)
	}
	got, err := s.Get("work")
	if err != nil {
		t.Fatal(err)
	}
	if got != "secret-key" {
		t.Fatalf("Get = %q, want %q", got, "secret-key")
	}
	if err := s.Delete("work"); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Get("work"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("after delete, Get err = %v, want ErrNotFound", err)
	}
}

func TestAvailableTrueWithMock(t *testing.T) {
	keyring.MockInit()
	if !Available() {
		t.Fatal("Available() = false with mock keyring, want true")
	}
}
