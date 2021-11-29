package config

var (
	MysqlC  *MysqlConfig
	ServerC *ServerConfig
)

type MysqlConfig struct {
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int64
	MaxOpenConns int64
}

type ServerConfig struct {
	RunMode             string
	HeartBeatServerPort int64
}

func (c *Config) ReadSection(k string, v interface{}) error {
	err := c.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
