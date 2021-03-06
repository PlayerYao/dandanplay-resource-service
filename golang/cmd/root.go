package cmd

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"dandanplay-resource-service/config"
	"dandanplay-resource-service/rotuers"
	"dandanplay-resource-service/utils"
	"dandanplay-resource-service/utils/logger"
)

var (
	isVersion bool
)

var rootCmd = &cobra.Command{
	Use:   config.Name,
	Short: config.ShortDescription,
	Long:  config.LongDescription,
	Run:   rootHandler,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Define global flags
	//rootCmd.PersistentFlags().BoolVarP()

	// Define local flags
	rootCmd.PersistentFlags().BoolVarP(
		&isVersion, "version", "V", false,
		fmt.Sprintf("Print the version number of %s", config.Name),
	)
	rootCmd.PersistentFlags().StringVarP(&config.Host, "host", "H", "",
		`Listen address for the API, such as "127.0.0.1", "0.0.0.0", or "localhost"`)
	rootCmd.PersistentFlags().StringVarP(&config.Port, "port", "P", "8080",
		"Listen address for the API")
	rootCmd.PersistentFlags().StringVarP(&config.Proxy, "proxy", "x", "",
		`Proxy address for web scraper, "http" and "socks5" are supported`)
}

func rootHandler(cmd *cobra.Command, args []string) {
	if isVersion == true {
		utils.PrintVersionInfo()
		return
	}

	gin.SetMode(gin.ReleaseMode)

	r := rotuers.InitRouter()

	if config.Port != "" {
		config.Port = ":" + config.Port
	}
	logger.Infof("{{Listening and serving HTTP on %s%s}}", config.Host, config.Port)

	if err := r.Run(config.Host + config.Port); err != nil {
		logger.Errorf("{{Server start failed.}} %v\n", err)
	}
}
