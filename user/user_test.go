package user
import (
	"testing"
	"github.com/melman-go/aliopengo"
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
	TokenValidate(client, "12345")
}

//func TestSendSms(t *testing.T) {
//	client := createClient()
//	SendVerCodeToMobile(client, "15010035751")
//}

func TestCheckSms(t *testing.T) {
	client := createClient()
	params := CheckSmsParam{
		VerCode:"1234",
		Mobile:"15010035751",
	}
	CheckVerCode(client, params)
}

//func TestSendMsg(t *testing.T){
//	client := createClient()
//	paramMap := map[string]interface{}{}
//	paramMap["code"] = "三胖"
//	SendMsg(client,"15010035751",952,paramMap)
//}




