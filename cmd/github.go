package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"regexp"
)

func getLocalGitURL() (string, error) {
	getRemoteURL := exec.Command("git", "remote", "get-url", "origin")
	remoteGitURL, err := getRemoteURL.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("Unable to get local git remote url: %v", err)
	}

	findGit := regexp.MustCompile("git@github.com:")
	removedGit := findGit.ReplaceAllString(string(remoteGitURL), "http://github.com/")

	log.Printf("Found github url for local repo %s", removedGit)

	return removedGit, nil
}

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Opens up a default browser with master remote url for github.",
	Run: func(_ *cobra.Command, _ []string) {
		log.Printf("Opening github repo url in the browser")

		url, err := getLocalGitURL()

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
	rootCmd.AddCommand(githubCmd)
}
