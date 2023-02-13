package generation_test

import (
	"passgen/pkg/generation"
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {

	p := generation.GeneratePassword(12)
	if len(p) != 12 {
		t.Errorf("the length of the password generated (%v) isn't equal to the value passed as argument (12)", len(p))
	}

	if !strings.ContainsAny(p, "0123456789") {
		t.Errorf("expected password to contain numbers, but got %s", p)
	}

	if !strings.ContainsAny(p, "~=+%^*/()[]{}/!@#$?|") {
		t.Errorf("expected password to contain symbols, but got %s", p)
	}

	if !strings.ContainsAny(p, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz") {
		t.Errorf("expected password to contain letters, but got %s", p)
	}
}
