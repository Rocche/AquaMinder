package notification

import (
	"math/rand"

	"github.com/gen2brain/beeep"
)

func getPhrase() string {
	return phrases[rand.Intn(len(phrases))]
}

func Notify() error {
	return beeep.Notify("AquaMinder", getPhrase(), "")
}
