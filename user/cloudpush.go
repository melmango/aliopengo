package user
import (
	. "github.com/melman-go/aliopengo"
)

const (
	PUSH_DEVICE_TYPE_IOS = 1
	PUSH_DEVICE_TYPE_ANDROID = 2
)


/**
target 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
targetValue 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
deviceType 1-iOS 2-Android
 */
func CloudPushMessage(client *AliHttpClient, deviceType int, body string, target string, targetValue string) (*ErrorResponse, *SendPushResult) {
	params := map[string]string{}
	params["body"] = body
	params["target"] = target
	params["target_value"] = targetValue
	var method, respKey string
	switch deviceType{
	case PUSH_DEVICE_TYPE_IOS:{
		method = "taobao.cloudpush.message.ios"
		respKey = "cloudpush_message_ios_response"
	}
	case PUSH_DEVICE_TYPE_ANDROID:{
		method = "taobao.cloudpush.message.android"
		respKey = "cloudpush_message_android_response"
	}
	}
	resp := client.PostRequest(method, params)
	isOk, data, errorResponse := client.ParserRespBodyWithoutData(respKey, resp)
	var result SendPushResult
	if isOk && data!=nil {
		result = SendPushResult{}
		result.Parse(data.(map[string]interface{}))
	}
	return &result, errorResponse
}

func CloudPushNoticeAndroid(client *AliHttpClient, summary string, target string, target_value string, title string) {
	params := map[string]string{}
	params["summary"] = summary
	params["target"] = target
	params["target_value"] = target_value
	params["title"] = title
	resp := client.PostRequest("taobao.cloudpush.notice.android", params)
	isOk, data, errorResponse := client.ParserRespBodyWithoutData("cloudpush_notice_android_response", resp)
	var result SendPushResult
	if isOk && data!=nil {
		result = SendPushResult{}
		result.Parse(data.(map[string]interface{}))
	}
	return &result, errorResponse
}

func CloudPushNoticeiOS(client *AliHttpClient, summary string, target string, target_value string, env string, ext string) (*ErrorResponse, *SendPushResult) {
	params := map[string]string{}
	params["summary"] = summary
	params["target"] = target
	params["target_value"] = target_value
	params["env"] = env
	params["ext"] = ext
	resp := client.PostRequest("taobao.cloudpush.notice.ios", params)
	isOk, data, errorResponse := client.ParserRespBodyWithoutData("cloudpush_notice_ios_response", resp)
	var result SendPushResult
	if isOk && data!=nil {
		result = SendPushResult{}
		result.Parse(data.(map[string]interface{}))
	}
	return &result, errorResponse
}

//TODO push 高级接口
//func CloadPushAdvance() {
//
//}

type SendPushResult struct {
	RequestErrorCode float64
	RequestErrorMsg  string
	IsSuccess        bool
}

func (this *SendPushResult) Parse(data map[string]interface{}) {
	for k, v := range data {
		switch k{
		case "request_error_code":this.RequestErrorCode = v.(float64)
		case "request_error_msg":this.RequestErrorMsg = v.(string)
		case "is_success":this.IsSuccess = v.(bool)
		}
	}
}