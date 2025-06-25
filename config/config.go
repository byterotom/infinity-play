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
}

func LoadConfig() *Config {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		AccountId:       viper.Get("R2_ACCOUNT_ID").(string),
		AccessKeyId:     viper.Get("R2_ACCESS_KEY_ID").(string),
		AccessKeySecret: viper.Get("R2_ACCESS_KEY_SECRET").(string),
		BucketName:      viper.Get("R2_BUCKET_NAME").(string),
	}
}
