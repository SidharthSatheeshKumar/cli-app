/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-app/app"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// compareCmd represents the compare command
var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare the responses of the two endpoints",
	Long: `compare command is used to compare the JSON responses of two endpoints as given as the arguments.
	At first, obtain the JSON response of each endpoint in string format
	The two responses were converted to JSON encoded format and compared those formats and the difference will be returned as a JSON response`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiRespOriginal, err := app.GetApiResponse(args[0])
		if err != nil {
			log.Fatalf("api call issue:%s  | endpoint: %s", err.Error(), args[0])
		}
		apiRespComp, err := app.GetApiResponse(args[1])
		if err != nil {
			log.Fatalf("api call issue:%s | endpoint: %s", err.Error(), args[1])
		}

		responseSummary, err := app.ResponseCheck(apiRespOriginal, apiRespComp)
		if err != nil {
			log.Fatalf("checking issue: %s", err.Error())
		}

		if responseSummary != (app.ApiResponseDifference{}) {
			byteResp, err := json.Marshal(responseSummary)
			if err != nil {
				log.Fatalf("encode issue: %s", err.Error())
			}
			fmt.Println("seems there a difference: ", string(byteResp))
		} else {
			fmt.Println("provided endoints have same response")
		}
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
