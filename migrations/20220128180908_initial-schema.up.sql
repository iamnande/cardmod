BEGIN;

CREATE TABLE IF NOT EXISTS "cards" (
    "id" uuid NOT NULL, 
    "name" varchar UNIQUE NOT NULL, 
    PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "magics" (
    "id" uuid NOT NULL, 
    "name" varchar UNIQUE NOT NULL, 
    PRIMARY KEY("id")
);

COMMIT;