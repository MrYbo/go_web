package config

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

type Yaml struct {
	Server   `yaml:"server"`
	Mysql    `yaml:"mysql"`
	Session  `yaml:"session"`
	Redis    `yaml:"redis"`
	Log      `yaml:"log"`
	Redirect `yaml:"redirect"`
	Token    `yaml:"token"`
}

type Server struct {
	Port               int    `yaml:"port"`
	Mode               string `yaml:"mode"`
	Authentication     string `yaml:"authentication"`
	AuthenticationSave string `yaml:"authenticationSave"`
}

type Mysql struct {
	User                      string `yaml:"user"`
	Password                  string `yaml:"password"`
	Path                      string `yaml:"path"`
	Database                  string `yaml:"database"`
	Config                    string `yaml:"config"`
	LogLevel                  int    `yaml:"logLevel"`
	MaxIdleConns              int    `yaml:"maxIdleConns"`
	MaxOpenConns              int    `yaml:"maxOpenConns"`
	AutoMigrate               bool   `yaml:"autoMigrate"`
	DefaultStringSize         uint   `yaml:"defaultStringSize"`
	DisableDatetimePrecision  bool   `yaml:"disableDatetimePrecision"`
	DontSupportRenameIndex    bool   `yaml:"dontSupportRenameIndex"`
	DontSupportRenameColumn   bool   `yaml:"dontSupportRenameColumn"`
	SkipInitializeWithVersion bool   `yaml:"skipInitializeWithVersion"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Session struct {
	Storage  string `yaml:"storage"`
	Key      string `yaml:"key"`
	Size     int    `yaml:"size"`
	MaxAge   int    `yaml:"maxAge"`
	Path     string `yaml:"path"`
	Domain   string `yaml:"domain"`
	HttpOnly bool   `yaml:"httpOnly"`
}

type Token struct {
	Secret string `yaml:"secret"`
	Issuer string `yaml:"issuer"`
	Expire int64  `yaml:"expire"`
}

type Redirect struct {
	Url string `yaml:"url"`
}

type Log struct {
	Debug    bool          `yaml:"debug"`
	MaxAge   time.Duration `yaml:"maxAge"`
	FileName string        `yaml:"fileName"`
	DirName  string        `yaml:"dirName"`
}

var Conf *Yaml

func init() {
	var defaultConfigFile string
	if len(os.Getenv("SERVER_ENV")) != 0 {
		defaultConfigFile = fmt.Sprintf("app/config/yaml/config.%s.yaml", os.Getenv("SERVER_ENV"))
	} else {
		defaultConfigFile = fmt.Sprintf("app/config/yaml/config.%s.yaml", "dev")
	}

	confFile := flag.String("c", defaultConfigFile, "help config path")
	flag.Parse()
	yamlConf, err := ioutil.ReadFile(*confFile)
	if err != nil {
		logrus.Error("get config file error:", err)
	}
	c := &Yaml{}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		logrus.Error("config init unmarshal failed:", err)
	}
	logrus.Info("config file load init success.")
	Conf = c
}
