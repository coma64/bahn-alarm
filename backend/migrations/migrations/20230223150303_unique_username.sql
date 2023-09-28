-- +goose Up
-- +goose StatementBegin
alter table users alter column name type varchar(255);
alter table users add constraint users_unique_name unique (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users drop constraint users_unique_name;
alter table users alter column name type text;
-- +goose StatementEnd
