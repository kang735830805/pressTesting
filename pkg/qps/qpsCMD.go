package qps

import "github.com/spf13/cobra"



var (
	loop int
	threadNum int
	txId string
	sdkPath string
)

func QpsCMD() *cobra.Command {

	keyCmd := &cobra.Command {
		Use:   "qps",
		Short: "Pressing chainMaker key command",
		Long:  "Pressing chainMaker key command",
		RunE: func(_ *cobra.Command, _ []string) error {
			return RunQps()
		},
	}
	flags := keyCmd.Flags()
	flags.IntVarP(&loop, "loop", "l", 1, "合约执行循环数量 eg. 1000")
	flags.IntVarP(&threadNum, "threadNum", "t", 1, "合约执行循环数量 eg. 1000")
	flags.StringVarP(&txId, "txId", "i", "", "合约参数")
	flags.StringVarP(&sdkPath, "sdkPath", "s", "", "SdkConfig路径")
	return keyCmd
}