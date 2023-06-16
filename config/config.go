package config

var (
	ServerGlobalConfig *GlobalConfig = &GlobalConfig{}
)

type GlobalConfig struct {
	ServerConfig       ServerConfig       `mapstructure:"server"`
	LoggerConfig       LoggerConfig       `mapstructure:"logger"`
	DBConfig           DBConfig           `mapstructure:"db"`
	CoreConfig         CoreConfig         `mapstructure:"core"`
	AliyunOssConfig    AliyunOssConfig    `mapstructure:"aliyun_oss"`
	AliyunSmsConfig    AliyunSmsConfig    `mapstructure:"aliyun_sms"`
	PaymentCallback    PaymentCallback    `mapstructure:"payment_callback"`
	RedisConfig        RedisConfig        `mapstructure:"redis"`
	JwtConfig          JwtConfig          `mapstructure:"jwt"`
	WeChatConfig       WeChatConfig       `mapstructure:"wechat"`
	AliPayConfig       AliPayConfig       `mapstructure:"alipay"`
	OrganizationConfig OrganizationConfig `mapstructure:"organization"`
	WeilaConfig        WeilaConfig        `mapstructure:"lanpice"`
	TencentCosConfig   TencentCosConfig   `mapstructure:"tencent_cos"`
}

type ServerConfig struct {
	Port    int    `mapstructure:"port"`
	Salt    string `mapstructure:"salt"`
	Router  string `mapstructure:"router"`
	Crontab int    `mapstructure:"crontab"`
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

type UploadConfig struct {
	FilePath string `mapstructure:"file_path"`
}

type AliyunOssConfig struct {
	Region           string `mapstructure:"region"`
	EndPoint         string `mapstructure:"endpoint"`
	AccessKeyId      string `mapstructure:"access_key_id"`
	AccessSecret     string `mapstructure:"access_secret"`
	BucketName       string `mapstructure:"bucket_name"`
	PublicBucketName string `mapstructure:"public_bucket_name"`
	OssUrl           string `mapstructure:"oss_url"`
}

type AliyunSmsConfig struct {
	AccessKeyId      string `mapstructure:"access_key_id"`
	AccessKeySecret  string `mapstructure:"access_key_secret"`
	SmsSignName      string `mapstructure:"sms_sign_name"`
	TemplateCode     string `mapstructure:"template_code"`
	TemplateParamKey string `mapstructure:"template_param_key"`
}

type PaymentCallback struct {
	Alipay string `mapstructure:"alipay"`
	Wechat string `mapstructure:"wechat"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JwtConfig struct {
	JwtSecret string   `mapstructure:"jwt_secret"`
	JwtExcuse []string `mapstructure:"jwt_excuse"`
}

type WeChatConfig struct {
	Pay         WeChatPayConfig         `mapstructure:"pay"`
	Miniprogram WeChatMiniProgramConfig `mapstructure:"miniprogram"`
	WeChatApp   WeChatAppConfig         `mapstructure:"app"`
}

type WeChatPayConfig struct {
	MchID                string `mapstructure:"mch_id"`
	MerchantSerialNumber string `mapstructure:"merchant_serial_number"`
	ApiV3Key             string `mapstructure:"api_v3_key"`
}

type WeChatMiniProgramConfig struct {
	AppId     string `mapstructure:"app_id"`
	AppSecret string `mapstructure:"app_secret"`
}

type WeChatAppConfig struct {
	AppId     string `mapstructure:"app_id"`
	AppSecret string `mapstructure:"app_secret"`
}

type AliPayConfig struct {
	AppId           string `mapstructure:"app_id"`
	AppPrivateKey   string `mapstructure:"app_private_key"`
	AlipayPublicKey string `mapstructure:"alipay_public_key"`
}

type CoreConfig struct {
	Host string `mapstructure:"host"`
}

type OrganizationConfig struct {
	JoinUrl   string `mapstructure:"join_url"`
	ReviewUrl string `mapstructure:"review_url"`
}

type WeilaConfig struct {
	AppId  string `mapstructure:"app_id"`
	AppKey string `mapstructure:"app_key"`
}

type TencentCosConfig struct {
	SecretId  string `mapstructure:"secret_id"`
	SecretKey string `mapstructure:"secret_key"`
	Buket     string `mapstructure:"buket"`
	Region    string `mapstructure:"region"`
}
