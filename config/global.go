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
	Host         string `toml:"host"`         // 服务器地址:端口
	Port         string `toml:"port"`         //:端口
	Username     string `toml:"username"`     // 数据库用户名
	Password     string `toml:"password"`     // 数据库密码
	Config       string `toml:"config"`       // 高级配置
	Dbname       string `toml:"dbName"`       // 数据库名
	MaxIdleConns int    `toml:"maxIdleConns"` // 空闲中的最大连接数
	MaxOpenConns int    `toml:"maxOpenConns"` // 打开到数据库的最大连接数
}

func (m *MysqlConfig) Dsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.Dbname, m.Config)
	return dsn
}

type LogConfig struct {
	Path  string `toml:"path"`
	Level string `toml:"level"`
}

type Run struct {
	Mode          string `toml:"mode"`
	HttpAddr      string `toml:"httpAddr"`
	WebsocketAddr string `toml:"websocketAddr"`
}

type JwtConfig struct {
	PublicKey  string `toml:"publicKey"`
	PrivateKey string `toml:"privateKey"`
}

type FileConfig struct {
	FilePath string `toml:"filePath"`
}

type Server struct {
	Mysql *MysqlConfig `toml:"mysql"`
	Log   *LogConfig   `toml:"log"`
	Run   *Run         `toml:"run"`
	Jwt   *JwtConfig   `toml:"jwt"`
	File  *FileConfig  `toml:"file"`
}
