CREATE TABLE "stadium" (
  "id" text PRIMARY KEY,
  "user_id" text,
  "name" text,
  "description" text,
  "phone" text,
  "card" text,
  "card_person" text,
  "size" text,
  "price" text,
  "Location" text,
  "latitude" text,
  "longitude" text,
  "created_at" text,
  "updated_at" text
);

CREATE TABLE "users" (
  "id" text PRIMARY KEY,
  "name" text,
  "phone" text,
  "login" text,
  "password_hash" text,
  "created_at" text,bd28
  "updated_at" text
);

CREATE TABLE "images" (
  "id" text PRIMARY KEY,
  "stadium_id" text,
  "path" text,
  "created_at" text,
  "updated_at" text
);

CREATE TABLE "time_planning" (
  "id" text,
  "day" text,
  "clock" text,
  "cleint_name" text,
  "cleint_phone" text
);

CREATE TABLE "people" (
  "id" text PRIMARY KEY,
  "phone" text unique,
  "first_name" text,
  "favorite_stadium" text,
  "created_at" text
);

CREATE TABLE "likes" (
  "person_id" text,
  "stadium_id" text,
  "created_at" text
);

ALTER TABLE "stadium" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "images" ADD FOREIGN KEY ("stadium_id") REFERENCES "stadium" ("id");

ALTER TABLE "time_planning" ADD FOREIGN KEY ("id") REFERENCES "stadium" ("id");

ALTER TABLE "people" ADD FOREIGN KEY ("favorite_stadium") REFERENCES "stadium" ("id");

ALTER TABLE "likes" ADD FOREIGN KEY ("person_id") REFERENCES "people" ("id");

ALTER TABLE "likes" ADD FOREIGN KEY ("stadium_id") REFERENCES "stadium" ("id");