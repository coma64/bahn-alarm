-- +goose Up
-- +goose StatementBegin
alter table departureInfos alter column day set default now();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table departureInfos alter column day drop default;
-- +goose StatementEnd
