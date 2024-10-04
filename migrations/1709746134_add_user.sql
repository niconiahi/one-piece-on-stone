-- +goose Up
-- +goose StatementBegin
create table user (
    id integer primary key autoincrement,
    name text not null,
    surname text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
-- +goose StatementEnd
