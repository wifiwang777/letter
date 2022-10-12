package config

import (
	"encoding/json"
	"fmt"
	"github.com/aisuosuo/ecdsa"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func init() {
	var err error
	if AppPath, err = filepath.Abs(filepath.Dir(filepath.Dir(os.Args[0]))); err != nil {
		panic(err)
	}

	configPath := fmt.Sprintf("%s/config/config.toml", AppPath)
	viper.SetConfigFile(configPath)
	viper.SetConfigType("toml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err = viper.Unmarshal(&GlobalConfig); err != nil {
		fmt.Println(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = viper.Unmarshal(&GlobalConfig); err != nil {
			fmt.Println(err)
		}
	})

	err = RefreshJwt(GlobalConfig.Jwt)
	if err != nil {
		panic(err)
	}

	marshal, _ := json.Marshal(GlobalConfig)
	fmt.Println(fmt.Sprintf("load config success:%s, config:%s", configPath, marshal))
}

func RefreshJwt(jwtConfig *JwtConfig) (err error) {
	JwtPublicKey, err = ecdsa.NewPublicKeyFromHex(jwtConfig.PublicKey)
	JwtPrivateKey, err = ecdsa.NewPrivateKeyFromHex(jwtConfig.PrivateKey)
	return
}
