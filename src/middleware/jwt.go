package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"myBlog/model"
	"net/http"
	"strings"
	"time"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(model.ServerConf.JwtKey),
	}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 定义错误
var (
	TokenExpired     = errors.New("token已过期,请重新登录")
	TokenNotValidYet = errors.New("token无效,请重新登录")
	TokenMalformed   = errors.New("token不正确,请重新登录")
	TokenInvalid     = errors.New("这不是一个token,请重新登录")
)

// CreateToken 生成token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token
func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {

	return func(c *gin.Context) {

		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" { // 没有token这个字段，直接返回结束
			code = model.ErrTokenExists
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": model.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 将header 按空格进行划分
		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": model.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 前端固定传过来的两个一个字段，格式为Bearer token
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": model.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 解析token
		j := NewJWT()
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {

			// token过期
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status":  model.Error,
					"message": "token授权已过期,请重新登录",
					"data":    nil,
				})
				c.Abort()
				return
			}

			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status":  model.Error,
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Set("username", claims)
		c.Next()
	}
}

// token生成函数
func (j *JWT) SetToken(c *gin.Context, user *model.User) (token string, err error) {

	// 定义jwt参数
	claims := MyClaims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "myBlog",
		},
	}

	// 生成jwt token
	token, err = j.CreateToken(claims)
	return
}
