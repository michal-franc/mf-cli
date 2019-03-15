package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func getTravisUrlFromGitRepo() (string, error) {
	getRemoteURL := exec.Command("git", "remote", "get-url", "origin")
	remoteGitURL, err := getRemoteURL.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("Unable to get local git remote url: %v", err)
	}

	findGit := regexp.MustCompile("git@github.com:")
	removedGit := findGit.ReplaceAllString(string(remoteGitURL), "https://travis-ci.com/")

	removedGit = strings.TrimSuffix(removedGit, ".git\n")

	log.Printf("Found travis url for local repo %s", removedGit)

	return removedGit, nil
}

// githubCmd represents the github command
var travisCommand = &cobra.Command{
	Use:   "travis",
	Short: "Opens up a default browser with travis url to git repo.",
	Run: func(_ *cobra.Command, _ []string) {
		log.Printf("Opening travis job url in the browser")

		url, err := getTravisUrlFromGitRepo()

		if err != nil {
			log.Printf("Github remote url empty cannot open browser")
			return
		}

		openBrowserCommand := exec.Command("xdg-open", url)

		if err := openBrowserCommand.Run(); err != nil {
			log.Fatalf("Command finished with error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(travisCommand)
}
