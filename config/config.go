package config

import (
	"bytes"
	"encoding/json"
	"strings"
)

// AppConfig is global var represent to application config
var AppConfig Config

type Config struct {
	Log                  LoggerConfig         `json:"log" mapstructure:"log"`
	Server               server.Config        `json:"server" mapstructure:"server"`
	MySQL                database.MySQLConfig `json:"mysql" mapstructure:"mysql"`
	MigrationsFolder     string               `json:"migrations_folder" mapstructure:"migrations_folder"`
	PubSub               PubSub               `json:"pubsub" mapstructure:"pubsub"`
	Env                  string               `json:"env" mapstructure:"env"`
	ConsumerWorkerNumber int                  `json:"consumer_worker" mapstructure:"consumer_worker"`
	Vietguys             Vietguys             `json:"vietguys" mapstructure:"vietguys"`
	OneID                OneID                `json:"oneid" mapstructure:"oneid"`
	VNPay                VNPay                `json:"vnpay" mapstructure:"vnpay"`
	App                  ApplicationConfig    `json:"app" mapstructure:"app"`
	Redis                Redis                `json:"redis" mapstructure:"redis"`
	Jwt                  Jwt                  `json:"jwt" mapstructure:"jwt"`
	Iam                  IamConfig            `json:"iam" mapsstructure:"iam"`
}

type ApplicationConfig struct {
	EnableSourceID bool `json:"enable_source_id" mapstructure:"enable_source_id"`
}

type Vietguys struct {
	SendURL  string `json:"send_url" mapstructure:"send_url"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

type OneID struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	TokenURL string `json:"token_url" mapstructure:"token_url"`
	SendURL  string `json:"send_url" mapstructure:"send_url"`
}

type VNPay struct {
	SendURL string     `json:"send_url" mapstructure:"send_url"`
	PV      VNPayBrand `json:"pv" mapstructure:"pv"`
	VNShop  VNPayBrand `json:"vnshop" mapstructure:"vnshop"`
}

type VNPayBrand struct {
	PartnerCode string `json:"partner_code" mapstructure:"partner_code"`
	SecretKey   string `json:"secret_key" mapstructure:"secret_key"`
}

type Jwt struct {
	Secret string
}

func Load() (*Config, error) {

	// You should set default config value here
	c := &Config{
		MySQL:                database.MySQLDefaultConfig(),
		Log:                  LoggerDefaultConfig(),
		Server:               server.DefaultConfig(),
		MigrationsFolder:     "file://migrations",
		Env:                  constant.LocalEnvironment,
		ConsumerWorkerNumber: 10,
		Jwt:                  Jwt{Secret: "mnbvcxz"},
		Iam:                  IamDefaultConfig(),
	}

	// --- hacking to load reflect structure config into env ----//
	viper.SetConfigType("json")
	configBuffer, err := json.Marshal(c)

	if err != nil {
		return nil, err
	}

	viper.ReadConfig(bytes.NewBuffer(configBuffer))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// -- end of hacking --//
	viper.AutomaticEnv()
	err = viper.Unmarshal(c)

	AppConfig = *c
	return c, err
}
