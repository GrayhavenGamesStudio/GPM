package userconfig

import (
	"os"
	"path/filepath"
)

// GetConfigPath returns the appropriate config file path based on OS
func GetConfigPath(appName string) (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, appName, "config.yaml"), nil
}
