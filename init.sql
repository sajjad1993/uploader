-- table images
CREATE TABLE "images" IF NOT EXSITS
(
    "sha"        varchar(255)             NOT NULL,
    "size"       bigint                   NOT NULL,
    "chunk_size" bigint                   NOT NULL,
    "status"     varchar(255),
    "data"       varchar(255),
    "created_at" timestamp with time zone NOT NULL
);

CREATE INDEX ON "images" ("sha") IF NOT EXSITS;

-- table images
CREATE TABLE "chunks" IF NOT EXSITS
(
    "sha"        varchar(255)             NOT NULL,
    "id"         int                      NOT NULL,
    "size"       bigint                   NOT NULL,
    "data"       varchar(255),
    "created_at" timestamp with time zone NOT NULL
);

CREATE INDEX ON "chunks" ("sha") IF NOT EXSITS;