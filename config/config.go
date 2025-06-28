package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AccountId       string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	DatabaseUrl     string
}

func LoadConfig() *Config {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		AccountId:       viper.GetString("R2_ACCOUNT_ID"),
		AccessKeyId:     viper.GetString("R2_ACCESS_KEY_ID"),
		AccessKeySecret: viper.GetString("R2_ACCESS_KEY_SECRET"),
		BucketName:      viper.GetString("R2_BUCKET_NAME"),
		DatabaseUrl:     viper.GetString("DATABASE_URL"),
	}
}
