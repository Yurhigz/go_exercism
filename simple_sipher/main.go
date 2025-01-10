package main

import (
	"fmt"
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// Define the shift and vigenere types here.
type shift struct {
	distance int
}

type vigenere struct {
	key string
}

// Both types should satisfy the Cipher interface.

func isLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func NewCaesar() Cipher {
	return shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance <= 25 && distance != 0 && distance >= -25 {
		return shift{distance: distance}
	}
	return nil
}

func (c shift) Encode(input string) string {
	encodedInput := ""
	for _, letter := range strings.ToLower(input) {
		if isLetter(letter) {
			shift := int32(c.distance)
			if c.distance > 0 {
				encodedInput += string('a' + ((letter - 'a' + shift) % 26))
			} else {
				encodedInput += string('a' + ((letter - 'a' + shift + 26) % 26))
			}
		}
	}
	return encodedInput
}

func (c shift) Decode(input string) string {
	decodedInput := ""
	for _, letter := range strings.ToLower(input) {
		if isLetter(letter) {
			shift := int32(-c.distance)
			if c.distance < 0 {
				decodedInput += string('a' + ((letter - 'a' + shift) % 26))
			} else {
				decodedInput += string('a' + ((letter - 'a' + shift + 26) % 26))
			}
		}
	}
	return decodedInput
}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	countA := 0
	for _, letter := range key {
		if letter < 'a' || letter > 'z' {
			return nil
		}
		if letter == 'a' {
			countA += 1
		}
	}
	if countA == len(key) {
		return nil
	}
	return vigenere{key}
}

func (v vigenere) Encode(input string) string {
	encodedInput := ""
	i := 0
	for _, letter := range strings.ToLower(input) {
		if isLetter(letter) {
			shift := int32(v.key[i%len(v.key)] - 'a')
			encodedInput += string('a' + ((letter-'a')+shift)%26)
			i++
		}
	}
	return encodedInput
}

func (v vigenere) Decode(input string) string {
	decodedInput := ""
	i := 0
	for _, letter := range strings.ToLower(input) {
		if isLetter(letter) {
			shift := int32(v.key[i%len(v.key)] - 'a')
			decodedInput += string('a' + (letter-'a'-shift+26)%26)
			i++
		}
	}
	return decodedInput
}

func main() {
	v := NewVigenere("lemon")
	fmt.Println(v.Encode("ATTACKATDAWN"))
}
