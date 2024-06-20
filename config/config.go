package config

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	RestServer struct {
		Port string
	}
	DbURL   string
	JsonRpc string
	ChainID int64
	Kafka   struct {
		Brokers []string
	}
	Redis struct {
		Addr     string
		Password string
		Username string
	}
	JwtSecret string
	Contracts struct {
		Vdex string
	}
	Tokens            map[string]string
	DispatcherWallets []string
	GaslessWallets    []string
	Env               string
	PushServer        struct {
		Port string
		Path string
	}
}

var conf *Configuration

func Init(confPath string) {
	conf = &Configuration{}
	viper.SetConfigFile(confPath)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Warning(err, "config file not found")
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Panic(err, "Error Unmarshal Viper Config")
	}
	logrus.SetReportCaller(true)
	spew.Dump("vdex: ", conf)
}

func GetConfig() *Configuration {
	return conf
}

func GetAddressBySymbol(s string) string {
	for addr, symbol := range conf.Tokens {
		if strings.EqualFold(symbol, s) {
			return addr
		}
	}

	return ""
}

func GetDispatcherByID(i int) string {
	return conf.DispatcherWallets[i]
}

func GetGaslessDispatcherByID(i int) string {
	return conf.GaslessWallets[i]
}

func GetRedisURL() string {
	urlStr := fmt.Sprintf("redis://%v:%v@%v/%v", GetConfig().Redis.Username, GetConfig().Redis.Password, GetConfig().Redis.Addr, 0)
	if GetConfig().Env != "local" {
		urlStr = fmt.Sprintf("rediss://%v:%v@%v/%v", GetConfig().Redis.Username, GetConfig().Redis.Password, GetConfig().Redis.Addr, 0)
	}

	return urlStr
}
