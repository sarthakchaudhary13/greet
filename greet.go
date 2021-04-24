package greet

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello will genrate random greet for the
// name and a `nil`. If no name is provided the it will
// return an emty string with an error.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init will seed the rand package.
func init() {
	rand.Seed(time.Now().UnixNano())
}

//Return Random geet format
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Aur, %v. Kya haal h bhai!",
		"Aur Miya, %v kya haal hai?",
	}

	return formats[rand.Intn(len(formats))]
}
