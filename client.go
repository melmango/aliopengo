package gobaichuan
import (
	"net/http"
	"net/url"
	"github.com/melman-go/gobaichuan/util"
	"fiwbee/helpers"
	"bytes"
)

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

func (this *AliHttpClient) ParserRespBody(resp *http.Response) map[string]string{

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
