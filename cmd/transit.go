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

// transitCmd represents the transit command
var transitCmd = &cobra.Command{
	Use:   "transit",
	Short: "Create download link for sending a file",
	Long:  `transit: Setup gate to send file across the internet.  Configure with file, gate, and policy.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("transit requested without args")
		} else {
			fmt.Println("transit requested with args")
		}
	},
}

func init() {
	rootCmd.AddCommand(transitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	filename := transitCmd.PersistentFlags().String("filename", "", "path to file")
	gate := transitCmd.PersistentFlags().String("gate", "", "name your transit")
	auth := transitCmd.PersistentFlags().String("auth", "", "enforce policy on transit")
	fmt.Println(*filename)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
