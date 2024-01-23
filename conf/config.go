package conf

import (
	"bufio"
	"encoding/json"
	"os"
)

type config struct {
	Env   string `json:"env"`
	Port  int    `json:"port"`
	Mysql Mysql  `json:"mysql"`
	Redis Redis  `json:"redis"`
	Log   Log    `json:"log"`
	Email Email  `json:"email"`
	Nacos Nacos  `json:"nacos"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       int    `json:"db"`
	Password string `json:"password"`
}

type Log struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

type NacosEndpoint struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}

type Nacos struct {
	Endpoints []NacosEndpoint `json:"endpoints"`
	Username  string          `json:"username"`
	Password  string          `json:"password"`
	Namespace string          `json:"namespace"`
}

type Email struct {
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Sender   string   `json:"sender"`
	To       []string `json:"to"`
}

var Config *config

var LoginExcludeAuth = map[string][]string{
	"GET":    {"/auth/*", "/public_key"},
	"POST":   {},
	"PUT":    {},
	"DELETE": {},
} // 存放不校验的URL

func InitConfig() {
	Config = &config{}
	file, err := os.Open("conf/config.json")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&Config); err != nil {
		panic(err)
	}
}
