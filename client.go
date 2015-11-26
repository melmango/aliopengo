package aliopengo
import (
	"net/http"
	"net/url"
	"github.com/melman-go/aliopengo/util"
	"fiwbee/helpers"
	"bytes"
	"io/ioutil"
	"github.com/ngaut/log"
	"fmt"
	"sort"
	"strings"
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
	Code       float64 `json:"code"`
}

func (this *ResponseEntity) Parse(res map[string]interface{}) interface{} {
	var data interface{}
	for k, v := range res {
		switch k{
		case "code":this.Code = v.(float64)
		case "message": this.Message = v.(string)
		case "successful":this.Successful = v.(bool)
		case "data":data = v
		case "datas": data = v
		}
	}
	return data
}

type ErrorResponse struct {
	Code      float64 `json:"code"`
	Msg       string `json:"msg"`
	SubCode   string `json:"sub_code"`
	SubMsg    string `json:"sub_msg"`
	RequestId string `json:"request_id"`
}

func (this *ErrorResponse) Parse(res map[string]interface{}) {
	for k, v := range res {
		switch k{
		case "code":this.Code = v.(float64)
		case "msg":this.Msg = v.(string)
		case "sub_code":this.SubCode = v.(string)
		case "sub_msg":this.SubMsg = v.(string)
		case "request_id": this.RequestId = v.(string)
		}
	}
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

func (this *AliHttpClient) String() string {
	return fmt.Sprintf("CLIENT: url:%s,key:%s,secret:%s\n", this.baseUrl, this.appKey, this.appSecret)
}

func (this *AliHttpClient)SendRequest(method string, values url.Values) *http.Response {
	this.bindDefaultParams(method, values)
	log.Error(values)
	resp, err := this.client.PostForm(this.baseUrl, values)
	CheckHttpError(err)
	return resp
}

func CheckHttpError(err error) {

}

func (this *AliHttpClient)bindDefaultParams(method string, values url.Values) {
	values.Set("method", method)
	values.Set("app_key", this.appKey)
	if this.session!="" {
		values.Set("session", this.session)
	}
	values.Set("timestamp", util.GetCurrentTimeStamp())
	values.Set("format", this.format)
	values.Set("v", this.ver)
	if this.partnerId!="" {
		values.Set("partner_id", this.partnerId)
	}
	if this.simplify {
		values.Set("simplify", "true")
	}else {
		values.Set("simplify", "false")
	}
	values.Set("sign_method", this.signMethod)
	values.Set("sign", this.CalcSign(values))
}


func (this *AliHttpClient) CalcSign(values url.Values) string {
	var keys []string
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	buf := new(bytes.Buffer)
	buf.WriteString(this.appSecret)
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(values.Get(k))
	}
	buf.WriteString(this.appSecret)
	sign := strings.ToUpper(helpers.EncodeMd5(buf.String()))
	return sign
}

func (this *AliHttpClient) ParserRespBody(method string, entityKey string, resp *http.Response) (bool, interface{}, *ResponseEntity, *ErrorResponse) {
	if resp.StatusCode!=http.StatusOK {
		return false, nil, nil, nil
	}else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		CheckHttpError(err)
		resMap := map[string]interface{}{}
		util.JsonDecodeB(body, &resMap)
		errorMap := resMap["error_response"]
		if errorMap!=nil {
			errorResponse := &ErrorResponse{}
			errorResponse.Parse(errorMap.(map[string]interface{}))
			log.Error(errorResponse)
			return false, nil, nil, errorResponse
		}
		successMap := resMap[method]
		var data interface{}
		var respEntity ResponseEntity
		if successMap!=nil {
			entityMap := successMap.(map[string]interface{})[entityKey]
			if entityMap!=nil {
				respEntity = ResponseEntity{}
				data = respEntity.Parse(entityMap.(map[string]interface{}))
				log.Error(respEntity)
			}
		}
		return true, data, &respEntity, nil
	}

}

func NewAliHttpClient(baseUrl string, appkey string, appSecret string, partnerId string) *AliHttpClient {
	client := AliHttpClient{
		client:http.DefaultClient,
		baseUrl:baseUrl,
		appKey: appkey,
		appSecret: appSecret,
		format:"json",
		ver:"2.0",
		partnerId:partnerId,
		simplify:false,
		signMethod:"md5",
	}
	return &client
}
