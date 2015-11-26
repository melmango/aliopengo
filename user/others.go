package user

import (
	. "github.com/melman-go/aliopengo"
)

func CheckText(client *AliHttpClient, content string, appId string) (*ResponseEntity, *ErrorResponse, CheckTextResult) {
	params := map[string]string{}
	params["content"] = content
	params["appId"] = appId
	resp := client.PostRequest("taobao.user.textcheck.get", params)
	isOk, data, respEntity, errorResponse := client.ParserRespBody("user_textcheck_get_response", "result", resp)
	var result CheckTextResult
	if isOk && data!=nil {
		result = CheckTextResult{}
		for k, v := range data.(map[string]interface{}) {
			switch k{
			case "target":result.Target = v.(string)
			case "original":result.Original = v.(string)
			case "isIllegal":result.IsIllegal = v.(bool)
			}
		}
	}
	return respEntity, errorResponse, result
}

type CheckTextResult struct {
	Target    string
	Original  string
	IsIllegal bool
}