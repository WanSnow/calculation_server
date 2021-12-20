package config

var (
	MysqlC  *MysqlConfig
	ServerC *ServerConfig
	NsqC    *NsqConfig
	RedisC  *RedisConfig
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
	StartGameServerPort int64
	AcceptServerPort    int64
}

type NsqConfig struct {
	NsqdUrl       string
	NsqLookupdUrl string
}

type RedisConfig struct {
	Addr         string
	PoolSize     int
	MinIdleConns int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	PoolTimeout  int
	IdleTimeout  int
}

func (c *Config) ReadSection(k string, v interface{}) error {
	err := c.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
