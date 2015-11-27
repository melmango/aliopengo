package user
import (
	. "github.com/melman-go/aliopengo"
	"strconv"
	"github.com/ngaut/log"
)

//open account token 验证
//http://open.taobao.com/doc2/apiDetail?spm=0.0.0.0.lpRM2M&apiId=25270
func TokenValidate(client *AliHttpClient, token string) (*ResponseEntity, *ErrorResponse, *OpenAccount, int64) {
	params := map[string]string{}
	params["param_token"] = token
	resp := client.PostRequest("taobao.open.account.token.validate", params)
	isOk, data, respEntity, errorResponse := client.ParserRespBody("open_account_token_validate_response", "data", resp)
	log.Error(data)
	var openAccountId int64
	var account OpenAccount
	if isOk && data!=nil {
		for k, v := range data.(map[string]interface{}) {
			switch k{
			case "open_account_id":openAccountId = int64(v.(float64))
			case "ext":{
				openAccountMap := v.(map[string]interface{})["open_account"]
				if openAccountMap!=nil {
					account = OpenAccount{}
					account.Parse(openAccountMap.(map[string]interface{}))
				}
			}
			}
		}
	}
	return respEntity, errorResponse, &account, openAccountId
}

//申请免登Open Account Token
//http://open.taobao.com/doc2/apiDetail?spm=0.0.0.0.oBca68&apiId=25271
func TokenApply(client *AliHttpClient, tokenTimeStamp int, openAccountId int, isvAccountId string, uuid string, loginStateExpireIn int) (*ResponseEntity, *ErrorResponse, string) {
	params := map[string]string{}
	if tokenTimeStamp>0 {
		params["token_timestamp"]= strconv.Itoa(tokenTimeStamp)
	}
	if openAccountId>0 {
		params["open_account_id"]= strconv.Itoa(openAccountId)
	}
	if isvAccountId!="" {
		params["isv_account_id"]= isvAccountId
	}
	if uuid!="" {
		params["uuid"]= uuid
	}
	if loginStateExpireIn>0 {
		params["login_state_expire_in"]= strconv.Itoa(loginStateExpireIn)
	}
	resp := client.PostRequest("taobao.open.account.token.apply", params)
	_, data, respEntity, errorResponse := client.ParserRespBody("open_account_token_apply_response", "data", resp)
	return respEntity, errorResponse, data.(string)
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
	LogonPwdIntensity  int64  `json:"login_pwd_intensity"`
	Id                 int64  `json:"id"`
	Type               int64  `json:"type"`
	Status             int64  `json:"status"`
	Version            int64  `json:"version"`
	LoginPwdEncryption int64  `json:"login_pwd_encryption"`
	Gender             int64  `json:"gender"`
	Name               string `json:"name"`
	Birthday           string `json:"birthday"`
	WangWang           string `json:"wangwang"`
	Weixin             string `json:"weixin"`
	OauthPlatform      int64  `json:"oauth_plateform"`
}



func (this *OpenAccount) Parse(res map[string]interface{}) {
	for k, v := range res {
		switch k{
		case "login_id":               this.LoginId = v.(string)
		case "create_device_id":this.CreateDeviceId = v.(string)
		case "alipay_id":this.AlipayId = v.(string)
		case "locale":this.Locale = v.(string)
		case "bank_card_no":this.BandCardNo = v.(string)
		case "isv_account_id":this.IsvAccountId = v.(string)
		case "email":this.Email = v.(string)
		case "avatar_url":this.AvatarUrl = v.(string)
		case "bank_card_owner_name":this.BandCardOwnerName = v.(string)
		case "display_name":this.DisplayName = v.(string)
		case "login_pwd_salt":this.LoginPwdSalt = v.(string)
		case "login_pwd":this.LiginPwd = v.(string)
		case "open_id":this.OpenId = v.(string)
		case "mobile":this.Mobile = v.(string)
		case "create_location":this.CreateLocation = v.(string)
		case "ext_infos":this.ExtInfos = v.(string)
		case "login_pwd_intensity":this.LogonPwdIntensity = int64(v.(float64))
		case "id":this.Id = int64(v.(float64))
		case "type":this.Type = int64(v.(float64))
		case "status":this.Status = int64(v.(float64))
		case "version":this.Version = int64(v.(float64))
		case "login_pwd_encryption":this.LoginPwdEncryption = int64(v.(float64))
		case "gender":this.Gender = int64(v.(float64))
		case "name":this.Name = v.(string)
		case "birthday":this.Birthday = v.(string)
		case "wangwang":this.WangWang = v.(string)
		case "weixin":this.Weixin = v.(string)
		case "oauth_plateform":this.OauthPlatform = int64(v.(float64))
		}
	}
}
