package function

import (
	"testing"
)

func TestHandlerBadSSL(t *testing.T) {
	x := Handle([]byte("expired.badssl.com:443"))
	if x != "EXPIRED" {
		t.Errorf("Testing against expired.badssl.com, expected EXPIRED got %s", x)
	}
}

func TestHandlerGoodSSL(t *testing.T) {
	x := Handle([]byte("badssl.com:443"))
	if x != "OK" {
		t.Errorf("Testing against badssl.com, expected GOOD got %s", x)
	}
}
