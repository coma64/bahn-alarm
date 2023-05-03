package main

import (
	"github.com/SherClockHolmes/webpush-go"
	"github.com/rs/zerolog/log"
)

func main() {
	private, public, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		log.Err(err).Msg("failed to generate vapid keys")
	}

	log.Info().Str("privateKey", private).Str("publicKey", public).Msg("generated")
}
