package global

type GlobalConfig struct {
	ServerConfig ServerConfig `mapstructure:"server"`
	LoggerConfig LoggerConfig `mapstructure:"logger"`
	DBConfig     DBConfig     `mapstructure:"db"`
}

type ServerConfig struct {
	Port   int    `mapstructure:"port"`
	Salt   string `mapstructure:"salt"`
	Router string `mapstructure:"router"`
}

type LoggerConfig struct {
	FilePath   string `mapstructure:"file_path"`
	FileName   string `mapstructure:"file_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type DBConfig struct {
	Uri      string `mapstructure:"uri"`
	Username string `mapstructure:"username"`
	Pwd      string `mapstructure:"pwd"`
	DBName   string `mapstructure:"dbname"`
}

type OrganizationConfig struct {
	JoinUrl string `mapstructure:"join_url"`
}
