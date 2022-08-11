package cmd

import (
	"github.com/spf13/cobra"
	"strings"
)

var (

	LoopNum = RootCmd.PersistentFlags().IntP("loopNum", "l", 1, "循环数量. eg. 10")
	ThreadNum = RootCmd.PersistentFlags().IntP("threadNum", "t", 1, "进程数量. eg. 10")
	Interval = RootCmd.PersistentFlags().IntP("interval", "i", 1, "实验间隔. eg. 10为每10s发送一次实验")
	Duration = RootCmd.PersistentFlags().IntP("duration", "d", 1, "压测持续时间. eg. 10为10s")
	SdkPath = RootCmd.PersistentFlags().StringP("sdkPath", "s", "", "SdkConfig路径")
	RootCmd = &cobra.Command{
		Use:   "press",
		Short: "Press test cli",
		Long: strings.TrimSpace(`Command line interface for press ChainMaker`),
	}

)
