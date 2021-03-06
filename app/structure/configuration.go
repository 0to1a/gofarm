package structure

var (
	SystemConf configure
)

type configure struct {
	ServicePort    int    `json:"service_port"`
	SecretKey      string `json:"secret_key"`
	ServiceLog     bool   `json:"use_log"`
	ServiceMonitor bool   `json:"use_monitor"`
	ServiceCronJob bool   `json:"use_cronjob"`

	Database         string `json:"database"`
	DatabaseHost     string `json:"database_host"`
	DatabaseUsername string `json:"database_username"`
	DatabasePassword string `json:"database_password"`
	DatabaseName     string `json:"database_name"`
	UseMigration     bool   `json:"use_migration"`

	ServiceRedis  bool   `json:"use_redis"`
	RedisHost     string `json:"redis_host"`
	RedisPassword string `json:"redis_password"`
	RedisDatabase int    `json:"redis_database"`
}
