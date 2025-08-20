package cmd

import (
	"fmt"
	"os"

	"github.com/mabushelbaia/gitfetch/internal/dashboard"
	"github.com/mabushelbaia/gitfetch/internal/github"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitfetch [username]",
	Short: "GitFetch CLI - fetch GitHub contributions",
	Long:  `GitFetch CLI fetches GitHub contributions and prints a Neofetch-style dashboard.`,
	Args:  cobra.ExactArgs(1), // Require exactly one positional argument
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]

		// Read the sample flag
		useSample, err := cmd.Flags().GetBool("sample")
		if err != nil {
			fmt.Printf("failed to read flag: %v\n", err)
			return
		}

		var user *github.UserInfo

		if useSample {
			// Load user from sample JSON
			user, err = github.LoadSampleUser()
			if err != nil {
				fmt.Printf("failed to load sample user: %v\n", err)
				return
			}
		} else {
			// Fetch user from GitHub
			user, err = github.FetchUserInfo(username)
			if err != nil {
				fmt.Printf("failed to fetch user info: %v\n", err)
				return
			}
		}

		dashboard.PrintDashboard(user)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("sample", "s", false, "Load sample user")
}
