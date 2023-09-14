CREATE TABLE "county" (
  "id" text PRIMARY KEY,
  "name" text,
  "location" text,
  "created_at" text,
  "updated_at" text,
  "deleted_at" text
);

CREATE TABLE "images" (
  "id" text PRIMARY KEY,
  "country_id" text,
  "path" text,
  "created_at" text,
  "updated_at" text,
  "deleted_at" text
);

CREATE TABLE "device" (
  "id" text PRIMARY KEY,
  "user_id" text,
  "country_id" text,
  "name" text,
  "created_at" text,
  "updated_at" text,
  "deleted_at" text
);

CREATE TABLE "users" (
  "id" text PRIMARY KEY,
  "login" text,
  "phone" text,
  "password_hash" text,
  "created_at" text,
  "updated_at" text
);

CREATE TABLE "sun" (
  "id" text,
  "result_sun" text,
  "device_id" text,
  "created_at" text,
  "updated_at" text,
  "deleted_at" text
);

CREATE TABLE "wind" (
  "id" text PRIMARY KEY,
  "result_wind" text,
  "device_id" text,
  "created_at" text,
  "updated_at" text,
  "deleted_at" text
);

CREATE TABLE "humidity" (
  "id" text PRIMARY KEY,
  "result_humidity" text,
  "device_id" text,
  "created_at" text,
  "updated_at" text,
  "deleted_at" text
);

ALTER TABLE "images" ADD FOREIGN KEY ("country_id") REFERENCES "county" ("id");

ALTER TABLE "device" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "device" ADD FOREIGN KEY ("country_id") REFERENCES "county" ("id");

ALTER TABLE "sun" ADD FOREIGN KEY ("device_id") REFERENCES "device" ("id");

ALTER TABLE "wind" ADD FOREIGN KEY ("device_id") REFERENCES "device" ("id");

ALTER TABLE "humidity" ADD FOREIGN KEY ("device_id") REFERENCES "device" ("id");