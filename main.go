package main

import (
	"chainpress/pkg/cmd"
	"chainpress/pkg/ctps"
	"chainpress/pkg/qps"
	"chainpress/pkg/tps"
	"fmt"
	"github.com/spf13/cobra"
)

func InitCmd()  {
	cmd.RootCmd.AddCommand(tps.TpsCMD())
	cmd.RootCmd.AddCommand(ctps.CtpsCMD())
	cmd.RootCmd.AddCommand(qps.QpsCMD())
}

func GetRootCmd() *cobra.Command {
	return cmd.RootCmd
}

func main()  {

	InitCmd()
	err := GetRootCmd().Execute()
	if err != nil {
		fmt.Errorf(err.Error())
	}
}
