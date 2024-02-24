package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

// Config struct for .env.yml
type Config struct {
	ForceTLS                       bool
	ImplementOrganizationStructure bool
	RegisterAutoLogin              bool
	AllowMultipleLogin             bool
	DefaultTimezone                *time.Location
	Server                         ServerConfig
	DB                             DBConfig
	DBLog                          DBConfig
	Redis                          RedisServer
	JWTConfig                      JWTConfig
	S3                             S3Configuration
	Firebase                       Firebase
	Mailer                         Mailer
	Whatsapp                       Whatsapp
	FrontEnd                       FrontEnd
	Facematch                      Facematch
	Oauth                          Oauth
	AppConfig                      map[string]interface{}
}

// ServerConfig struct to handle server configuration
type ServerConfig struct {
	AppName                string
	AppUrl                 string
	Addr                   string
	WriteTimeout           int
	ReadTimeout            int
	GraceFulTimeout        int
	MaxReceivedMessageSize int
	MaxImageHeight         int
	StorageProvider        string
	LocalStoragePath       string
	FrontEndLandingURL     string
}

// DBConfig struct to handle database configuration
type DBConfig struct {
	Name            string
	Host            string
	DbName          string
	CronHost        string
	CronDbName      string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime int
}

// RedisServer struct to handle redis configuration
type RedisServer struct {
	Addr     string
	Password string
	Timeout  int
	MaxIdle  int
}

// JWTConfig struct to handle JWT configuration
type JWTConfig struct {
	Issuer             string
	Secret             string
	ExpireAccessToken  int // by minutes
	ExpireRefreshToken int // by month
}

// S3Configuration struct to handle S3 configuration
type S3Configuration struct {
	Key        string
	Secret     string
	Region     string
	Bucket     string
	RootFolder string
	PublicUrl  string
}

// Mailer struct to handle mailer configuration
type Mailer struct {
	Server     string
	Port       int
	Username   string
	Password   string
	UseTls     bool
	Sender     string
	MaxAttempt int
}

// Firebase struct to handle Firebase configuration
type Firebase struct {
	ProjectID           string
	KeyFileDir          string
	Bucket              string
	StoragePublicAccess bool
}

// Whatsapp struct to handle Whatsapp configuration
type Whatsapp struct {
	BaseUrl   string
	AppSecret string
}

// Facematch struct to handle facematch configuration
type Facematch struct {
	SecretNumber string
}

// FrontEnd struct for response link to frontend
type FrontEnd struct {
	LinkUrl string
}

type Oauth struct {
	GoogleProjectID     string
	AppleAppID          string
	AppleTeamID         string
	AppleKeyID          string
	AppleCredentialPath string
}

// InitConfig function to init configuration, returns Config struct
func InitConfig(defaultTimezone *time.Location) Config {
	viper.SetConfigName(".env")
	if os.Getenv("ENV") != "" {
		viper.SetConfigName(".env-" + os.Getenv("ENV"))
	}

	viper.AddConfigPath(".")

	var configuration Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	if defaultTimezone != nil {
		configuration.DefaultTimezone = defaultTimezone
	} else {
		configuration.DefaultTimezone = time.UTC
	}

	return configuration
}
