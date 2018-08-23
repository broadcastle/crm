package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// var cfgFile string
// var debugLog bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "crm",
	Short: "A simple customer relationship management tool.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {

	if err := RootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}

	// if debugLog {
	// 	logrus.SetLevel(logrus.DebugLevel)
	// 	logrus.Info("debug text enabled")
	// }

	// manage.Init()
	// defer manage.Close()
}

func init() {
	// cobra.OnInitialize(initConfig)

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.crm.yaml)")
	// RootCmd.PersistentFlags().BoolVar(&debugLog, "debug", false, "debug output")

}

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			logrus.Fatal(err)
// 		}

// 		// Search config in home directory with name ".crm" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".crm")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		logrus.Debugf("Using config file: %s", viper.ConfigFileUsed())
// 	}
// }
