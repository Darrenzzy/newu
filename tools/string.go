package tools

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func StringToBool(e string) (bool, error) {
	return strconv.ParseBool(e)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05.9999")
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func FormatTimeStr(timeStr string) (string, error) {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02T15:04:05.000Z", timeStr, loc)
	return theTime.Format("2006/01/02 15:04:05"), err
}
func FormatTimeStrSimple(timeStr string) (string, error) {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02", timeStr, loc)
	return theTime.Format("2006/01/02 15:04:05"), err
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func GetBodyString(c *gin.Context) (string, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return string(body), nil
	} else {
		return "", err
	}
}

func JsonStrToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	} else {
		return nil, err
	}
}

func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

func EncodeToString(max int) string {
	b := make([]byte, max)
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
