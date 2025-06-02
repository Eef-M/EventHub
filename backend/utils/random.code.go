package utils

import (
	"fmt"
	"math/rand"
)

func GenerateTicketCode() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 4)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	num := rand.Intn(10000)
	return fmt.Sprintf("%s-%04d", string(b), num)
}
