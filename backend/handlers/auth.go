package handlers

import (
	"database/sql"
	"github.com/coma64/bahn-alarm-backend/auth"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (b *BahnAlarmApi) PostAuthLogin(ctx echo.Context) error {
	var body server.PostAuthLoginJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	var user models.User
	if err := db.Db.GetContext(ctx.Request().Context(), &user, "select * from users where name = $1", body.Username); err != nil {
		if err == sql.ErrNoRows {
			return ctx.NoContent(http.StatusBadRequest)
		} else {
			return err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password)); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	token, expiresAt, err := auth.GenerateJwt(user.Name)
	if err != nil {
		return err
	}

	ctx.SetCookie(&http.Cookie{
		Name:     config.Conf.Jwt.Cookie,
		Value:    token,
		Expires:  expiresAt,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	return ctx.JSON(http.StatusOK, &server.User{
		CreatedAt: user.CreatedAt,
		Id:        user.Id,
		IsAdmin:   user.IsAdmin,
		Name:      user.Name,
	})
}

func (b *BahnAlarmApi) PostAuthLogout(ctx echo.Context) error {
	ctx.SetCookie(&http.Cookie{
		Name:     config.Conf.Jwt.Cookie,
		Value:    "",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
	return ctx.NoContent(http.StatusNoContent)
}

func (b *BahnAlarmApi) GetAuthMe(ctx echo.Context) error {
	var user models.User
	if err := db.Db.GetContext(ctx.Request().Context(), &user, "select * from users where name = $1", ctx.Get("username")); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user.ToSchema())
}

func (b *BahnAlarmApi) PostAuthRegister(ctx echo.Context) error {
	var body server.PostAuthRegisterJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	errMsg := ""
	if len(body.Username) < 3 {
		errMsg = "username must be at least 3 chars long"
	} else if len(body.Password) < 4 {
		errMsg = "password must be at least 4 characters long"
	}

	if errMsg != "" {
		return ctx.JSON(http.StatusBadRequest, server.ValidationFailed{Message: &errMsg})
	}

	var inviteToken models.InviteToken
	if err := db.Db.GetContext(ctx.Request().Context(), &inviteToken, "select * from inviteTokens where token = $1", body.InviteToken); err != nil {
		if err == sql.ErrNoRows {
			return ctx.NoContent(http.StatusBadRequest)
		} else {
			return err
		}
	}

	if inviteToken.UsedById != nil || inviteToken.ExpiresAt.Before(time.Now()) {
		return ctx.NoContent(http.StatusGone)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if _, err = db.Db.ExecContext(
		ctx.Request().Context(),
		`
		with userCreation as (insert into users (name, passwordHash) values ($1, $2) returning id)
		update inviteTokens set usedById = (select id from userCreation) where id = $3
		`,
		body.Username,
		passwordHash,
		inviteToken.Id,
	); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}
