package account
import (
	"github.com/melman-go/aliopengo/util"
	. "github.com/melman-go/aliopengo"
	"net/url"
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

//open account token 验证
func TokenVAlidate(client *AliHttpClient, token string) (*OpenAccount, *ResponseEntity, *ErrorResponse) {
	values := *url.Values{
		"param_token":token,
	}
	resp := client.SendRequest("taobao.open.account.token.validate", values)
	account := *OpenAccount{}
	isOk, data, respEntity, errorResponse := client.ParserRespBody("open_account_token_validate_response", resp)
	if isOk {
		util.JsonDecodeS(data, account)
	}
	return account, respEntity, errorResponse
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
