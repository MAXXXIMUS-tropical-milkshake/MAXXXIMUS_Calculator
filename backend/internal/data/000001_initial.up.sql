CREATE TABLE IF NOT EXISTS "history" (
  "id" serial PRIMARY KEY,
  "user_id" integer NOT NULL,
  "expression" text NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS "users" (
  "id" serial PRIMARY KEY
);

ALTER TABLE "history" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
