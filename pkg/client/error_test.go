package client

import (
	"fmt"
	"testing"
)

func TestNewTypeNotFoundError(t *testing.T) {
	msg := "Testing Not Found Error"
	err := NewTypeNotFoundError(msg)
	if err.message != msg {
		t.Error(err)
	}
}

func TestError(t *testing.T) {
	msg := "Testing Error()"
	err := NewTypeNotFoundError(msg)
	if err.Error() != msg {
		t.Error(err)
	}
}

func TestIsResourceTypeNotFound(t *testing.T) {
	msg := "Testing IsResourceTypeNotFound()"
	err := NewTypeNotFoundError(msg)
	if !IsResourceTypeNotFound(err) {
		t.Error(err)
	}
	errwrap := fmt.Errorf("Testing Wrap Error: %w", err)
	if !IsResourceTypeNotFound(errwrap) {
		t.Error(errwrap)
	}
	errf := fmt.Errorf("Testing Not ResourceTypeNotFound")
	if IsResourceTypeNotFound(errf) {
		t.Error(err)
	}
}
