package structure

var (
	SystemConf configure
)

type configure struct {
	// Core Configuration - default Framework
	ServicePort         int    `json:"service_port"`
	SecretKey           string `json:"secret_key"`
	ServiceWebserverLog bool   `json:"use_webserver_log"`
	ServiceCronJob      bool   `json:"use_cronjob"`

	// Database Configuration - default Framework
	Database         string `json:"database"`
	DatabaseHost     string `json:"database_host"`
	DatabaseUsername string `json:"database_username"`
	DatabasePassword string `json:"database_password"`
	DatabaseName     string `json:"database_name"`
	UseMigration     bool   `json:"use_migration"`

	// Redis Configuration - default Framework
	ServiceRedis  bool   `json:"use_redis"`
	RedisHost     string `json:"redis_host"`
	RedisPassword string `json:"redis_password"`
	RedisDatabase int    `json:"redis_database"`
	RedisPrefix   string `json:"redis_prefix"`
}
