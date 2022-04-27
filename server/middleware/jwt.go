package middleware

import (
	"errors"
	"strconv"
	"time"

	"myapp/global"
	"myapp/model/common/request"
	"myapp/model/common/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}


var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)


func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")

		if token == "" {
			response.FailWithMessage("未登录或非法访问！", c)
		}

		j := &JWT{
			[]byte(global.GB_CONFIG.JWT.SigningKey),
		}
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == ErrTokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// 当过期时间 - 当前时间 < 缓冲时间时，进行续期
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {

			claims.ExpiresAt = time.Now().Unix() + global.GB_CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)

			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid

	} else {
		return nil, ErrTokenInvalid
	}
}

// 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.GB_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.GB_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GB_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    global.GB_CONFIG.JWT.Issuer,                          // 签名的发行者
		},
	}
	return claims
}