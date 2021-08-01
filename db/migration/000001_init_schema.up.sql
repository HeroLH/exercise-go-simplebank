CREATE TABLE "accounts"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "owner"      varchar     NOT NULL,
    "balance"    bigint      NOT NULL,
    "currency"   varchar     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "account_id" bigint      NOT NULL,
    "amount"     bigint      NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers"
(
    "id"              BIGSERIAL PRIMARY KEY,
    "from_account_id" bigint      NOT NULL,
    "to_account_id"   bigint      NOT NULL,
    "amount"          bigint      NOT NULL,
    "created_at"      timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "entries"
    ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
    ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
    ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

CREATE
INDEX ON "accounts" ("owner");
CREATE
INDEX ON "entries" ("account_id");
CREATE
INDEX ON "transfers" ("from_account_id");
CREATE
INDEX ON "transfers" ("to_account_id");
CREATE
INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT
ON COLUMN "accounts"."id" IS '帐户ID';
COMMENT
ON COLUMN "accounts"."owner" IS '所属人';
COMMENT
ON COLUMN "accounts"."created_at" IS '创建时间';
COMMENT
ON COLUMN "entries"."amount" IS 'can be negative or positive';
COMMENT
ON COLUMN "entries"."created_at" IS '创建时间';
COMMENT
ON COLUMN "transfers"."amount" IS 'must be positive';
COMMENT
ON COLUMN "transfers"."created_at" IS '创建时间';