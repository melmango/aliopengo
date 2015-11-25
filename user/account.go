package user
import (
	"github.com/melman-go/aliopengo/util"
	. "github.com/melman-go/aliopengo"
	"net/url"
	"strconv"
)

//open account token 验证
//http://open.taobao.com/doc2/apiDetail?spm=0.0.0.0.lpRM2M&apiId=25270
func TokenVAlidate(client *AliHttpClient, token string) (*OpenAccount, *ResponseEntity, *ErrorResponse) {
	values := *url.Values{
		"param_token":token,
	}
	resp := client.SendRequest("taobao.open.account.token.validate", values)
	isOk, data, respEntity, errorResponse := client.ParserRespBody("open_account_token_validate_response","data", resp)
	account := *OpenAccount{}
	if isOk {
		util.JsonDecodeS(data, account)
	}
	return account, respEntity, errorResponse
}

//申请免登Open Account Token
//http://open.taobao.com/doc2/apiDetail?spm=0.0.0.0.oBca68&apiId=25271
func TokenApply(client *AliHttpClient, tokenTimeStamp int, openAccountId int, isvAccountId string, uuid string, loginStateExpireIn int) (string, *ResponseEntity, *ErrorResponse) {
	values := *url.Values{}
	if tokenTimeStamp>0 {
		values.Set("token_timestamp", strconv.Itoa(tokenTimeStamp))
	}
	if openAccountId>0 {
		values.Set("open_account_id", strconv.Itoa(openAccountId))
	}
	if isvAccountId!="" {
		values.Set("isv_account_id", strconv.Itoa(isvAccountId))
	}
	if uuid!="" {
		values.Set("uuid", strconv.Itoa(uuid))
	}
	if loginStateExpireIn>0 {
		values.Set("login_state_expire_in", strconv.Itoa(loginStateExpireIn))
	}
	resp := client.SendRequest("taobao.open.account.token.apply", values)
	_, data, respEntity, errorResponse := client.ParserRespBody("open_account_token_apply_response", "data",resp)
	return data, respEntity, errorResponse
}


type OpenAccount struct {
	LoginId            string `json:"login_id"`
	CreateDeviceId     string `json:"create_device_id"`
	AlipayId           string `json:"alipay_id"`
	Locale             string `json:"locale"`
	BandCardNo         string `json:"bank_card_no"`
	IsvAccountId       string `json:"isv_account_id"`
	Email              string `json:"email"`
	AvatarUrl          string `json:"avatar_url"`
	BandCardOwnerName  string `json:"bank_card_owner_name"`
	DisplayName        string `json:"display_name"`
	LoginPwdSalt       string `json:"login_pwd_salt"`
	LiginPwd           string `json:"login_pwd"`
	OpenId             string `json:"open_id"`
	Mobile             string `json:"mobile"`
	CreateLocation     string `json:"create_location"`
	ExtInfos           string `json:"ext_infos"`
	LogonPwdIntensity  int32 `json:"login_pwd_intensity"`
	Id                 int32 `json:"id"`
	Type               int32 `json:"type"`
	Status             int32 `json:"status"`
	Version            int32 `json:"version"`
	LoginPwdEncryption int32 `json:"login_pwd_encryption"`
	Gender             int32 `json:"gender"`
	Name               string `json:"name"`
	Birthday           string `json:"birthday"`
	WangWang           string `json:"wangwang"`
	Weixin             string `json:"weixin"`
	OauthPlatform      int32 `json:"oauth_plateform"`
}
