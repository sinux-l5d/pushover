/*
Copyright © 2022 Simon LEONARD git-1001af4@sinux.sh

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pushover",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		checkFlags(cmd, args)
		sendPayload(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("token", "t", "", "pushover api token")
	rootCmd.Flags().StringP("user", "u", "", "pushover user key")
	rootCmd.Flags().StringP("message", "m", "", "message to send")
	rootCmd.Flags().StringP("title", "T", "", "title of message")
}

func checkFlags(cmd *cobra.Command, args []string) {
	// check that all flags are set to non-empty values with for loop
	stop := false
	for _, flag := range []string{"token", "user", "message", "title"} {
		if cmd.Flag(flag).Value.String() == "" {
			fmt.Println("flag", flag, "is not set")
			stop = true
		}
	}
	if stop {
		os.Exit(1)
	}
}

func sendPayload(cmd *cobra.Command, args []string) {
	// post form payload to pushover using http library
	resp, err := http.PostForm("https://api.pushover.net/1/messages.json",
		url.Values{
			"token":   {cmd.Flag("token").Value.String()},
			"user":    {cmd.Flag("user").Value.String()},
			"message": {cmd.Flag("message").Value.String()},
			"title":   {cmd.Flag("title").Value.String()},
		},
	)
	// print response
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
