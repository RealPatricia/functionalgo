package functionalgo

import (
	"testing"
)

func TestSanity(t *testing.T) {
	if 2+2 != 4 {
		t.Errorf("Sorry, math's closed, Moose out front should have told ya")
	}
}
