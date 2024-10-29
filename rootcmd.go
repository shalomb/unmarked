package main

import (
	"fmt"
	"os"
	"path"

	"github.com/adrg/xdg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile        string
	cfgDir         string = path.Join(xdg.StateHome, "unmarked")
	defaultCfgFile string = path.Join(cfgDir, "unmarked.yaml")
	userLicense    string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "unmarked",
		Short: "Unmarked marks windows for use with shortcuts",
		Long:  `"Unmarked marks windows for use with shortcuts"`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	var debug bool
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug logging")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "default config file")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	// Cobra also supports local flags, which will only run when this
	// action is called directly.
	// res, err := rootCmd.Flags().GetBool("toggle")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Set a toggle")
}

// InitCobra adds all child commands to the root command and sets flags appropriately.
// This is called by main.init(). It only needs to happen once to the rootCmd.
func InitCobra() error {
	return rootCmd.Execute()
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(xdg.Home)
		viper.AddConfigPath(path.Join(xdg.ConfigHome, AppName))
		viper.AddConfigPath(path.Join(xdg.StateHome, AppName))
		viper.SetConfigType("yaml")
		viper.SetConfigName(AppName)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warnf("Error, config file not found: %v", cfgFile)
		}
	}
	log.Debugf("Using config file: %v", viper.ConfigFileUsed())

	rootCmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when
		// represented in the config file
		configName := f.Name
		log.Debugf("Processing flag: %v", f.Name)

		// Apply the viper config value to the flag
		// when the flag is not set and viper has a value
		if !f.Changed && viper.IsSet(configName) {
			val := viper.Get(configName)
			rootCmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})

	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}
}
