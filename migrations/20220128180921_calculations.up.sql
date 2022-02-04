BEGIN;

CREATE TABLE IF NOT EXISTS "calculations" (
    "id" uuid NOT NULL,
    "card_id" uuid NOT NULL,
    "magic_id" uuid NOT NULL,
    "card_ratio" int NOT NULL,
    "magic_ratio" int NOT NULL,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_card" FOREIGN KEY("card_id") REFERENCES "cards"("id") ON DELETE CASCADE,
    CONSTRAINT "fk_magic" FOREIGN KEY("magic_id") REFERENCES "magics"("id") ON DELETE CASCADE
);

COMMIT;