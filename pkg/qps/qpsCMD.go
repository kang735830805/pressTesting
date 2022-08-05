package qps

import "github.com/spf13/cobra"



var (
	threadNum int
	concurrency int
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
	flags.IntVarP(&threadNum, "threadNum", "t", 1, "进程数量 eg. 1000")
	flags.IntVarP(&concurrency, "concurrency", "c", 1, "进程内交易并发数量. eg. 10")
	flags.StringVarP(&txId, "txId", "i", "", "合约参数")
	flags.StringVarP(&sdkPath, "sdkPath", "s", "", "SdkConfig路径")
	return keyCmd
}