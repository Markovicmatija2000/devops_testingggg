package utils

import (
	"strconv"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	// Run many times since GenerateCode is random - a single call could get
	// lucky and hide a formatting bug (e.g. missing zero-padding for small numbers).
	for i := 0; i < 1000; i++ {
		code := GenerateCode()

		if len(code) != 6 {
			t.Fatalf("GenerateCode() = %q, want length 6, got length %d", code, len(code))
		}

		n, err := strconv.Atoi(code)
		if err != nil {
			t.Fatalf("GenerateCode() = %q, want a numeric string: %v", code, err)
		}

		if n < 0 || n > 999999 {
			t.Fatalf("GenerateCode() = %q, want value in [0, 999999]", code)
		}
	}
}
