package config

import ("os"
 "encoding/json"
  "fmt" 
  "path/filepath")

type Config struct {
	CurrentUsername string "json:'username'"
	DbUrl string "json:'url'"
}

func Read () (Config, error) {
	var cfg Config
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("Error retrieving filepath: %v\n", err)
	}
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, fmt.Errorf("Error splicing Home directory with json: %v\n", err)
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("Error Unmarshalling: %v\n", err)
	}
	return cfg, nil
}

func SetUser (cfg Config) (Config, error) {
	err := write(cfg)
	if err != nil {
		return Config{}, fmt.Errorf("Error Unmarshalling: %v\n", err)
	}
	return cfg, nil
}

func getConfigFilePath() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error retrieving Home directory: %v\n", err)
	}
	fullPath := filepath.Join(userHome, configFileName)
	return fullPath, nil
}

func write (cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Error Marshalling: %v\n", err)
	}
	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Error retrieving filepath: %v\n", err)
	}
	err = os.WriteFile(filePath, data, 0600)
	if err != nil {
		return fmt.Errorf("Error Writing: %v\n", err)
	}
	return nil
}

const configFileName = ".gatorconfig.json"