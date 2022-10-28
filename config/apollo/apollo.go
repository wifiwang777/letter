package apollo

import (
	globalconfig "github.com/aisuosuo/letter/config"
	"github.com/aisuosuo/letter/config/log"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"
	"github.com/apolloconfig/agollo/v4/env/config"
)

var GlobalApolloConfig agcache.CacheInterface

func init() {
	c := &config.AppConfig{
		AppID:          globalconfig.GlobalConfig.ApolloConfig.AppId,
		Cluster:        "dev",
		IP:             globalconfig.GlobalConfig.ApolloConfig.Ip,
		NamespaceName:  "application",
		IsBackupConfig: true,
		Secret:         globalconfig.GlobalConfig.ApolloConfig.SecretKey,
	}
	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	GlobalApolloConfig = client.GetConfigCache(c.NamespaceName)
	log.Logger.Info("初始化Apollo配置成功")
}
