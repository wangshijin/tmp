package main

import (
	"errors"
)

var (
	ErrKeyTooLong = errors.New("Key too long")
	ErrKeyIsEmpty = errors.New("Key is empty")
)

type Cipher struct {
	s    [256]byte
	i, j int
}

func NewCipher(key []byte) (*Cipher, error) {
	keyLen := len(key)
	if keyLen == 0 {
		return nil, ErrKeyIsEmpty
	}
	if keyLen > 256 {
		return nil, ErrKeyTooLong
	}

	cipher := new(Cipher)

	T := make([]byte, 256)
	for i := 0; i < 256; i++ {
		cipher.s[i] = byte(i)
		T[i] = key[i%keyLen]
	}

	j := 0
	for i := 0; i < 256; i++ {
		j = (j + int(cipher.s[i]+T[i])) & 0xFF
		cipher.s[i], cipher.s[j] = cipher.s[j], cipher.s[i]
	}

	cipher.i = 0
	cipher.j = 0

	return cipher, nil
}

func (c *Cipher) XORKeyStream(dst, src []byte) {
	if len(dst) < len(src) {
		return
	}

	for n := 0; n < len(src); n++ {
		c.i = (c.i + 1) & 0xFF
		c.j = (c.j + int(c.s[c.i])) & 0xFF
		c.s[c.i], c.s[c.j] = c.s[c.j], c.s[c.i]
		t := int(c.s[c.i]+c.s[c.j]) & 0xFF
		dst[n] = c.s[t] ^ src[n]
	}
}
