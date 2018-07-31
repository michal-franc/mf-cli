// Copyright © 2018 Michal Franc
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
	Short: "Opens up a default browser with master remote url for github.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Opening github repo url in the browser")

		url := getLocalGitURL()

		if len(url) <= 0 {
			log.Printf("Github remote url empty cannot open browser")
			return
		}

		openBrowserCommand := exec.Command("xdg-open", url)
		err := openBrowserCommand.Run()

		if err != nil {
			log.Printf("Command finished with error: %v", err)
		}

		return

	},
}

func init() {
	rootCmd.AddCommand(githubCmd)
}
