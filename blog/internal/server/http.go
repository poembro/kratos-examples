package server

import (
	"context"

	v1 "github.com/go-kratos/examples/blog/api/blog/v1"
	"github.com/go-kratos/examples/blog/internal/biz"
	"github.com/go-kratos/examples/blog/internal/conf"
	"github.com/go-kratos/examples/blog/internal/service"
	"github.com/go-kratos/examples/blog/pkg/codec"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdJwt "github.com/golang-jwt/jwt/v4"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, blog *service.BlogService, articleBiz *biz.ArticleUsecase) *http.Server {
	jwtAlg := "HS256"
	jwtKey := "nmGDsxKwFE5yssxox3399PXnzlcY2bZbcp5cE61T0J7gsqZeTxuo5knGt9cbpCfK"

	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			selector.Server(
				jwt.Server(func(token *stdJwt.Token) (interface{}, error) {
					return []byte(jwtKey), nil
				}, jwt.WithSigningMethod(stdJwt.GetSigningMethod(jwtAlg)), jwt.WithClaims(func() stdJwt.Claims {
					return &stdJwt.RegisteredClaims{}
				})),
			).Match(NewWhiteListMatcher()).Build(),
			selector.Server(
				articleBiz.Verify(),
			).Match(ShopBlacklistByGame()).Build(),
			validate.Validator(),
		),
		http.ErrorEncoder(codec.NormalizeErrorResponse),
		http.ResponseEncoder(codec.NormalizeResponse),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterBlogServiceHTTPServer(srv, blog)
	return srv
}

// jwt 白名单接口
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/blog.api.v1.BlogService/GetArticle"] = struct{}{}
	whiteList["/blog.api.v1.BlogService/UpdateArticle"] = struct{}{}
	whiteList["/blog.api.v1.BlogService/ListArticle"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; !ok {
			return true
		}
		return false
	}
}
