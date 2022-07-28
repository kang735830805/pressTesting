package main

import (
	"chainpress/pkg/ctps"
	"chainpress/pkg/qps"
	"chainpress/pkg/tps"
	"github.com/spf13/cobra"
	"strings"
)



func main()  {
	//var winit=flag.Int("init",0,"是否需要初始化(安装链码，注册用户)")
	//var winvoke=flag.Int("invoke",0,"是否需要测试链码调用")
	//var wtest=flag.Int("test",1,"性能测试")

	mainCmd := &cobra.Command{
		Use:   "press",
		Short: "Press test cli",
		Long: strings.TrimSpace(`Command line interface for press ChainMaker
`),
	}

	mainCmd.AddCommand(tps.TpsCMD())
	mainCmd.AddCommand(ctps.CtpsCMD())
	mainCmd.AddCommand(qps.QpsCMD())


	//if 1==*winit {
	//	tps.Init()
	//}
	//if 1==*winvoke{
	//
		//Invoke()
	//}
	//if 1==*wtest{

	//tps.RunTps(*loopNum, *concurrencyNum, *contractName, *method, *args, *sdkConfigPath)
	//tps.RunTps(1, 1, "fact", "save", "{\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"6543234\"}", "./sdk_config.yml")
	//}
	mainCmd.Execute()
}
