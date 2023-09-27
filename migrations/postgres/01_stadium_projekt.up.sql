CREATE TABLE "country" (
  "id" text PRIMARY KEY,
  "name" text,
  "location" text,
  "created_at" text,
  "updated_at" text,
  "deleted_at" text
);

CREATE TABLE "images" (
  "id" text PRIMARY KEY,
  "feild_id" text,
  "path" text,
  "created_at" text,
  "updated_at" text,
  "deleted_at" text
);

CREATE TABLE "feild" (
  "id" text PRIMARY KEY,
  "country_id" text,
  "name" text,
  "created_at" text,
  "updated_at" text,
  "deleted_at" text
);

CREATE TABLE "users" (
  "login" text,
  "password_hash" text,
  "token" text,
  "created_at" text
);

CREATE TABLE "weather" (
  "id" text PRIMARY KEY,
  "result_humidity" text,
  "result_sun" text,
  "result_wind" text,
  "device_id" text,
  "created_at" text,
  "updated_at" text,
  "deleted_at" text
);

ALTER TABLE "images" ADD FOREIGN KEY ("feild_id") REFERENCES "feild" ("id");

ALTER TABLE "feild" ADD FOREIGN KEY ("country_id") REFERENCES "country" ("id");

ALTER TABLE "weather" ADD FOREIGN KEY ("device_id") REFERENCES "feild" ("id");