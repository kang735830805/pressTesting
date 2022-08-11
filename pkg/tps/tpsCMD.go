package tps

import "github.com/spf13/cobra"

var (
	name string
	method string
	parameter string
)

func TpsCMD() *cobra.Command {
	keyCmd := &cobra.Command{
		Use:   "tps",
		Short: "Pressing chainMaker key command",
		Long:  "Pressing chainMaker key command",
		RunE: func(_ *cobra.Command, _ []string) error {
			return RunTps()
		},
	}
	flags := keyCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "合约名称")
	flags.StringVarP(&method, "method", "m", "", "合约内的方法")
	flags.StringVarP(&parameter, "parameter", "p", "", "合约参数")
	return keyCmd
}

