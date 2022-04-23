package configs

import (
	"github.com/jinzhu/configor"
	"quizbe/utils"
)

type Configurations struct {
	Sql             DatabaseDetails
	Redis           RedisClients
	Jwt             JwtToken
	SecretKeySentry string
	ProjectSentry   string
	Port            string
	AppName         string
}

type JwtToken struct {
	SecretKey           string
	SecretExpireMinutes string
}

type DatabaseDetails struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisClients struct {
	Name      string
	Host      string
	Port      string
	Password  string
	RedisDb   string
	KeyPrefix string
}

var (
	Configs    Configurations
	configPath string
)

func LoadConfig() {
	pathName := utils.CurrentDirectoryPathName()

	configPath = pathName + "/configs/config.dev.yaml"
	_ = configor.New(&configor.Config{Debug: true}).Load(&Configs, configPath)
}
