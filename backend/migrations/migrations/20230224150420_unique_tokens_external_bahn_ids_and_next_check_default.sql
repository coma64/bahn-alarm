-- +goose Up
-- +goose StatementBegin
alter table bahnStations alter column externalId type varchar(128);
alter table inviteTokens alter column token type varchar(128);
alter table inviteTokens add constraint inviteTokensUniqueTokens unique (token);
alter table departures alter column nextCheck SET default now();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table bahnStations alter column externalId type varchar(255);
alter table inviteTokens alter column token type text;
alter table inviteTokens drop constraint inviteTokensUniqueTokens;
alter table departures alter column nextCheck drop default;
-- +goose StatementEnd
