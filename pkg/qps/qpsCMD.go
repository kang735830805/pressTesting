package qps

import (
	"github.com/spf13/cobra"
)



var (
	txId string
)

func QpsCMD() *cobra.Command {

	qpsCmd := &cobra.Command {
		Use:   "qps",
		Short: "Pressing chainMaker key command",
		Long:  "Pressing chainMaker key command",
		RunE: func(_ *cobra.Command, _ []string) error {
			return RunQps()
		},
	}
	flags := qpsCmd.Flags()
	flags.StringVarP(&txId, "txId", "x", "", "合约参数")
	return qpsCmd
}