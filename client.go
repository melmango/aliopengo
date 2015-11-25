package aliopengo
import (
	"net/http"
	"net/url"
	"github.com/melman-go/aliopengo/util"
	"fiwbee/helpers"
	"bytes"
	"io/ioutil"
	"strconv"
)

const (
	ERROR_CODE_SERVER_ERROR = 10000//系统异常了，无法处理
	ERROR_CODE_ILLEGAL_PARAM = 10001//参数错误
	ERROR_CODE_RECORD_ALREADY_EXIST = 20001//记录已经存在，判断依据是根据mobile、email、login_id、isv_account_id、open_id
	ERROR_CODE_RECORD_NOT_EXIST = 20002//记录不存在，如删除的时候传了错误的id或isv_account_id
	ERROR_CODE_DOMAIN_MISMATCH = 20004//数据域检查错误，操作了不属于自己域的数据)
	ERROR_CODE_UPDATE_FAIL = 20005//更新时DB异常
	ERROR_CODE_INSERT_FAIL = 20006//插入时DB异常
	ERROR_CODE_INSERT_INDEX_FAIL = 20007//写索引DB异常
	ERROR_CODE_INVALID_TOKEN = 20008//Token错误或者验证时间超过了1分钟
)


type ResponseEntity struct {
	Message    string `json:"message"`
	Successful bool `json:"successful"`
	Code       int32 `json:"code"`
}

type ErrorResponse struct {
	Code    int32 `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

type AliHttpClient struct {
	client     *http.Client
	baseUrl    string
	appKey     string
	appSecret  string
	session    string
	format     string
	ver        string
	partnerId  string
	simplify   bool
	signMethod string
}

func (this *AliHttpClient)SendRequest(method string, values *url.Values) *http.Response {
	this.bindDefaultParams(method, values)
	resp, err := this.client.PostForm(this.baseUrl, values)
	CheckHttpError(err)
	return resp
}

func CheckHttpError(err error) {

}

func (this *AliHttpClient)bindDefaultParams(method string, values *url.Values) {
	values["method"] = method
	values["app_key"] = this.appKey
	if this.appKey!=nil {
		values["session"] = this.session
	}
	values["timestamp"] = util.GetCurrentTimeStamp()
	values["format"] = this.format
	values["v"] = this.ver
	if this.partnerId!=nil {
		values["partner_id"] = this.partnerId
	}
	if this.simplify {
		values["simplify"] = "true"
	}else {
		values["simplify"] = "false"
	}
	values["sign_method"] = this.signMethod
	values["sign"] = this.CalcSign(values)
}


func (this *AliHttpClient) CalcSign(values *url.Values) string {
	//TODO finish sign calc logic
	var keys []string
	for k := range values {
		keys = append(keys, k)
	}
	buf := new(bytes.Buffer)
	buf.WriteString(this.appSecret)
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(values[k])
	}
	buf.WriteString(this.appKey)
	sign := helpers.EncodeMd5(buf.String())
	return sign
}

func (this *AliHttpClient) ParserRespBody(method string,entityKey string,resp *http.Response) (bool, string, *ResponseEntity, *ErrorResponse) {
	if resp.StatusCode!=http.StatusOK {
		return false, "", nil, nil
	}else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		CheckHttpError(err)
		resMap := map[string]string{}
		util.JsonDecodeB(body, resMap)
		var errorResponse *ErrorResponse
		var respEntity *ResponseEntity
		errorStr := resMap["error_response"]
		successStr := resMap[method]
		var data string
		if errorStr!=nil {
			util.JsonDecodeS(errorStr, &errorResponse)
		}
		if successStr!=nil {
			var dataMap map[string]string
			util.JsonDecodeS(successStr, dataMap)
			data= dataMap["data"]
			respEntity = *ResponseEntity{
				Message:dataMap["message"],
				Successful:dataMap["code"]=="true",
				Code:strconv.Atoi(dataMap["code"]),
			}
		}
		return true, data, respEntity, errorResponse
	}
}

func NewAliHttpClient(baseUrl string, appkey string, appSecret string, partnerId string) *AliHttpClient {
	client := &AliHttpClient{
		client:http.DefaultClient,
		appkey: appkey,
		appSecret: appSecret,
		format:"json",
		ver:"2.0",
		partnerId:partnerId,
		simplify:false,
		signMethod:"md5",
	}
	return client
}
