package conf

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

type Mysql struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}
type redis struct {
	RedisAddr string
}
type service struct {
	AppMode  string
	HttpPort string
}
type Oss struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	MainDirectory   string
}

type RabbitMQ struct {
	Port     string
	Username string
	Password string
}

var (
	OSS        *Oss
	sql        *Mysql
	PageSize   int
	RedisAddr  string
	ServerAddr string
	MQ         *RabbitMQ
)

func Init() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./pkg/conf")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		} else {
			fmt.Println("config file was found but another error was produced")
		}
	}
	OSS = LoadOSS(viper.GetStringMapString("oss"))

	PageSize, _ = strconv.Atoi(viper.GetStringMapString("page")["page-size"])

	RedisAddr = viper.GetStringMapString("redis")["addr"]

	ServerAddr = viper.GetStringMapString("server")["addr"]

	MQ = LoadRabbitMQ(viper.GetStringMapString("rabbitmq"))

	sql = LoadMysql(viper.GetStringMapString("mysql"))
	path := strings.Join([]string{sql.user, ":", sql.password, "@tcp(", sql.host, ":", sql.port, ")/", sql.dbname, "?charset=utf8mb4&parseTime=True"}, "")
	return path
}

func LoadMysql(myConf map[string]string) *Mysql {
	return &Mysql{
		host:     myConf["host"],
		port:     myConf["port"],
		user:     myConf["user"],
		password: myConf["password"],
		dbname:   myConf["dbname"],
	}
}
func LoadOSS(myConf map[string]string) *Oss {
	return &Oss{
		EndPoint:        myConf["endpoint"],
		AccessKeyId:     myConf["accesskeyid"],
		AccessKeySecret: myConf["accesskeysecret"],
		BucketName:      myConf["bucketname"],
		MainDirectory:   myConf["main-directory"],
	}
}

func LoadRabbitMQ(myConf map[string]string) *RabbitMQ {
	return &RabbitMQ{
		Port:     myConf["port"],
		Username: myConf["name"],
		Password: myConf["password"],
	}
}
