package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sillyhatxu/convenient-utils/envconfig"
	"github.com/sillyhatxu/convenient-utils/logconfig"
	"github.com/sillyhatxu/convenient-utils/mysqlclient"
	"github.com/sillyhatxu/table-backend/api"
	"github.com/sillyhatxu/table-backend/config"
	"github.com/sirupsen/logrus"
)

func init() {
	cfgFile := flag.String("c", "config-local.conf", "configuration file")
	flag.Parse()
	envconfig.ParseConfig(*cfgFile, func(content []byte) {
		err := toml.Unmarshal(content, &config.Conf)
		if err != nil {
			logrus.Panicf("unmarshal toml object error. %v", err)
		}
	})
	logConfig := logconfig.NewLogConfig(
		logrus.InfoLevel, true,
		config.Conf.Log.Project,
		config.Conf.Log.Module,
		config.Conf.Log.OpenLogstash,
		config.Conf.Log.LogstashAddress,
		config.Conf.Log.OpenLogfile,
		config.Conf.Log.FilePath,
	)
	logConfig.InitialLogConfig()
}

func main() {
	mysqlclient.InitialDBClient(config.Conf.MysqlDB.DataSource, config.Conf.MysqlDB.MaxIdleConns, config.Conf.MysqlDB.MaxOpenConns)
	api.InitialAPI()
	logrus.Info("---------- Project Close ----------")
}
