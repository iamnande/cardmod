BEGIN;

CREATE TABLE IF NOT EXISTS "users" (
    "id" uuid NOT NULL,
    "name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "create_time" timestamp with time zone NOT NULL,
    "update_time" timestamp with time zone NOT NULL,
    "delete_time" timestamp with time zone NULL,
    PRIMARY KEY("id")
);

COMMIT;