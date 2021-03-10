package cmd

import (
	"fmt"
	"os"

	"github.com/GulshanArora7/ipsec_exporter/exporter"
	"github.com/GulshanArora7/ipsec_exporter/ipsec"
	"github.com/spf13/cobra"
)

const (
	flagIPSecConfigFile  = "config-path"
	flagWebListenAddress = "web.listen-address"
	flagExcludeConnName  = "exclude-conn"
)

// Version Variable
var Version string

// RootCmd Variable
var RootCmd = &cobra.Command{
	Use:     "ipsec_exporter",
	Short:   "Prometheus exporter for ipsec status.",
	Long:    "",
	Run:     defaultCommand,
	Version: Version,
}

func init() {
	RootCmd.PersistentFlags().StringVar(&exporter.IPSecConfigFile, flagIPSecConfigFile,
		"/etc/strongswan/ipsec.conf",
		"Path to the ipsec config file.")

	RootCmd.PersistentFlags().StringVar(&exporter.WebListenAddress, flagWebListenAddress,
		"0.0.0.0:9110",
		"Address on which to expose metrics.")

	RootCmd.PersistentFlags().StringVar(&ipsec.ExcludeConnName, flagExcludeConnName,
		"",
		"Tunnel Connection Name to Exclude.")
}

// Execute Function
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func defaultCommand(_ *cobra.Command, _ []string) {
	exporter.Serve()
}
