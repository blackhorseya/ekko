package biz

import (
	"crypto/sha256"
	"fmt"

	"github.com/blackhorseya/ekko/internal/app/domain/user/biz/repo"
	"github.com/blackhorseya/ekko/internal/pkg/errorx"
	"github.com/blackhorseya/ekko/internal/pkg/tokenx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	ub "github.com/blackhorseya/ekko/pkg/entity/domain/user/biz"
	um "github.com/blackhorseya/ekko/pkg/entity/domain/user/model"
	"github.com/blackhorseya/ekko/pkg/genx"
	"go.uber.org/zap"
)

type impl struct {
	repo      repo.IRepo
	node      genx.Generator
	tokenizer tokenx.Tokenizer
}

// NewImpl serve caller to create an IBiz
func NewImpl(repo repo.IRepo, node genx.Generator, tokenizer tokenx.Tokenizer) ub.IBiz {
	return &impl{
		repo:      repo,
		node:      node,
		tokenizer: tokenizer,
	}
}

func (i *impl) Signup(ctx contextx.Contextx, username, password string) (info *um.Profile, err error) {
	if username == "" {
		return nil, errorx.ErrInvalidUsername
	}

	if password == "" {
		return nil, errorx.ErrInvalidPassword
	}

	newUser := &um.Profile{
		Id:        i.node.Int64(),
		Username:  username,
		Password:  fmt.Sprintf("%x", sha256.Sum256([]byte(password))),
		Token:     "",
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	ret, err := i.repo.Register(ctx, newUser)
	if err != nil {
		ctx.Error(errorx.ErrRegisterProfile.Error(), zap.Error(err))
		return nil, errorx.ErrRegisterProfile
	}

	return ret, nil
}

func (i *impl) Login(ctx contextx.Contextx, username, password string) (info *um.Profile, err error) {
	if username == "" {
		return nil, errorx.ErrInvalidUsername
	}

	if password == "" {
		return nil, errorx.ErrInvalidPassword
	}

	exists, err := i.repo.GetProfileByUsername(ctx, username)
	if err != nil {
		ctx.Error(errorx.ErrGetProfile.Error(), zap.Error(err), zap.String("username", username))
		return nil, errorx.ErrGetProfile
	}
	if exists == nil {
		ctx.Error(errorx.ErrUserNotFound.Error(), zap.String("username", username))
		return nil, errorx.ErrUserNotFound
	}

	if exists.Password != fmt.Sprintf("%x", sha256.Sum256([]byte(password))) {
		ctx.Error(errorx.ErrInvalidPassword.Error(), zap.String("username", username))
		return nil, errorx.ErrInvalidPassword
	}

	token, err := i.tokenizer.NewToken(exists)
	if err != nil {
		ctx.Error(errorx.ErrNewToken.Error(), zap.Error(err))
		return nil, errorx.ErrNewToken
	}

	ret, err := i.repo.UpdateToken(ctx, exists, token)
	if err != nil {
		ctx.Error(errorx.ErrUpdateToken.Error(), zap.Error(err))
		return nil, errorx.ErrUpdateToken
	}

	return ret, nil
}

func (i *impl) Logout(ctx contextx.Contextx, who *um.Profile) error {
	if who == nil {
		return errorx.ErrInvalidProfile
	}

	_, err := i.repo.UpdateToken(ctx, who, "")
	if err != nil {
		ctx.Error(errorx.ErrUpdateToken.Error(), zap.Error(err))
		return errorx.ErrUpdateToken
	}

	return nil
}

func (i *impl) WhoAmI(ctx contextx.Contextx, token string) (info *um.Profile, err error) {
	// todo: 2023/4/16|sean|impl me
	panic("implement me")
}
