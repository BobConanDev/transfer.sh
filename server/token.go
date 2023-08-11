/*
The MIT License (MIT)

Copyright (c) 2020- Andrea Spacca and Stefan Benten.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package server

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

var seed *rand.Rand

func init() {
	var seedBytes [8]byte
	if _, err := cryptoRand.Read(seedBytes[:]); err != nil {
		panic("cannot obtain cryptographically secure seed")
	}

	seed = rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(seedBytes[:]))))
}

const (
	// SYMBOLS characters used for short-urls
	SYMBOLS = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// generate a token
func token(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		x := seed.Intn(len(SYMBOLS) - 1)
		result = string(SYMBOLS[x]) + result
	}

	return result
}
