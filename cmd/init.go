package cmd

import (
	"GPM/internal/packagefile"
	"GPM/internal/util"
	"fmt"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a package_file.gpm file and a packages directory (if one does not already exist)",
	Long: `The "init" command sets up the necessary files for GPM.

It will:
  - Prompt for basic package details.
  - Create a package_file.gpm with the provided values.
  - Ensure a packages/ directory exists.`,
	Run: func(cmd *cobra.Command, args []string) {
		defaultPackageName := util.CurrentDirectory()

		packageName := util.GetUserInput(fmt.Sprintf("Enter package name (default is %s): ", defaultPackageName))
		if packageName == "" {
			packageName = defaultPackageName
		}

		defaultPackageURL := fmt.Sprintf("https://github.com/user/%s", packageName)
		// TODO Store and fetch user config data
		packageURL := util.GetUserInput(fmt.Sprintf("Enter package url (default is %s): ", defaultPackageURL))
		if packageURL == "" {
			packageURL = defaultPackageURL
		}

		// Create package_file.gpm with user input
		err := packagefile.CreatePackageFile(packageName, packageURL)
		if err != nil {
			fmt.Println("⚠ Error creating package file:", err)
		} else {
			fmt.Println("✔ Created package_file.gpm")
		}

		// Ensure packages/ directory exists
		if err := packagefile.EnsurePackagesDir(); err != nil {
			fmt.Println("⚠", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
