package sys

/*
Make initial variable from env
*/

type SysEnv struct {
	AppEnv      string `mapstructure:"APP_ENV"`
	AppName     string `mapstructure:"APP_TITLE"`
	AppVersion  string `mapstructure:"APP_VERSION"`
	AppJWTSalt  string `mapstructure:"APP_JWT_SALT"`
	AppTimeZone string `mapstructure:"APP_TIMEZONE"`
	AppPort     string `mapstructure:"APP_PORT"`
	AppCors     bool   `mapstructure:"APP_CORS"`

	DbHost            string `mapstructure:"DB_HOST"`
	DbPort            string `mapstructure:"DB_PORT"`
	DbName            string `mapstructure:"DB_NAME"`
	DbUSer            string `mapstructure:"DB_USER"`
	DbPassword        string `mapstructure:"DB_PASSWORD"`
	DbMaxOpenConn     int    `mapstructure:"DB_MAX_OPEN_CONN"`
	DbMaxConnLifeTime int    `mapstructure:"DB_MAX_CONN_LIFETIME"`
}
