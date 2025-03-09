package cmd

import (
	"fmt"
	"os"

	"github.com/Gabriel-Schiestl/archgen/internal/architectures"
	"github.com/spf13/cobra"
)

var arch string
var dir string

var generateCmd = &cobra.Command{
	Use: "generate <arch>",
	Short: "Generates an architecture structure",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you need to specify an architecture:\n\nclean\nmvc")
		}

		arch = args[0]

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch arch {
		case "clean":
			architectures.Clean(dir)
		case "mvc":
			architectures.MVC(dir)
		default:
			fmt.Println("Arch does not satisfies:\n\nclean\nmvc")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	rootCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "", "Defines the dir to create the structure")
}