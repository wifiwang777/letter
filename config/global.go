package config

import (
	"fmt"
	"github.com/aisuosuo/ecdsa"
)

var (
	GlobalConfig  Server
	AppPath       string
	JwtPublicKey  *ecdsa.PublicKey
	JwtPrivateKey *ecdsa.PrivateKey
)

type MysqlConfig struct {
	Host         string `toml:"host"`
	Port         string `toml:"port"`
	Username     string `toml:"username"`
	Password     string `toml:"password"`
	Config       string `toml:"config"`
	Dbname       string `toml:"dbName"`
	MaxIdleConns int    `toml:"maxIdleConns"` // 空闲中的最大连接数
	MaxOpenConns int    `toml:"maxOpenConns"` // 打开到数据库的最大连接数
}

func (m *MysqlConfig) Dsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.Dbname, m.Config)
	return dsn
}

type LogConfig struct {
	Path  string `mapstructure:"path"`
	Level string `mapstructure:"level"`
}

type Run struct {
	Mode          string `mapstructure:"mode"`
	HttpAddr      string `mapstructure:"httpAddr"`
	WebsocketAddr string `mapstructure:"websocketAddr"`
}

type JwtConfig struct {
	PublicKey  string `mapstructure:"publicKey"`
	PrivateKey string `mapstructure:"privateKey"`
}

type FileConfig struct {
	FilePath string `mapstructure:"filePath"`
}

type ApolloConfig struct {
	AppId     string `mapstructure:"appId"`
	Ip        string `mapstructure:"ip"`
	SecretKey string `mapstructure:"secretKey"`
}

type Server struct {
	MysqlConfig  *MysqlConfig  `mapstructure:"mysql"`
	LogConfig    *LogConfig    `mapstructure:"log"`
	RunConfig    *Run          `mapstructure:"run"`
	JwtConfig    *JwtConfig    `mapstructure:"jwt"`
	FileConfig   *FileConfig   `mapstructure:"file"`
	ApolloConfig *ApolloConfig `mapstructure:"apollo"`
}
