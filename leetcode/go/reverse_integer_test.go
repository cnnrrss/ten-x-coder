package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	got := reverse2(123)
	if got != 321 {
		t.Errorf("reverse(123) = %d; want 321", got)
	}

	got = reverse2(-321)
	if got != -123 {
		t.Errorf("reverse(-321) = %d; want -123", got)
	}

	// Overflow example
	got = reverse2(210)
	if got != 12 {
		t.Errorf("reverse(210) = %d; want 12", got)
	}

	got = reverse2(-2222)
	if got != -2222 {
		t.Errorf("reverse(-2222) = %d; want -2222", got)
	}
}
