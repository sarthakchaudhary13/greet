package greet

import (
	"regexp"
	"testing"
)

// TestHelloName calls greet.Hello with a name
// to check for a valid return value.
func TestHello(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greet.Hello with empty
// string
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("Gladys" = %q, %v, want "" , error`, msg, err)
	}
}
