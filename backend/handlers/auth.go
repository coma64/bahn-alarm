package handlers

import (
	"context"
	"github.com/coma64/bahn-alarm-backend/server"
)

func (b *BahnAlarmApi) PostAuthLogin(ctx context.Context, request server.PostAuthLoginRequestObject) (server.PostAuthLoginResponseObject, error) {
	return server.PostAuthLogin204Response{}, nil
}

func (b *BahnAlarmApi) PostAuthLogout(ctx context.Context, request server.PostAuthLogoutRequestObject) (server.PostAuthLogoutResponseObject, error) {
	return server.PostAuthLogout204Response{}, nil
}

func (b *BahnAlarmApi) GetAuthMe(ctx context.Context, request server.GetAuthMeRequestObject) (server.GetAuthMeResponseObject, error) {
	return server.GetAuthMe200JSONResponse{}, nil
}

func (b *BahnAlarmApi) PostAuthRegister(ctx context.Context, request server.PostAuthRegisterRequestObject) (server.PostAuthRegisterResponseObject, error) {
	return server.PostAuthRegister201Response{}, nil
}
