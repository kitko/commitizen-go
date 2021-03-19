package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kitko/commitizen-go/commit"
	"github.com/kitko/commitizen-go/git"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() error {
	return rootCmd.Execute()
}

var (
	all     bool
	hook    bool
	debug   bool
	rootCmd = &cobra.Command{
		Use:          "commitizen-go",
		Long:         `Command line utility to standardize git commit messages, golang version.`,
		Run:          RootCmd,
		SilenceUsage: true,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolVarP(&all, "all", "a", false, "tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected")
	rootCmd.Flags().BoolVarP(&hook, "hook", "", false, "tell the command to save message to .git/COMMIT_EDITMSG file")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug mode, output debug info to debug.log")

	// viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(InstallCmd)
	rootCmd.AddCommand(InstallHookCmd)
}

func initConfig() {
	if !debug {
		log.SetOutput(ioutil.Discard)
	} else {
		f, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		// defer f.Close()
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.SetOutput(f)
	}

	// Find home directory.
	if home, err := homedir.Dir(); err != nil {
		log.Printf("Get home dir failed, err=%v\n", err)
		os.Exit(1)
	} else {
		// Search config in home directory with name ".git-czrc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".git-czrc")
		viper.SetConfigType("json")

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
				log.Printf("can not find config file\n")
			} else {
				// Config file was found but another error was produced
				log.Printf("read config failed, err=%v\n", err)
			}
		} else {
			log.Printf("read config success\n")
		}
	}
}

func RootCmd(command *cobra.Command, args []string) {
	// exit if not git repo
	if _, err := git.IsCurrentDirectoryGitRepo(); err != nil {
		fmt.Print(err)
		return
	}

	if message, err := commit.FillOutForm(); err == nil {

		if hook {
			err := git.PrepareCommitMessage(message)
			if err != nil {
				log.Printf("run pre commit hook failed, err=%v\n", err)
			}
			return
		}

		// do git commit
		result, err := git.CommitMessage(message, all)
		if err != nil {
			log.Printf("run git commit failed, err=%v\n", err)
			log.Printf("commit message is: \n\n%s\n\n", string(message))
		}
		fmt.Print(result)
	}
}
