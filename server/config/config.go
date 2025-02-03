package config

type EnvConfig struct{
	ServerPort string `env:"SERVER_PORT"`
	DB_HOST string `env:"DB_HOST"`
	DB_NAME string `env:"DB_NAME"`
	DB_USER string `env:"DB_USER"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
	DB_SSLMODE string `env:"DB_SSLMODE"`
}

func NewEnvConfig() *EnvConfig {
	err := godotenv()

	if err != nil{
		log.Fatal("Error loading .env file")
	}

	config := &EnvConfig{}

	err = env.Parse(config)

	if err != nil{
		log.Fatal("Error parsing .env file")
	}

	return config
}