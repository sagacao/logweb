package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetString(str string, reply map[string][]byte) string {
	value, err := reply[str]
	if !err {
		return ""
	}
	return string(value)
}

func GetQuery(str string, qurey map[string]string) string {
	if qurey == nil {
		return ""
	}
	value, err := qurey[str]
	if !err {
		return ""
	}
	return value
}

func GetDefaultQuery(str string, qurey map[string]string, defaultstr string) string {
	if qurey == nil {
		return defaultstr
	}
	value, err := qurey[str]
	if !err {
		return defaultstr
	}
	return value
}

func GetIntQuery(iVal int, reply map[int]string) string {
	value, err := reply[iVal]
	if !err {
		return "0"
	}
	return value
}

func GetInterfaceQuery(str string, qurey map[string]interface{}) string {
	if qurey == nil {
		return "undefined"
	}
	value, err := qurey[str]
	if !err {
		return "undefined"
	}
	val, ok := value.(string)
	if ok {
		return val
	}
	return "undefined"
}

func GetInterfaceInt64(str string, qurey map[string]interface{}) int64 {
	if qurey == nil {
		return 0
	}
	value, err := qurey[str]
	if !err {
		return 0
	}
	val, ok := value.(int64)
	if ok {
		return val
	}
	return 0
}

func GetInterfaceUint32(str string, qurey map[string]interface{}) uint32 {
	if qurey == nil {
		return 0
	}
	value, err := qurey[str]
	if !err {
		return 0
	}
	val, ok := value.(uint32)
	if ok {
		return val
	}
	return 0
}

func GetInt(str string, reply map[string][]byte) int {
	value, err1 := reply[str]
	if !err1 {
		return 0
	}
	ret, err2 := strconv.Atoi(string(value))
	if err2 != nil {
		return 0
	}
	return ret
}

func GetInt64(str string, reply map[string][]byte) int64 {
	value, err1 := reply[str]
	if !err1 {
		return 0
	}
	ret, err2 := strconv.ParseInt(string(value), 10, 64)
	if err2 != nil {
		return 0
	}
	return ret
}

func GetFloat32(str string, reply map[string][]byte) float32 {
	value, err1 := reply[str]
	if !err1 {
		return 0
	}
	ret, err2 := strconv.ParseFloat(string(value), 32) //strconv.ParseInt(string(value), 10, 64)
	if err2 != nil {
		return 0
	}
	return float32(ret)
}

func ParseInt(value string) int {
	ret, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return ret
}

func ParseUint32(value string) uint32 {
	ret, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return uint32(ret)
}

func ParseInt64(value string) int64 {
	ret, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	return ret
}

func StringToInt(val string) int {
	return ParseInt(val)
	// intVal, err := strconv.Atoi(val)
	// if err != nil {
	// 	return 0
	// }
	// return intVal
}

func FormatUint32(value uint32) string {
	return strconv.FormatUint(uint64(value), 10)
}

func FormatInt(value int) string {
	return strconv.Itoa(value)
}

func FormatPercent(numerator, denominator float32) string {
	if denominator == 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f", float32(numerator)/float32(denominator)*100)
}

func CatchError(flag string, err error) bool {
	return strings.Contains(err.Error(), flag)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

const (
	y1  = `0123456789`
	y2  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	y3  = `0000000000111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999`
	y4  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	mo1 = `000000000111`
	mo2 = `123456789012`
	d1  = `0000000001111111111222222222233`
	d2  = `1234567890123456789012345678901`
	h1  = `000000000011111111112222`
	h2  = `012345678901234567890123`
	mi1 = `000000000011111111112222222222333333333344444444445555555555`
	mi2 = `012345678901234567890123456789012345678901234567890123456789`
	s1  = `000000000011111111112222222222333333333344444444445555555555`
	s2  = `012345678901234567890123456789012345678901234567890123456789`
)

func FormatTimeHeader(when time.Time) ([]byte, int) {
	y, mo, d := when.Date()
	h, mi, s := when.Clock()
	//len("2006/01/02 15:04:05 ")==20
	var buf [20]byte

	buf[0] = y1[y/1000%10]
	buf[1] = y2[y/100]
	buf[2] = y3[y-y/100*100]
	buf[3] = y4[y-y/100*100]
	buf[4] = '/'
	buf[5] = mo1[mo-1]
	buf[6] = mo2[mo-1]
	buf[7] = '/'
	buf[8] = d1[d-1]
	buf[9] = d2[d-1]
	buf[10] = ' '
	buf[11] = h1[h]
	buf[12] = h2[h]
	buf[13] = ':'
	buf[14] = mi1[mi]
	buf[15] = mi2[mi]
	buf[16] = ':'
	buf[17] = s1[s]
	buf[18] = s2[s]
	buf[19] = ' '

	return buf[0:], d
}

func MapToStruct(m map[string]interface{}, v interface{}) error {
	outrsp, err := json.Marshal(m)
	if err != nil {
		return err
	}

	fmt.Printf("json: %v \n", string(outrsp))

	err = json.Unmarshal(outrsp, &v)
	if err != nil {
		return err
	}
	return nil
}
