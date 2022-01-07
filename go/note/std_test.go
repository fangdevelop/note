package note

import (
	"errors"
	"testing"
)

func TestIsNotNegative(log *testing.T) {
	err := errors.New("Is Negative")
	if IsNotNegative(-1) {
		log.Log("OK")
	} else {
		log.Fatal(err)
	}
	if IsNotNegative(1) {
		log.Log("OK")
	} else {
		log.Error(err)
	}
}
