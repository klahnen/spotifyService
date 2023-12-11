package config

import "gitlab.com/avarf/getenvs"

type Config struct {
	Port         int
	ClientID     string
	ClientSecret string
	DBName       string
	RedirectURI  string
	BearerToken  string
}

func GetConfig() Config {
	var config Config

	config.Port, _ = getenvs.GetEnvInt("PORT", 8000)
	config.ClientID = getenvs.GetEnvString("CLIENT_ID", "99001291aa1f4dacbf5b381f1fd8af71")
	config.ClientSecret = getenvs.GetEnvString("CLIENT_SECRET", "4c7f1a7905d14e9bbc91116909a7bea1")
	config.DBName = getenvs.GetEnvString("DB_NAME", "/tmp/dev.db")
	config.RedirectURI = getenvs.GetEnvString("REDIRECT_URI", "http://127.0.0.1:8000/callback")
	config.BearerToken = getenvs.GetEnvString("BEARER_TOKEN", "")

	return config
}
