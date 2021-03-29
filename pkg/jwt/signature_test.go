package jwt

import (
	"os"
	"testing"
)

func TestSignatureReturnProperHash(t *testing.T)  {
	_ = os.Setenv("JWT_SECRET", "secret")
	result := signature("abcdef", "12345")
	expected := "vjL2HxuvHXBxxfSye1_Mya7a8PpgrrwcYJMvpagoDTIBr6_SBe_eHkp4YSdmgADS6uqM6gPUtnv9auGnhBZz_A"

	if result != expected {
		t.Errorf("signature %v does not match %v", result, expected)
	}
}