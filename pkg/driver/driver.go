package driver

import (
	"fmt"

	"github.com/DinnieJ/njav-downloader/pkg/config"
	"github.com/DinnieJ/njav-downloader/pkg/logger"
	"github.com/tebeka/selenium"
)

type WebDriver struct {
	service   *selenium.Service
	webdriver selenium.WebDriver
}

var LOGGER *logger.Logger

func init() {
	LOGGER = logger.GetLogger(&logger.LoggerConfig{
		Name:  "Driver",
		Level: logger.DEBUG,
	})
}

func (w *WebDriver) InitWebDriver(config *config.Config) error {
	fullPath, err := GetDriverPath(config)
	if err != nil {
		return err
	}

	service, err := selenium.NewChromeDriverService(fullPath, 9515)
	if err != nil {
		print(err.Error())
		return err
	}
	w.service = service

	caps := selenium.Capabilities{
		"browserName": "chrome",
		"goog:chromeOptions": map[string]interface{}{
			"args": []string{
				"--headless", // Run Chrome in headless mode (optional)
				"--disable-gpu",
				"--no-sandbox",
			},
			"perfLoggingPrefs": map[string]interface{}{
				"enableNetwork": true, // Capture network events
				"enablePage":    true, // Capture page events (e.g., page load time)
			},
		},
		"goog:loggingPrefs": map[string]interface{}{
			"performance": "ALL",
		},
	}
	wd, err := selenium.NewRemote(caps, "http://localhost:9515/wd/hub")
	if err != nil {
		return err
	}

	w.webdriver = wd
	return nil
}

func (w *WebDriver) GetWebDriver() selenium.WebDriver {
	return w.webdriver
}

func (w *WebDriver) Terminate() error {
	if err := w.webdriver.Quit(); err != nil {
		return fmt.Errorf("Failed to stop webdriver")
	}

	if err := w.service.Stop(); err != nil {
		return fmt.Errorf("Failed to stop selenium service")
	}
	return nil
}
