package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (b *BahnAlarmApi) GetNotificationsVapidKeys(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, server.VapidKeys{PublicKey: config.Conf.PushNotifications.VapidKeys.Public})
}

func (b *BahnAlarmApi) GetNotificationsPushSubscriptions(ctx echo.Context, params server.GetNotificationsPushSubscriptionsParams) error {
	var body server.GetNotificationsPushSubscriptionsParams
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	offset, size := defaultPagination(params.Page, params.Size)

	var subs []models.PushNotificationSub
	if err := db.Db.SelectContext(
		ctx.Request().Context(),
		&subs,
		"select s.* from pushNotificationSubs s join users u on s.ownerId = u.id where u.name = $1 offset $2 fetch first $3 rows only",
		ctx.Get("username"),
		offset,
		size,
	); err != nil {
		return fmt.Errorf("error getting push subs: %w", err)
	}

	var pagination server.Pagination
	if err := db.Db.GetContext(
		ctx.Request().Context(),
		&pagination,
		"select count(s.id) totalItems from pushNotificationSubs s join users u on s.ownerId = u.id where u.name = $1",
		ctx.Get("username"),
	); err != nil {
		return fmt.Errorf("error getting total items: %w", err)
	}

	subSchemas := make([]server.PushNotificationSubscription, 0, len(subs))
	for _, sub := range subs {
		if schema, err := sub.ToSchema(); err != nil {
			return fmt.Errorf("error converting db push sub (%d) to schema: %w", sub.Id, err)
		} else {
			subSchemas = append(subSchemas, *schema)
		}
	}

	return ctx.JSON(http.StatusOK, server.PushNotificationSubscriptionList{
		Pagination:    pagination,
		Subscriptions: subSchemas,
	})
}

func (b *BahnAlarmApi) PostNotificationsPushSubscriptions(ctx echo.Context) error {
	body := server.PostNotificationsPushSubscriptionsJSONRequestBody{}
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	rawSub, err := json.Marshal(&body.Subscription)
	if err != nil {
		return fmt.Errorf("error converting raw sub to bytes: %w", err)
	}

	var pushSub models.PushNotificationSub
	if err := db.Db.GetContext(
		ctx.Request().Context(),
		&pushSub,
		"insert into pushNotificationSubs (ownerId, rawSubscription, name) values ((select id from users where name = $1), $2, $3) returning *",
		ctx.Get("username"),
		rawSub,
		body.Name,
	); err != nil {
		return fmt.Errorf("error inserting push sub: %w", err)
	}

	schema, err := pushSub.ToSchema()
	if err != nil {
		return fmt.Errorf("error converting push sub to schema: %w", err)
	}

	return ctx.JSON(http.StatusCreated, schema)
}

func (b *BahnAlarmApi) DeleteNotificationsPushSubscriptionsId(ctx echo.Context, id int) error {
	result, err := db.Db.ExecContext(
		ctx.Request().Context(),
		"delete from pushNotificationSubs s using users u where s.ownerId = u.id and u.name = $1 and s.id = $2",
		ctx.Get("username"),
		id,
	)
	if err != nil {
		return fmt.Errorf("error deleting push sub (%d): %w", id, err)
	}

	var count int64
	if count, err = result.RowsAffected(); err != nil {
		return fmt.Errorf("error getting affected row count while deleting (%d): %w", id, err)
	}

	if count == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (b *BahnAlarmApi) PutNotificationsPushSubscriptionsId(ctx echo.Context, id int) error {
	var body server.PutNotificationsPushSubscriptionsIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	rawSub, err := json.Marshal(body.Subscription)
	if err != nil {
		return fmt.Errorf("error marshaling push sub update: %w", err)
	}

	var result sql.Result
	result, err = db.Db.ExecContext(
		ctx.Request().Context(),
		"update pushNotificationSubs s set isenabled = $1, name = $2, rawSubscription = $3 from users u where s.ownerId = u.id and u.name = $4 and s.id = $5",
		body.IsEnabled,
		body.Name,
		rawSub,
		ctx.Get("username"),
		id,
	)
	if err != nil {
		return fmt.Errorf("error updating push sub (%d): %w", id, err)
	}

	var count int64
	count, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting affected rows for push sub (%d) update: %w", id, err)
	}

	if count == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.NoContent(http.StatusNoContent)
}
