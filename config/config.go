package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type RunMode string

const (
	Debug   RunMode = "debug"
	Release         = "release"
)

type Config struct {
	Database struct {
		Host                  string
		Port                  int
		User                  string
		Password              string
		Name                  string
		MaxConnection         int
		MinConnection         int
		MaxLifeTimeConnection float64
		MaxIdleTimeConnection float64
	}

	Cache struct {
		Size int
	}

	App struct {
		Port        int
		KeyPath     string
		CertPath    string
		Version     string
		Name        string
		EnablePPROF bool
	}
	UseCache              bool
	EnableCacheStatistics bool

	Kafka struct {
		Server      string
		Topic       string
		NumConsumer uint
		ClientID    string
		GroupID     string
		Assignor    string
		Version     string
	}
}

var (
	AppConfig = &Config{}
)

func InitConfig() {
	runMode := os.Getenv("gintoki_RUNMODE")
	if strings.ToLower(runMode) == "debug" || runMode == "" {
		viper.SetConfigFile(".env")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	} else {
		viper.AllowEmptyEnv(true) // Allows an environment variable to not exist and not blow up, I suggest using switch statements to handle these though
		viper.AutomaticEnv()      // Do the darn thing :D
	}

	// Database config
	AppConfig.Database.Host = viper.GetString("gintoki_DB_HOST")
	AppConfig.Database.Port = viper.GetInt("gintoki_DB_PORT")
	AppConfig.Database.User = viper.GetString("gintoki_DB_USER")
	AppConfig.Database.Password = viper.GetString("gintoki_DB_PASSWORD")
	AppConfig.Database.Name = viper.GetString("gintoki_DB_NAME")
	AppConfig.Database.MaxConnection = viper.GetInt("gintoki_DB_MAXCONN")
	AppConfig.Database.MinConnection = viper.GetInt("gintoki_DB_MINCONN")
	AppConfig.Database.MaxLifeTimeConnection = viper.GetFloat64("gintoki_DB_CONN_LIFETIME")
	AppConfig.Database.MaxIdleTimeConnection = viper.GetFloat64("gintoki_DB_CONN_IDLETIME")

	AppConfig.App.Port = viper.GetInt("gintoki_APP_PORT")
	AppConfig.App.Version = viper.GetString("gintoki_APP_VERSION")
	AppConfig.App.Name = viper.GetString("gintoki_APP_NAME")
	AppConfig.App.EnablePPROF = viper.GetBool("gintoki_ENABLE_PPROF")

	AppConfig.UseCache = viper.GetBool("gintoki_USE_CACHE")
	AppConfig.EnableCacheStatistics = viper.GetBool("gintoki_ENABLE_CACHE_STATISTICS")

	// Kafka
	AppConfig.Kafka.Server = viper.GetString("gintoki_KAFKA_SERVER")
	AppConfig.Kafka.Topic = viper.GetString("gintoki_KAFKA_TOPIC")
	AppConfig.Kafka.NumConsumer = viper.GetUint("gintoki_KAFKA_NUMCONSUMER")
	AppConfig.Kafka.ClientID = viper.GetString("gintoki_KAFKA_CLIENT_ID")
	AppConfig.Kafka.GroupID = viper.GetString("gintoki_KAFKA_GROUP_ID")
	AppConfig.Kafka.Assignor = viper.GetString("gintoki_KAFKA_ASSIGNOR")
	AppConfig.Kafka.Version = viper.GetString("gintoki_KAFKA_VERSION")
}
