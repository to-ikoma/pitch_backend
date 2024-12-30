-- +goose Up
-- +goose StatementBegin
create table hello (
  id integer, 
  name varchar(10)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table hello;
-- +goose StatementEnd
