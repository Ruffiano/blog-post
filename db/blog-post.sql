CREATE TABLE "article" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint,
  "title" varchar,
  "content" varchar,
  "updated_at" timestamtz DEFAULT (now()),
  "created_at" timestamtz DEFAULT (now())
);

CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "updated_at" timestamtz DEFAULT (now()),
  "created_at" timestamtz DEFAULT (now())
);

ALTER TABLE "article" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
