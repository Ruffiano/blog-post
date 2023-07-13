CREATE TABLE "article" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint,
  "title" varchar,
  "content" varchar,
  "updated_at"  timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE IF EXISTS "article" ADD CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "user" ("id");