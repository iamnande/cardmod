BEGIN;

CREATE TABLE IF NOT EXISTS cards (
    id SERIAL NOT NULL,
    name VARCHAR UNIQUE NOT NULL,
    level SMALLINT NOT NULL CHECK (
        level > 0
        AND level <= 10
    )
);

CREATE TABLE IF NOT EXISTS items (
    id SERIAL NOT NULL,
    name VARCHAR UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS limitbreaks (
    id SERIAL NOT NULL,
    name VARCHAR UNIQUE NOT NULL
);

CREATE TYPE magic_purpose AS ENUM ('Offensive', 'Restorative', 'Indirect');
CREATE TABLE IF NOT EXISTS magics (
    id SERIAL NOT NULL,
    name VARCHAR UNIQUE NOT NULL,
    purpose magic_purpose NOT NULL
);

COMMIT;