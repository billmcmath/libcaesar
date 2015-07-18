package libcaesar

import (
	"errors"
)

func Encrypt(key int8, plainText string) (string, error) {
	err := validateKey(key)
	if err != nil {
		return "", err
	}

	var cipherText string = ""

	letters := []byte(plainText)

	for _, letter := range letters {
		if isBasicLatinLetter(letter) {
			cipherText += shiftLetter(key, letter)
		} else {
			cipherText += string(letter)
		}
	}

	return cipherText, nil
}

func shiftLetter(key int8, letter byte) string {
	if letter >= 97 {
		iLetter := (int8(letter) + key - 97)
		if iLetter < 97 {
			iLetter += 26
		}
		iLetter = (iLetter % 26) + 97
	} else {
		iLetter := (int8(letter) + key - 65)
		if iLetter < 65 {
			iLetter += 26
		}
		iLetter = (iLetter % 26) + 65
		letter = byte(iLetter)
	}
	return string(letter)
}

func isBasicLatinLetter(letter byte) bool {
	if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
		return true
	} else {
		return false
	}
}

func validateKey(key int8) error {
	var err error = nil

	if key > 26 || key < -26 {
		err = errors.New("Invalid Key - Key must be in the range ±26")
	}

	return err
}

func Decrypt(key int8, cipherText string) (string, error) {
	err := validateKey(key)
	if err != nil {
		return "", err
	} else {
		key *= -1
	}
	return Encrypt(key, cipherText)
}
