package user
import (
	"testing"
	"github.com/melman-go/aliopengo"
	"github.com/ngaut/log"
	"github.com/melman-go/aliopengo/util"
)

const (
//	URL = "http://gw.api.tbsandbox.com/router/rest"
	URL = "http://gw.api.taobao.com/router/rest"
	APP_KEY = ""
	APP_SECRET = ""
)

func createClient() *aliopengo.AliHttpClient {
	return aliopengo.NewAliHttpClient(URL, APP_KEY, APP_SECRET, "")
}

func TestTokenValidate(t *testing.T) {
	client := createClient()
	_,_,account,id := TokenValidate(client, "5f+pfdl0AqOrQxHnW11Y2l584ER+xZF5mo/x4IepP3ABxMozXYXhFrZtVZK5NKQvPE+iLGoHXIVbGaS6uXKeUaJPiA88PfOIsIu+T1PBL9L+U2hjgXVpnAAmxkxitp3Ep44gDt4CIxcIEECfXS/f16fHwXNKpX1N089oJphVQtUig9yRT2/hgMQ7ihcJu+CVHWKYVIEs3WkOo4UVfwAj/0Kd328XutVgudZfTgRCYAmjyWFUhZr9wriCpZKgFih81ADoSND6wQEBz5GwOBHevH841Vk2rQH3OPKuxGZLe0vrIrivHieCu0LBGOLG9QeN4Pzk/YAbZrmWFiMHznbIgWRbFdXgPjxzlX0FCMnWakTuCKCI2djcyk3Fx5vKxSuRLFRh2JWGfKlqoJtZlWVDeEYEw61tfdUnyFHkq2uqLg8=")
	if account!=nil{
		log.Error(util.JsonEncodeS(account))
		log.Error(id)
	}
}

//func TestCheckText(t *testing.T){
//	client := createClient()
//	CheckText(client,"测试","11826")
//}

//func TestSendSms(t *testing.T) {
//	client := createClient()
//	SendVerCodeToMobile(client, "15010035751")
//}

//func TestCheckSms(t *testing.T) {
//	client := createClient()
//	params := CheckSmsParam{
//		VerCode:"1234",
//		Mobile:"15010035751",
//	}
//	CheckVerCode(client, params)
//}

//func TestSendMsg(t *testing.T){
//	client := createClient()
//	paramMap := map[string]interface{}{}
//	paramMap["code"] = "三胖"
//	SendMsg(client,"15010035751",952,paramMap)
//}




