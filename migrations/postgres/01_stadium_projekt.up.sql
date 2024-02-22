CREATE TABLE "country" (
  "id" text PRIMARY KEY,
  "person_id" text,
  "region_name" text,
  "district_name" text,
  "village_name" text,
  "device_id" text,
  "status" integer,
  "pleace_name" text,
  "created_at" text
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
  "created_at" text
);

CREATE TABLE "persons" (
  "id" text PRIMARY KEY,
  "fullname" text,
  "phone" text,
  "email" text,
  "status" text,
  "parol" text,
  "created_at" text
);

CREATE TABLE "acses" (
  "id" text PRIMARY KEY,
  "fullname" text,
  "person_id" text,
  "status" text,
  "created_at" text
);