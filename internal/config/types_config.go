package config

type Config struct {
	DatabaseURL string `json:"db_url"`
	UserName    string `json:"current_user_name"`
}
