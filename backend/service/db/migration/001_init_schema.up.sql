CREATE TABLE "profile" (
  "address" varchar(42) PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "message" (
  "transaction_hash" varchar(66) PRIMARY KEY,
  "address" varchar(42)  NOT NULL,
  "token" varchar(42) NOT NULL,
  "amount" varchar(256) NOT NULL,
  "fee" varchar(256) NOT NULL,
  "content" varchar(256) NOT NULL,
  "nick" varchar(64) NOT NULL,
  "added_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "piggy" (
  "address" varchar(42) PRIMARY KEY,
  "from_address" varchar(42) NOT NULL,
  "profile_address" varchar(42) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL,
  "added_at" timestamptz NOT NULL DEFAULT 'now()',
  "unlocks_at" timestamptz NOT NULL,
  "name" varchar(42) UNIQUE NULL
);

CREATE INDEX ON "piggy" ("name");

CREATE TABLE "session" (
  "id" uuid PRIMARY KEY,
  "address" varchar(42) NOT NULL,
  "refresh_token" text NOT NULL,
  "user_agent" text NOT NULL,
  "client_ip" text NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

COMMENT ON COLUMN "session"."id" IS 'randomly generated id';

COMMENT ON COLUMN "session"."address" IS 'blockchain address';