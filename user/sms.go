package user
import (
	"github.com/melman-go/aliopengo/util"
	. "github.com/melman-go/aliopengo"
	"net/url"
)
//发送短信验证码
//http://open.taobao.com/doc2/apiDetail?spm=0.0.0.0.aoyZAl&apiId=25596
func SendVerCode(client *AliHttpClient, token string) (*OpenAccount, *ResponseEntity, *ErrorResponse) {
	values := *url.Values{
//		"param_token":token,
	}
	resp := client.SendRequest("taobao.open.sms.sendvercode", values)
	isOk, data, respEntity, errorResponse := client.ParserRespBody("open_sms_sendvercode_response", "result", resp)
	account := *OpenAccount{}
	if isOk {
		util.JsonDecodeS(data, account)
	}
	return account, respEntity, errorResponse
}
