package middleware

import (
	"bibi/biz/model/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"log"
	"net/http"
	"time"
)

var identityKey = "id"

func JWT() *jwt.HertzJWTMiddleware {
	// the jwt middleware
	var loginRes user.LoginResp
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "BibiBibi",                  //用于设置所属领域名称，默认为 hertz jwt
		Key:         []byte("PANDORA PARADOXXX"), //用于设置签名密钥（必要配置）
		Timeout:     time.Hour * 24 * 7,          //token 过期时间
		MaxRefresh:  time.Hour * 24 * 7,          //最大 token 刷新时间
		IdentityKey: identityKey,

		//PayloadFunc
		//用于设置登录时为 token 添加自定义负载信息的函数，如果不传入这个参数，
		//则 token 的 payload 部分默认存储 token 的过期时间和创建时间，
		//额外存储了username
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user.LoginReq); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		//通过在 IdentityHandler 内配合使用 identityKey，将存储用户信息的 token 从请求
		//上下文中取出并提取需要的信息，封装成 User 结构，以 identityKey 为 key，
		//User 为 value 存入请求上下文当中以备后续使用。
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &user.LoginReq{
				Username: claims[identityKey].(string),
			}
		},

		/*Login Handler*/
		//登录时触发，用于认证用户的登录信息。
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			data, err, res := Login(ctx, c)
			loginRes = res
			if err == nil && loginRes.Status == e.SUCCESS {
				return data, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		//LoginResponse 作为 LoginHandler 的响应结果
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, model.TokenResponse{
				Status: e.SUCCESS,
				Data:   loginRes.Data,
				Msg:    loginRes.Msg,
				Token:  token,
				Expire: expire.Format(time.RFC3339),
			})
		},
		//Unauthorized 用于设置 jwt 授权失败后的响应函数，
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},

		//用于验证用户是否有访问权限
		//Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
		//	if _, ok := data.(*service.UserService); ok {
		//		return true
		//	}
		//	return false
		//},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}
