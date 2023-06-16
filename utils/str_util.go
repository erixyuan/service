package utils

import (
	"fmt"
	"github.com/gookit/goutil/arrutil"
	"github.com/spf13/cast"
	"math/rand"
	"service/global"
	"strconv"
	"strings"
	"time"
)

func SliceToString(ints []int) string {
	var slices []string
	for _, i := range ints {
		slices = append(slices, strconv.Itoa(i))
	}
	return strings.Join(slices, ",")
}

func StringToInts(s string) []int {
	if ints, err := arrutil.StringsToInts(strings.Split(s, ",")); err != nil {
		global.GetLogger().Errorf("StringToInts error: %v", err)
		return []int{}
	} else {
		return ints
	}
}

// AmountCentToString @Description: 金额分转换成元,字符串
func AmountCentToString(amount int) string {
	price := float64(amount) / float64(100)
	return fmt.Sprintf("%.2f", price)
}

// GenerateNumberCode @Description: 生成指定位数的数字
func GenerateNumberCode(length int) string {
	str := "1"
	for i := 0; i < length; i++ {
		str += "0"
	}
	return fmt.Sprintf("%0"+strconv.Itoa(length)+"v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(cast.ToInt32(str)))
}

func HasDuplicates(arr []string) bool {
	//创建一个map用于存储每个字符串的出现次数
	frequencyMap := make(map[string]int)

	//对于数组中的每个字符串
	for _, str := range arr {
		//如果字符串已经在map中
		if _, ok := frequencyMap[str]; ok {
			//有重复元素
			return true
		} else {
			//否则将该字符串放入map中
			frequencyMap[str] = 1
		}
	}
	//没有重复元素
	return false
}
