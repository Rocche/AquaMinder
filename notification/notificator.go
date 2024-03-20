package notification

import (
	"math/rand"

	"github.com/gen2brain/beeep"
)

func getPhrase() string {
	return phrases[rand.Intn(len(phrases))]
}

// Notify sends a notification with a random catchy phrase about the
// importance of drinking water at regular intervals
func Notify() error {
	return beeep.Notify("AquaMinder", getPhrase(), "")
}
