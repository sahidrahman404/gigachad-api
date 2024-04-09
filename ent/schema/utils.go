package schema

import (
	"log"
	"time"

	"github.com/segmentio/ksuid"
)

func generateKSUID() string {
	k, err := ksuid.NewRandomWithTime(time.Now())
	if err != nil {
		log.Fatal(err)
	}
	return k.String()
}

func generateTime() time.Time {
	return time.Now().UTC()
}
