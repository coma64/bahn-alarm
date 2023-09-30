-- +goose Up
-- +goose StatementBegin
alter table bahnapisearchresponses
    add column createdAt timestamp default now() not null,
    add column url text not null default '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table bahnapisearchresponses
    drop column createdAt,
    drop column url;
-- +goose StatementEnd
