package memory

import (
	"bytes"
	"testing"
)

func TestLocking(t *testing.T) {
	// Declare two slices to test on.
	dataOne := []byte("yellow submarine")
	dataTwo := []byte("yellow submarine")

	// Lock them.
	Protect(dataOne)
	Protect(dataTwo)

	// Check if they're zeroed out. They shouldn't be.
	if bytes.Equal(dataOne, make([]byte, 16)) || bytes.Equal(dataOne, make([]byte, 16)) {
		t.Error("Ctitical error: memory zeroed out early")
	}

	// Cleanup.
	Cleanup()

	// Check if data is zeroed out.
	for _, v := range dataOne {
		if v != 0 {
			t.Error("Didn't zero out memory; dataOne =", dataOne)
		}
	}
	for _, v := range dataTwo {
		if v != 0 {
			t.Error("Didn't zero out memory; dataTwo =", dataTwo)
		}
	}
}
