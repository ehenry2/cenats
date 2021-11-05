/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	subject, natsURL, eventPayloadFile string
)

// pubCmd represents the pub command
var pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "Publish a message in CloudEvents format to NATS",
	Long: `Publish a message in CloudEvents format to a NATS message queue`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pub called")
	},
}

func init() {
	rootCmd.AddCommand(pubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pubCmd.Flags().StringVarP(&subject, "subject", "s", "",
		"NATS subject to publish to. In NATS, subjects scope messages into streams or topics. (required)")
	_ = pubCmd.MarkFlagRequired("subject")
	pubCmd.Flags().StringVarP(&natsURL, "url", "u", "nats://127.0.0.1:4222",
		"URL of the NATS server")
	pubCmd.Flags().StringVarP(&eventPayloadFile, "payloadFile", "f", "",
		"Path to json file containing the event payload (required)")
	_ = pubCmd.MarkFlagRequired("payloadFile")
}
