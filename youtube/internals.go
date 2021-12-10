package youtube

import (
	"math/rand"
	"time"
)

func ComputeSignature(s string) string {
	chars := []byte(s)
	// Reverse string a.reverse()
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	// remove first item a.splice(0,b) where b is 1
	chars = chars[1:]
	// swap first character and character at b%a.length where b is 30
	c := chars[0]
	chars[0] = chars[30%len(chars)]
	chars[30%len(chars)] = c

	return string(chars)
}

func GenerateClientPlaybackNonce() string {
	rand.Seed(time.Now().Unix())

	possibleChars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_")
	var cpn []byte
	for i := 0; i < 16; i++ {
		cpn = append(cpn, possibleChars[rand.Intn(256)&63])
	}
	return string(cpn)
}
