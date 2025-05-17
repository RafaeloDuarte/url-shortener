package controllers

import (
	"testing"
)

func TestGenerateShortCode(t *testing.T) {
	code := generateShortCode()

	if len(code) != shortCodeLength {
		t.Errorf("Esperado %d caracteres, mas veio %d", shortCodeLength, len(code))
	}
}
