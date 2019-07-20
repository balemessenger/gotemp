package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	confOnce sync.Once
	config   *Config
)

type Config struct {
	Conf ConfYaml
}

func NewConfig() *Config {
	return &Config{}
}

func GetConfig() *Config {
	logOnce.Do(func() {
		config = NewConfig()
	})
	return config
}

var defaultConf = []byte(`
core:
  mode: "release" # release, debug, test
  user_work_pool_size: 1000
  group_work_pool_size: 100
postgres:
  host: ""
  port: 5432
  db: ""
  user: ""
  pass: ""
  batch_count: 5
kafka:
  bootstrap_servers: ""
  group_id: "random"
  auto_offset_reset: "earliest"
  topic: "{{ProjectName}}"
firebase:
  url: "http://fcm.googleapis.com/fcm/send"
  token: ""
  timeout: 10s
  max_connection: 10000
  work_pool_size: 10000
apple:
  enable: false
  key: 456367
  bundle_id: "ai.nasim.bale"
  path: "/home/amsjavan/PushCertificate.p12"
  password: "Elenoon@91"
prometheus:
  port: 8080
log:
  level: debug
endpoints:
  grpc:
    address: "127.0.0.1:5050"
  http:
    address: ":4040"
    user: "test"
    pass: "test"
`)

type ConfYaml struct {
	Core       SectionCore       `yaml:"core"`
	Postgres   SectionPostgres   `yaml:"postgres"`
	Kafka      SectionKafka      `yaml:"kafka"`
	Firebase   SectionFirebase   `yaml:"firebase"`
	Apple      SectionApple      `yaml:"apple"`
	Prometheus SectionPrometheus `yaml:"prometheus"`
	Log        SectionLog        `yaml:"log"`
	Endpoints  SectionEndpoints  `yaml:"endpoints"`
	Eureka     SectionEureka     `yaml:"eureka"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	Mode              string `yaml:"mode"`
	UserWorkPoolSize  int    `yaml:"user_work_pool_size"`
	GroupWorkPoolSize int    `yaml:"group_work_pool_size"`
}

// SectionPostgres is sub section of config.
type SectionPostgres struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	DB         string `yaml:"db"`
	User       string `yaml:"user"`
	Pass       string `yaml:"pass"`
	BatchCount int    `yaml:"batch_count"`
}

// SectionKafka is sub section of config.
type SectionKafka struct {
	BootstrapServers string `yaml:"bootstrap_servers"`
	GroupId          string `yaml:"group_id"`
	AutoOffsetReset  string `yaml:"auto_offset_reset"`
	Topic            string `yaml:"topic"`
}

type SectionFirebase struct {
	Url           string        `yaml:"url"`
	Token         string        `yaml:"token"`
	Timeout       time.Duration `yaml:"time"`
	MaxConnection int           `yaml:"max_connection"`
	WorkPoolSize  int           `yaml:"work_pool_size"`
}

type SectionApple struct {
	Enable   bool   `yaml:"enable"`
	Key      int32  `yaml:"key"`
	BundleId string `yaml:"bundle_id"`
	Path     string `yaml:"path"`
	Password string `yaml:"password"`
}

type SectionPrometheus struct {
	Port int `yaml:"port"`
}

type SectionLog struct {
	Level string `yaml:"level"`
}

type SectionEndpoints struct {
	Grpc SectionGrpc `yaml:"grpc"`
	Http SectionHttp `yaml:"http"`
}

type SectionEureka struct {
	Addresses  []string `yaml:"addresses"`
	AppName    string   `yaml:"app_name"`
	InstanceId string   `yaml:"instance_id"`
	VipAddress string   `yaml:"vip_address"`
	Ip         string   `yaml:"ip"`
	Port       int      `yaml:"port"`
}

type SectionGrpc struct {
	Address string `yaml:"address"`
}

type SectionHttp struct {
	Address string `yaml:"address"`
	User    string `yaml:"user"`
	Pass    string `yaml:"pass"`
}

type SectionGraylog struct {
	Level       string `yaml:"level"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Facility    string `yaml:"facility"`
	Compression bool   `yaml:"compression"`
}

type SectionStdout struct {
	Level string `yaml:"level"`
}

type SectionFile struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

// LoadConf load config from file and read in environment variables that match
func (config *Config) LoadConf(confPath string) (ConfYaml, error) {
	var conf ConfYaml

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()                  // read in environment variables that match
	viper.SetEnvPrefix("{{ProjectName}}") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			log.Errorf("File does not exist : %s", confPath)
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		// Search config in home directory with name ".pkg" (without extension).
		viper.AddConfigPath("/etc/{{ProjectName}}/")
		viper.AddConfigPath("$HOME/.{{ProjectName}}")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			// load default config
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
				return conf, err
			}
		}
	}

	// Core
	conf.Core.Mode = viper.GetString("core.mode")
	conf.Core.UserWorkPoolSize = viper.GetInt("core.user_work_pool_size")
	conf.Core.GroupWorkPoolSize = viper.GetInt("core.group_work_pool_size")

	// Postgres
	conf.Postgres.Host = viper.GetString("postgres.host")
	conf.Postgres.Port = viper.GetInt("postgres.port")
	conf.Postgres.DB = viper.GetString("postgres.db")
	conf.Postgres.User = viper.GetString("postgres.user")
	conf.Postgres.Pass = viper.GetString("postgres.pass")
	conf.Postgres.BatchCount = viper.GetInt("postgres.batch_count")

	// Kafka
	conf.Kafka.BootstrapServers = viper.GetString("kafka.bootstrap_servers")
	conf.Kafka.GroupId = viper.GetString("kafka.group_id")
	conf.Kafka.AutoOffsetReset = viper.GetString("kafka.auto_offset_reset")
	conf.Kafka.Topic = viper.GetString("kafka.topic")

	// Firebase
	conf.Firebase.Url = viper.GetString("firebase.url")
	conf.Firebase.Token = viper.GetString("firebase.token")
	conf.Firebase.Timeout = viper.GetDuration("firebase.timeout")
	conf.Firebase.MaxConnection = viper.GetInt("firebase.max_connection")
	conf.Firebase.WorkPoolSize = viper.GetInt("firebase.work_pool_size")

	// Apple
	conf.Apple.Enable = viper.GetBool("apple.enable")
	conf.Apple.Key = viper.GetInt32("apple.key")
	conf.Apple.BundleId = viper.GetString("apple.bundle_id")
	conf.Apple.Path = viper.GetString("apple.path")
	conf.Apple.Password = viper.GetString("apple.password")

	// Prometheus
	conf.Prometheus.Port = viper.GetInt("prometheus.port")

	//Log
	conf.Log.Level = viper.GetString("log.level")

	//Endpoints
	conf.Endpoints.Grpc.Address = viper.GetString("endpoints.grpc.address")
	conf.Endpoints.Http.Address = viper.GetString("endpoints.http.address")
	conf.Endpoints.Http.User = viper.GetString("endpoints.http.user")
	conf.Endpoints.Http.Pass = viper.GetString("endpoints.http.pass")

	return conf, nil
}

func (config *Config) Initialize(path string) {
	var err error
	config.Conf, err = config.LoadConf(path)
	if err != nil {
		log.Fatalf("Load yaml config file error: '%v'", err)
		return
	}
}
