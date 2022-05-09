package youtube

import (
	"math/rand"
	"time"
)

// The following is the youtube minified function for computing the signature as per 9th of May 2022
// Sourced from https://www.youtube.com/s/player/a4d8b401/player_ias.vflset/en_US/base.js
/*
var Ex = {
        mg: function(a) {
            a.reverse()
        },
        sI: function(a, b) {
            var c = a[0];
            a[0] = a[b % a.length];
            a[b % a.length] = c
        },
        Wm: function(a, b) {
            a.splice(0, b)
        }
    };
 tra = function(a) {
        a = a.split("");
        Ex.Wm(a, 2);
        Ex.mg(a, 10);
        Ex.Wm(a, 2);
        Ex.mg(a, 69);
        Ex.Wm(a, 3);
        Ex.mg(a, 59);
        Ex.Wm(a, 3);
        Ex.mg(a, 39);
        return a.join("")
    }
*/
func computeSignature(s string) string {
	// a = a.split("");
	a := []byte(s)
	// Ex.Wm(a, 2); a.splice(0, b)
	a = a[2:]
	// Ex.mg(a, 10); a.reverse()
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	// Ex.Wm(a, 2); a.splice(0, b)
	a = a[2:]
	// Ex.mg(a, 69); a.reverse()
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	// Ex.Wm(a, 3); a.splice(0, b)
	a = a[3:]
	// Ex.mg(a, 59); a.reverse()
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	// Ex.Wm(a, 3); a.splice(0, b)
	a = a[3:]
	// Ex.mg(a, 39); a.reverse()
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	// return a.join("")
	return string(a)
}

// this method still seems to work but looking at base.js it has been replaced by a newer version
func computeSignatureOld(s string) string {
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
