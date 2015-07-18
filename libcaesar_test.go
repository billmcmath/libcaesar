package libcaesar

import (
	"testing"
)

func TestEncrptyedStringIsNotTheSameAsPlainTextString(t *testing.T) {
	plainText := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG, the quick brown fox jumps over the lazy dog."
	cipherText, err := Encrypt(3, plainText)
	decodedTest, err := Decrypt(3, cipherText)

	if err != nil {
		t.Fatalf("Got the error \"%s\"", err)
	}
	if cipherText == "" {
		t.Fatalf("The cipher text was blank")
	}
	if cipherText == plainText {
		t.Fatalf("The string has not been encrypted")
	}
	if plainText != decodedTest {
		t.Fatalf("The string has not been decrypted properly")
	}
}
