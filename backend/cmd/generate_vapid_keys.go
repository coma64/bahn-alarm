package cmd

import (
	"github.com/SherClockHolmes/webpush-go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var generateVapidKeysCmd = &cobra.Command{
	Use:   "generate-vapid-keys",
	Short: "Generate and print a vapid key pair",
	Run: func(cmd *cobra.Command, args []string) {
		private, public, err := webpush.GenerateVAPIDKeys()
		if err != nil {
			log.Err(err).Msg("failed to generate vapid keys")
		}

		log.Info().Str("privateKey", private).Str("publicKey", public).Msg("generated")
	},
}

func init() {
	rootCmd.AddCommand(generateVapidKeysCmd)
}
