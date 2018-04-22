package debug

import (
	"testing"
)

func TestParseEnabledPrefixString(t *testing.T) {
	enabledPrefixesString := "github.com/invokit/go-util golang.org"

	r , err := parseEnabledPrefixString(enabledPrefixesString)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if l := len(r); l != 2 {
		t.Errorf("expected 2 got %d. r is: %s", l, r)
	}

	if !r[0].MatchString("github.com/invokit/go-util") {
		t.Errorf("expected match")
	}

	if !r[0].MatchString("github.com/invokit/go-util/debug") {
		t.Errorf("expected match")
	}

	if r[0].MatchString("github.com/invokit") {
		t.Errorf("expected no match")
	}

	if r[0].MatchString("invokit/go-util") {
		t.Errorf("expected no match")
	}
}