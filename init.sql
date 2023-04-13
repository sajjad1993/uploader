-- table images
CREATE TABLE  IF NOT EXISTS public."images"
(
    "sha"        varchar(255)             NOT NULL,
    "size"       bigint                   NOT NULL,
    "chunk_size" bigint                   NOT NULL,
    "status"     varchar(255),
    "data"       varchar(255),
    "created_at" timestamp with time zone NOT NULL
);

CREATE INDEX ON public."images" ("sha");

-- table images
CREATE TABLE IF NOT EXSITS public."chunks"
(
    "sha"        varchar(255)             NOT NULL,
    "id"         int                      NOT NULL,
    "size"       bigint                   NOT NULL,
    "data"       varchar(255),
    "created_at" timestamp with time zone NOT NULL
);

CREATE INDEX ON public."chunks" ("sha");