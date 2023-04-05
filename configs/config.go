package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBDriver      string           `mapstructure: "DBDRIVER"`
	DBHost        string           `mapstructure: "DBHOST"`
	DBPort        string           `mapstructure: "DBPORT"`
	DBUser        string           `mapstructure: "DBUSER"`
	DBPassword    string           `mapstructure: "DBPASSWORD"`
	DBName        string           `mapstructure: "DBNAME"`
	WebServerPort string           `mapstructure: "WEBSERVERPORT"`
	JWTSecret     string           `mapstructure: "JWTSECRET"`
	JWTExperesIn  int              `mapstructure: "JWTEXPERESIN"`
	TokenAuth     *jwtauth.JWTAuth `mapstructure: "TOKENAUTH"`
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("ApiStudy")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New(" HS256", []byte(cfg.JWTSecret), nil)

	return nil, err
}
