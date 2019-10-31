-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "builders" (
    id BIGINT PRIMARY KEY UNIQUE NOT NULL,
    name VARCHAR(60) NOT NULL
);

CREATE TABLE "models" (
    id BIGINT PRIMARY KEY UNIQUE NOT NULL,
    name VARCHAR(60) NOT NULL,
    builder_id BIGINT NOT NULL,

    FOREIGN KEY (builder_id) REFERENCES builders(id) ON DELETE CASCADE
);

CREATE TABLE "charters" (
    id BIGINT PRIMARY KEY UNIQUE NOT NULL,
    name VARCHAR(60) NOT NULL
);

CREATE TABLE "yachts" (
    id BIGINT PRIMARY KEY  NOT NULL,
    name VARCHAR(60) NOT NULL,
    model_id BIGINT NOT NULL,
    charter_id BIGINT NOT NULL,
    available_from TIMESTAMP WITH TIME ZONE,
    available_to TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (model_id) REFERENCES models(id) ON DELETE CASCADE,
    FOREIGN KEY (charter_id) REFERENCES charters(id) ON DELETE CASCADE
);

CREATE TABLE "update_info" (
    moment TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS "yachts";
DROP TABLE IF EXISTS "builders";
DROP TABLE IF EXISTS "models";
DROP TABLE IF EXISTS "charters";
DROP TABLE IF EXISTS "update_info";
