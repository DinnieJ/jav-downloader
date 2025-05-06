package cmd

import (
	"fmt"

	"github.com/DinnieJ/njav-downloader/pkg/config"
	msg "github.com/DinnieJ/njav-downloader/pkg/i18n"
	"github.com/spf13/cobra"
)

type App struct {
	Version string
	AppName string
	mainCmd *cobra.Command
}

func NewApplication(version string, appName string, BuildSrc string) *App {
	mainCmd := &cobra.Command{
		Use:   fmt.Sprintf("%s [code]", appName),
		Short: "NJAV Downloader",
		Long:  "Input the code, then wa la !",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf(msg.MSG_MISSING_CODE)
			}
			return nil
		},
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(msg.MSG_HELLO_WORLD)
			// t := driver.WebDriver{}
			// if err := t.InitWebDriver(); err != nil {
			// 	panic(err)
			// }
			config := &config.Config{}
			if err := config.Init(); err != nil {
				fmt.Println("Unable to init config", err)
			}

			// defer t.Terminate()
		},
	}
	// mainCmd.

	return &App{
		mainCmd: mainCmd,
		Version: version,
		AppName: appName,
	}
}

func (a *App) Start() error {
	if err := a.mainCmd.Execute(); err != nil {
		return err
	}

	return nil
}
