package ctps

import "github.com/spf13/cobra"



var (
	loop int
	concurrency int
	threadNum int
	name string
	method string
	args string
	sdkPath string
)

func CtpsCMD() *cobra.Command {

	keyCmd := &cobra.Command {
		Use:   "ctps",
		Short: "Pressing chainMaker key command",
		Long:  "Pressing chainMaker key command",
		RunE: func(_ *cobra.Command, _ []string) error {
			return RunCTps()
		},
	}
	flags := keyCmd.Flags()
	flags.IntVarP(&loop, "loop", "l", 1, "合约执行循环数量 eg. 1000")
	flags.IntVarP(&concurrency, "concurrency", "c", 1, "并发数量. eg. 10")
	flags.StringVarP(&name, "name", "n", "", "合约名称")
	flags.StringVarP(&method, "method", "m", "", "合约内的方法")
	flags.StringVarP(&args, "args", "a", "", "合约参数")
	flags.StringVarP(&sdkPath, "sdkPath", "s", "", "SdkConfig路径")
	return keyCmd
}