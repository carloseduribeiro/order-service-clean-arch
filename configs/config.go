package configs

import "github.com/spf13/viper"

type conf struct {
	DBDriver             string `mapstructure:"DB_DRIVER"`
	DBHost               string `mapstructure:"DB_HOST"`
	DBPort               string `mapstructure:"DB_PORT"`
	DBUser               string `mapstructure:"DB_USER"`
	DBPassword           string `mapstructure:"DB_PASSWORD"`
	DBName               string `mapstructure:"DB_NAME"`
	WebServerPort        string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort       string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort    string `mapstructure:"GRAPHQL_SERVER_PORT"`
	RabbitMQUserName     string `mapstructure:"RABBITMQ_USERNAME"`
	RabbitMQUserPassword string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQServerHost   string `mapstructure:"RABBITMQ_SERVER_HOST"`
	RabbitMQServerPort   string `mapstructure:"RABBITMQ_SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.AutomaticEnv()
	viper.AddConfigPath(path)
	viper.SetConfigName("app_config")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
