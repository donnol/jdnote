package authsrv

import "context"

type AuthMock struct {
	CheckLoginFunc func(ctx context.Context) error

	CheckPermFunc func(ctx context.Context, perms []string) error

	CheckUserExistFunc func(ctx context.Context) error

	CheckUserPermFunc func(ctx context.Context, perms []string) error
}

var _ IAuth = &AuthMock{}

func (mockRecv *AuthMock) CheckLogin(ctx context.Context) error {
	return mockRecv.CheckLoginFunc(ctx)
}

func (mockRecv *AuthMock) CheckPerm(ctx context.Context, perms []string) error {
	return mockRecv.CheckPermFunc(ctx, perms)
}

func (mockRecv *AuthMock) CheckUserExist(ctx context.Context) error {
	return mockRecv.CheckUserExistFunc(ctx)
}

func (mockRecv *AuthMock) CheckUserPerm(ctx context.Context, perms []string) error {
	return mockRecv.CheckUserPermFunc(ctx, perms)
}
