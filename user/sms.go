package user
import (
	"github.com/melman-go/aliopengo/util"
	. "github.com/melman-go/aliopengo"
)
//发送短信验证码
//http://open.taobao.com/doc2/apiDetail?spm=0.0.0.0.aoyZAl&apiId=25596
func SendVerCode(client *AliHttpClient, param SendSmsParam) (*ResponseEntity, *ErrorResponse, float64) {
	params := map[string]string{}
	params["send_ver_code_request"]=util.JsonEncodeS(param)
	resp := client.PostRequest("taobao.open.sms.sendvercode", params)
	isOk, data, respEntity, errorResponse := client.ParserRespBody("open_sms_sendvercode_response", "result", resp)
	var taskId float64
	if isOk {
		resMap := data.(map[string]interface{})
		if resMap["task_id"]!=nil {
			taskId = resMap["task_id"].(float64)
		}
	}
	return respEntity, errorResponse, taskId
}

func SendVerCodeToMobile(client *AliHttpClient, mobile string) (*ResponseEntity, *ErrorResponse, float64) {
	params := map[string]string{}
	subParams := map[string]interface{}{}
	subParams["mobile"] = mobile
	params["send_ver_code_request"] = util.JsonEncodeS(subParams)
	resp := client.PostRequest("taobao.open.sms.sendvercode", params)
	isOk, data, respEntity, errorResponse := client.ParserRespBody("open_sms_sendvercode_response", "result", resp)
	var taskId float64
	if isOk {
		resMap := data.(map[string]interface{})
		if resMap["task_id"]!=nil {
			taskId = resMap["task_id"].(float64)
		}
	}
	return respEntity, errorResponse, taskId
}


type SendSmsParam struct {
	ExpireTime         int `json:"expire_time"`           //可选 验证码失效时间，单位为秒
	SessionLimit       int `json:"session_limit`          //可选 session级别的发送次数限制
	DeviceLimit        int `json:"device_limit`           //可选 设备级别的发送次数限制
	DeviceLimitInTime  int `json:"device_limit_in_time`   //可选 发送次数限制的时间，单位为秒
	MobileLimit        int `json:"mobile_limit"`          //可选 手机号的次数限制
	SessionLimitInTime int `json:"session_limit_in_time"` //可选 发送次数限制的时间，单位为秒
	ExternalId         string `json:"external_id"`        //可选 外部的id
	MobileLimitInTime  int `json:"mobile_limit_in_time"`  //可选 手机号的次数限制的时间
	TemplateId         int `json:"template_id"`           //可选 模板id
	SignatureId        int `json:"signature_id"`          //可选 签名id
	SessionId          string `json:"session_id"`         //可选 session id
	Demain             string `json:"domain String"`      //可选 场景域，比如登录的验证码不能用于注册
	DeviceId           string `json:"device_id"`          //可选 设备id
	Mobile             string `json:"mobile"`             //必须 手机号
	VerCodeLength      int `json:"ver_code_length"`       //可选 短信内容替换上下文
														  // `json:"context"` Json //可选 短信内容替换上下文
}

//验证短信验证码
//http://open.taobao.com/doc2/apiDetail?spm=0.0.0.0.ZxwtyH&apiId=25597
func CheckVerCode(client *AliHttpClient, subParams CheckSmsParam) (*ResponseEntity, *ErrorResponse) {
	params := map[string]string{}
	params["check_ver_code_request"] = util.JsonEncodeS(subParams)
	resp := client.PostRequest("taobao.open.sms.checkvercode", params)
	_, _, respEntity, errorResponse := client.ParserRespBody("open_sms_checkvercode_response", "result", resp)
	return respEntity, errorResponse
}

type CheckSmsParam struct {
	Domain            string `json:"domain"`           //可选 短信验证码域
	CheckFailLimit    int `json:"check_fail_limit"`    // 可选 最多验证错误几次
	CheckSuccessLimit int `json:"check_success_limit"` // 可选 最多验证成功几次
	VerCode           string `json:"ver_code"`         // 必须 验证码
	Mobile            string `json:"mobile"`           // 必须 手机号
}


func SendMsg(client *AliHttpClient, mobile string, templateId int, context map[string]interface{}) (*ResponseEntity, *ErrorResponse, float64) {
	params := map[string]string{}
	subParams := map[string]interface{}{}
	subParams["mobile"] = mobile
	subParams["template_id"] = templateId
	subParams["context"] = context
	params["send_message_request"] = util.JsonEncodeS(subParams)
	resp := client.PostRequest("taobao.open.sms.sendmsg", params)
	isOk, data, respEntity, errorResponse := client.ParserRespBody("open_sms_sendmsg_response", "result", resp)
	var taskId float64
	if isOk {
		resMap := data.(map[string]interface{})
		if resMap["task_id"]!=nil {
			taskId = resMap["task_id"].(float64)
		}
	}
	return respEntity, errorResponse, taskId
}
