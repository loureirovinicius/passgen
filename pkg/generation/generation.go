package generation

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GeneratePassword(length uint) string {
	numbers := "0123456789"
	symbols := "~=+%^*/()[]{}/!@#$?|"
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	all := letters + numbers + symbols

	passwordSlice := make([]byte, length)
	for i := 0; uint(i) < length; i++ {
		passwordSlice[i] = all[rand.Intn(len(all))]
	}

	rand.Shuffle(int(length), func(i, j int) {
		passwordSlice[i], passwordSlice[j] = passwordSlice[j], passwordSlice[i]
	})

	password := string(passwordSlice)
	return password
}
