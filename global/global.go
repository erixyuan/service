package global

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
)

const (
	envIsDev  = "IS_DEV"
	envIsTest = "IS_TEST"
	envIsProd = "IS_PROD"
	envIsFake = "IS_FAKE"
)

func getEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}
func InitViperConfig(workDir string, config interface{}) {
	isProd := getEnvInfo(envIsProd)
	isTest := getEnvInfo(envIsTest)
	isDev := getEnvInfo(envIsDev)
	isFake := getEnvInfo(envIsFake)
	fmt.Println("Load config path:", workDir)
	var configFileName string
	if isProd {
		// 生产环境
		configFileName = path.Join(workDir, "application.prod.yml")
		fmt.Println("now env is product")
	} else if isTest {
		// 测试环境
		configFileName = path.Join(workDir, "application.test.yml")
		fmt.Println("now env is test")
	} else if isDev {
		// 开发环境
		configFileName = path.Join(workDir, "application.dev.yml")
		fmt.Println("now env is dev")
	} else if isFake {
		// 开发环境
		configFileName = path.Join(workDir, "application.fake.yml")
		fmt.Println("now env is fake")
	} else {
		// 本地开发环境
		// 这里要改为./ 不然本地起不来
		configFileName = path.Join("./", "application.dev.yml")
		fmt.Println("now local env is developer")
	}

	fmt.Println("Config filename:", configFileName)
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(config); err != nil {
		panic(err)
	}
	if _, err := fmt.Printf("Load config content : %v", config); err != nil {
		panic(err)
	}
}

func CheckIsDev() bool {
	return getEnvInfo(envIsDev)
}

func CheckIsTest() bool {
	return getEnvInfo(envIsTest)
}

func CheckIsProd() bool {
	return getEnvInfo(envIsProd)
}
