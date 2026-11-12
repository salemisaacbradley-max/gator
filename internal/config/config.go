package config

import ("os", "encoding/json", "fmt")

type Config struct {
	current_username string
	post 
}

func Read (data byte[]) (map[string]Config, error) {
	err := json.Unmarshal(data)
	if err != nil {
		return map[string]Config, fmt.Errorf("Error reading JSON: %v", err)
	}
	post := Post&

}

func SetUser (cfg Config) *Config {
	err := write(cfg.current_username)
}

func getConfigFilePath() (string, error) {

}

func write (cfg Config) error {
	err := 
}

const configFileName = ".gatorconfig.json"