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
		// Run: func(cmd *cobra.Command, args []string) {
		// },
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is %v)", defaultCfgFile))
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	log.Printf("cfgFile: %+v", cfgFile)

	// Cobra also supports local flags, which will only run when this
	// action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	log.Printf("cfgFile: %+v", rootCmd.Flags().Lookup("toggle").Value)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.init(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warnf("Error, config file not found: %v", cfgFile)
		} else {
			log.Printf("Using config file: %v", viper.ConfigFileUsed())
		}
	}

	rootCmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name
		log.Printf("Processing flag: %v", f.Name)

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && viper.IsSet(configName) {
			val := viper.Get(configName)
			rootCmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})

}
