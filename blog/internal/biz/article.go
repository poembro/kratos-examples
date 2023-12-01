package biz

import (
	"context"
	"fmt"
	"strconv"
	"time"

	pb "github.com/go-kratos/examples/blog/api/blog/v1"
	common "github.com/go-kratos/examples/blog/api/common/v1"
	"github.com/go-kratos/examples/blog/internal/data/model"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	stdJwt "github.com/golang-jwt/jwt/v4"
)

type ArticleRepo interface {
	// db
	ListArticle(ctx context.Context, typ string, page, limit int) ([]model.Article, int64, error)
	GetArticle(ctx context.Context, where ...interface{}) (*model.Article, error)
	CreateArticle(ctx context.Context, article *model.Article) (*model.Article, error)
	UpdateArticle(ctx context.Context, id int64, article *model.Article) error
	DeleteArticle(ctx context.Context, id int64) error

	// redis
	GetArticleLike(ctx context.Context, id int64) (rv int64, err error)
	IncArticleLike(ctx context.Context, id int64) error
}

type ArticleUsecase struct {
	logger *log.Helper
	repo   ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, logger: log.NewHelper(logger)}
}

func (d *ArticleUsecase) Verify() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			header, ok := transport.FromServerContext(ctx)
			if !ok {
				d.logger.Info("Verify:", "请求错误 ctx 没有 header 信息")
				return nil, errors.Unauthorized("UNAUTHORIZED", "请求错误")
			}
			//gameActivityUuid := header.RequestHeader().Get("game_activity_uuid")
			uid, _ := strconv.Atoi(header.RequestHeader().Get("uid"))
			if uid > 10 {
				d.logger.Info("Verify:", "test 参数缺失")

				return nil, common.ErrorUnauthorized("%s", "参数缺失")
			}
			// 获取活动信息，根据活动uuid
			return handler(ctx, req)
		}
	}
}

func (uc *ArticleUsecase) Login(ctx context.Context) error {
	jwtAlg := "HS256"
	jwtKey := "nmGDsxKwFE5yssxox3399PXnzlcY2bZbcp5cE61T0J7gsqZeTxuo5knGt9cbpCfK"
	//获取token
	ttl := time.Now().Add(7 * 24 * 3600)
	uid := 456

	token, err := stdJwt.NewWithClaims(stdJwt.GetSigningMethod(jwtAlg), &stdJwt.RegisteredClaims{
		ExpiresAt: stdJwt.NewNumericDate(ttl),
		ID:        fmt.Sprintf("%d", uid),
	}).SignedString([]byte(jwtKey))
	if err != nil {
		return common.ErrorUnauthorized("%s", "参数缺失")
	}

	uc.logger.Info("token : ", token)
	return nil
}

func (uc *ArticleUsecase) List(ctx context.Context) (ps []model.Article, err error) {
	ps, _, err = uc.repo.ListArticle(ctx, "", 1, 10)
	if err != nil {
		return
	}
	return
}

func (uc *ArticleUsecase) Get(ctx context.Context, id int64) (p *model.Article, err error) {
	uc.Login(ctx)

	p, err = uc.repo.GetArticle(ctx, "id = ?", id)
	if err != nil || p == nil {
		return
	}
	err = uc.repo.IncArticleLike(ctx, id)
	if err != nil {
		return
	}
	p.Like, err = uc.repo.GetArticleLike(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *ArticleUsecase) Create(ctx context.Context, req *pb.CreateArticleRequest) (*model.Article, error) {
	return uc.repo.CreateArticle(ctx, &model.Article{
		Title:   req.Title,
		Content: req.Content,
	})
}

func (uc *ArticleUsecase) Update(ctx context.Context, req *pb.UpdateArticleRequest) error {
	return uc.repo.UpdateArticle(ctx, req.Id, &model.Article{
		Title:   req.Title,
		Content: req.Content,
	})
}

func (uc *ArticleUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteArticle(ctx, id)
}
