// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"regexp"
)

func getLocalGitURL() string {
	getRemoteURL := exec.Command("git", "remote", "get-url", "origin")
	remoteGitURL, err := getRemoteURL.CombinedOutput()

	if err != nil {
		log.Printf("Unable to get local git remote url: %v", err)
		return ""
	}

	findGit := regexp.MustCompile("git@github.com:")
	removedGit := findGit.ReplaceAllString(string(remoteGitURL), "http://github.com/")

	log.Printf("Found github url for local repo")
	log.Printf(removedGit)

	return removedGit
}

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Opens up a default browser with an url to github repo - in directory.",
	Long: `It uses xdg-open to use default browser.
	
	A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Opening github repo url in the browser")

		url := getLocalGitURL()

		if len(url) > 0 {
			openBrowserCommand := exec.Command("xdg-open", url)
			err := openBrowserCommand.Run()

			if err != nil {
				log.Printf("Command finished with error: %v", err)
			}
			return
		}

		log.Printf("Github remote url empty cannot open browser")
		return
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// githubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// githubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
