-- +goose Up
Alter Table users Add api_key varchar(64) unique not null default (
    encode(sha256(random()::text::bytea),'hex')
);

-- +goose Down
Alter Table users Drop Column api_key