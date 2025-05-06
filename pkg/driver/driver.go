package driver

import (
	"fmt"

	"github.com/tebeka/selenium"
)

type WebDriver struct {
	tempDriverDir string
	service       *selenium.Service
	webdriver     selenium.WebDriver
}

func (w *WebDriver) InitWebDriver() error {
	fullPath, err := GetDriverPath()
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

func (w *WebDriver) Terminate() error {
	if err := w.webdriver.Quit(); err != nil {
		return fmt.Errorf("Failed to stop webdriver")
	}

	if err := w.service.Stop(); err != nil {
		return fmt.Errorf("Failed to stop selenium service")
	}
	return nil
}
