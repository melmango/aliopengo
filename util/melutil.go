package util
import (
	"encoding/json"
	"time"
)

func JsonDecodeS(jstr string, obj interface{}) {
	json.Unmarshal([]byte(jstr), obj)
}

func JsonDecodeB(value []byte,obj interface{}){
	json.Unmarshal(value,obj);
}

func JsonEncodeS(obj interface{}) string {
	res, _ := json.Marshal(obj)
	return string(res)
}

func JsonEncodeB(obj interface{}) []byte {
	res, _ := json.Marshal(obj)
	return res
}

func GetCurrentTimeStamp() string{
	return time.Now().Format("2006-01-02 15:04:05")
}