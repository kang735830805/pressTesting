package tps

import "github.com/spf13/cobra"

var (
	//loop int
	concurrency int
	threadNum int
	name string
	method string
	parameter string
	sdkPath string
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
	//flags.IntVarP(&loop, "loop", "l", 1, "交易执行总数量 eg. 1000")
	flags.IntVarP(&concurrency, "concurrency", "c", 1, "进程内交易数量. eg. 10")
	flags.IntVarP(&threadNum, "threadNum", "t", 1, "进程数量. eg. 10")
	flags.StringVarP(&name, "name", "n", "", "合约名称")
	flags.StringVarP(&method, "method", "m", "", "合约内的方法")
	flags.StringVarP(&parameter, "parameter", "p", "", "合约参数")
	flags.StringVarP(&sdkPath, "sdkPath", "s", "", "SdkConfig路径")
	return keyCmd
}

