package main

import (
	"chainpress/pkg/ctps"
	"chainpress/pkg/qps"
	"chainpress/pkg/tps"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)




func main()  {

	var mainCmd = &cobra.Command{
		Use:   "press",
		Short: "Press test cli",
		Long: strings.TrimSpace(`Command line interface for press ChainMaker`),
	}
	mainCmd.AddCommand(tps.TpsCMD())
	mainCmd.AddCommand(ctps.CtpsCMD())
	mainCmd.AddCommand(qps.QpsCMD())

	if err := mainCmd.Execute(); err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
}
