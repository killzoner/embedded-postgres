-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE beer_catalogue (
    id SERIAL PRIMARY KEY,
    name TEXT,
    consumed BOOL DEFAULT TRUE,
    rating DOUBLE PRECISION,
    tags TEXT [] NOT NULL
);

INSERT INTO beer_catalogue (name, consumed, rating, tags)
VALUES ('Punk IPA', true, 68.29, '{"1", "2", "3"}');
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.