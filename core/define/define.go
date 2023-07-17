package define

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

// JwtKey jwt密钥
var JwtKey = "cloud-disk-key"

var MailPassword = "XZMBIILEXVINNTMG"

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间
var CodeExpire = 300

var TencentSecretKey = os.Getenv("xxx")
var TencentSecretID = os.Getenv("xxx")
var CosBucket = "https://merhythm-1318328416.cos.ap-nanjing.myqcloud.com"

// PageSize 分页的默认参数
var PageSize = 20

var DateTime = "2006-01-02 15:01:05"

// TokenExpire Token有效期
var TokenExpire int = 3600

var RefreshTokenExpire int = 3600
