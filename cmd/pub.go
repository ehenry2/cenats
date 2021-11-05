package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ehenry2/cenats/internal/pub"
	"github.com/spf13/cobra"
)

var (
	eventPayloadFile, eventSource, eventType string
)

// pubCmd represents the pub command
var pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "Publish a message in CloudEvents format to NATS",
	Long: `Publish a message in CloudEvents format to a NATS message queue`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pub called")
		s, err := pub.NewNatsSender(subject, natsURL)
		if err != nil {
			log.Fatalf("failed to initialize nats client: %s\n", err)
		}
		f, err := os.Open(eventPayloadFile)
		if err != nil {
			log.Fatalf("failed to open payload file: %s\n", err)
		}
		h, err := pub.NewHandler(s, f)
		if err != nil {
			log.Fatalf("failed to initialize cloud events publisher: %s\n", err)
		}
		ctx := context.Background()
		if err = h.Handle(ctx, eventType, eventSource); err != nil {
			log.Fatalf("failed to publish event: %s\n", err)
		}
		log.Println("message published successfully")
	},
}

func init() {
	rootCmd.AddCommand(pubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	pubCmd.Flags().StringVarP(&eventPayloadFile, "payloadFile", "f", "",
		"Path to json file containing the event payload (required)")
	_ = pubCmd.MarkFlagRequired("payloadFile")
	pubCmd.Flags().StringVarP(&eventType, "eventType", "t", "custom.default_event",
		"CloudEvents event type of your event")
	pubCmd.Flags().StringVarP(&eventSource, "eventSource", "", "mysource",
		"CloudEvents event source")
}
