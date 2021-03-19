package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kitko/commitizen-go/git"
	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install this tool to git-core as git-cz",
	Run: func(cmd *cobra.Command, args []string) {
		appFilePath, _ := exec.LookPath(os.Args[0])
		if path, err := git.InstallSubCmd(appFilePath, "cz"); err != nil {
			fmt.Printf("Install commitizen failed, err=%v\n", err)
		} else {
			fmt.Printf("Install commitizen to %s\n", path)
		}
	},
}

var InstallHookCmd = &cobra.Command{
	Use:   "install-hook",
	Short: "Install this tool to ./git/hooks/prepare-commit-msg",
	Run: func(cmd *cobra.Command, args []string) {
		if err := git.InstallHookCmd(); err != nil {
			fmt.Printf("Install commitizen as hook failed, err=%v\n", err)
		} else {
			fmt.Printf("Successful install commitizen as git hook in .git/hoooks/prepare-commit-msg\n")
		}
	},
}
