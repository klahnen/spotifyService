package config

import "gitlab.com/avarf/getenvs"

type Config struct {
	Port         int
	PostgresURL  string
	ClientID     string
	ClientSecret string
}

func GetConfig() Config {
	var config Config

	config.PostgresURL = getenvs.GetEnvString("DB_URL", "postgres://postgres:postgres@127.0.0.1:5432/spotifyService?sslmode=disable")
	config.Port, _ = getenvs.GetEnvInt("PORT", 8000)
	config.ClientID = getenvs.GetEnvString("CLIENT_ID", "99001291aa1f4dacbf5b381f1fd8af71")
	config.ClientSecret = getenvs.GetEnvString("CLIENT_SECRET", "4c7f1a7905d14e9bbc91116909a7bea1")

	return config
}
