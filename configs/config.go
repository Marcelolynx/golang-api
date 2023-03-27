package config

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPass       string `mapstructure:"DB_PASS"`
	DBName       string `mapstructure:"DB_NAME"`
	WebPort      string `mapstructure:"WEB_PORT"`
	jwtSecret    string `mapstructure:"JWT_SECRET"`
	JwtExperesIn int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth    *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path)
	viper.SetConfigFile(path + "/app_config.json")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.jwtSecret), nil)
	return cfg, nil
}
