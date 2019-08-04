package config

var Conf config

type mysqlDB struct {
	DataSource   string `toml:"data_source"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxOpenConns int    `toml:"max_open_conns"`
}

type http struct {
	Listen      string `toml:"listen"`
	Host        string `toml:"host"`
	Environment string `toml:"environment"`
}

type logConf struct {
	OpenLogstash    bool   `toml:"open_logstash"`
	OpenLogfile     bool   `toml:"open_logfile"`
	FilePath        string `toml:"file_path"`
	Project         string `toml:"project"`
	Module          string `toml:"module"`
	LogstashAddress string `toml:"logstash_address"`
}

type resources struct {
	FileFolder string `toml:"file_folder"`
}

type config struct {
	Http      http      `toml:"http"`
	Log       logConf   `toml:"log_conf"`
	MysqlDB   mysqlDB   `toml:"mysql_db"`
	Resources resources `toml:"resources"`
}
