package components

import (
	"reflect"
	"runtime"
	"service/cache"
	"service/dto"
	"service/enum/err_enum"
	"service/global"
	"strings"
	"time"
)

func GetFunctionName(i interface{}) string {
	// 获取函数名称
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	if fn != "" {
		arr := strings.Split(fn, "/")
		return arr[len(arr)-1]
	}
	return ""
}

func WithLock(baseParam dto.BaseParam, lKey string, fn func(baseParam dto.BaseParam) *err_enum.ErrorObj, other ...any) *err_enum.ErrorObj {
	functionName := GetFunctionName(fn)
	lockKey := "lock-" + lKey
	lockValue := cache.RedisClient.Get(lockKey)
	if cache.RedisClient.IsExist(lockKey) {
		global.GetLogger().InfoByTrace(&baseParam, "%s 存在锁  key:%s, value:%v", functionName, lockKey, lockValue)
		return err_enum.LockErr
	} else {
		global.GetLogger().InfoByTrace(&baseParam, "%s 锁不存在，开始写入锁标记 %s", functionName, lockKey)
		cache.RedisClient.Set(lockKey, 1, 30*time.Second)
	}
	defer func() {
		if err := recover(); err != nil {
			global.GetLogger().ErrorByTrace(&baseParam, "%s 执行异常，解除锁 %s err: %v", functionName, lockKey, err)
			cache.RedisClient.Delete(lockKey)
			panic(err)
		}
	}()
	defer func(baseParam dto.BaseParam, lockKey string) {
		global.GetLogger().InfoByTrace(&baseParam, "%s 执行完成，解除锁 %s", functionName, lockKey)
		cache.RedisClient.Delete(lockKey)
	}(baseParam, lockKey)
	if err := fn(baseParam); err != nil {
		global.GetLogger().WarnByTrace(&baseParam, "%s 执行异常 err: %+v", functionName, err)
		return err
	}
	return nil
}
