package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"my-blog-golang/server/utils"
	"my-blog-golang/server/utils/errmsg"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour) // token有效期
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 时间
			Issuer:    "ginblog",         // 签发人
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey) // 加入key生成token
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCSE

}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	var claims MyClaims
	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok { // 官方写法照抄就行
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errmsg.ERROR_TOKEN_WRONG
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errmsg.ERROR_TOKEN_RUNTIME
			} else {
				return nil, errmsg.ERROR_TOKEN_TYPE_WRONG
			}
		}
	}
	if setToken != nil {
		if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
			return key, errmsg.SUCCSE
		} else {
			return nil, errmsg.ERROR_TOKEN_WRONG
		}
	}
	return nil, errmsg.ERROR_TOKEN_WRONG
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")

		// token不存在
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort() // 不调用c.Next()后续的函数
			return
		}

		// token格式错误
		chkToken := strings.Split(tokenHeader, " ")
		if len(chkToken) == 0 {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// token格式错误
		if len(chkToken) != 2 && chkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		key, tCode := CheckToken(chkToken[1])

		// token不正确
		if tCode != errmsg.SUCCSE {
			code = tCode
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// token正确
		c.Set("username", key) // 保存token到上下文
		c.Next()               // 调用后续的函数
	}
}
