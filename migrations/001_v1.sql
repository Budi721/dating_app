-- Write your migrate up statements here
CREATE TABLE "public"."member_access"
(
    "member_id"           varchar(36),
    "user_name"           varchar(100),
    "user_password"       varchar(16),
    "join_date"           date,
    "verification_status" varchar(2),
    PRIMARY KEY ("member_id")
);
CREATE UNIQUE INDEX "member_user_name_uq" ON "public"."member_access" USING BTREE ("user_name");
INSERT INTO "public"."member_access"(member_id, user_name, user_password, join_date, verification_status)
VALUES ('0c490c0f-13d7-4acc-85d8-bb9f9bfcc53f', 'doni@enigmacamp.com', 'v60lXypb_XaBfCU', '2022-01-11', 'N');
INSERT INTO "public"."member_access"(member_id, user_name, user_password, join_date, verification_status)
VALUES ('a9127c2f-f878-40ab-8903-917954ab814c', 'tika@enigmacamp.com', '8BOMGK1idWdUz76', '2022-01-11', 'Y');

CREATE TABLE "public"."member_preference"
(
    "preference_id"     varchar(36),
    "member_id"         varchar(36),
    "looking_gender"    varchar(1),
    "looking_domicile"  varchar(100),
    "looking_start_age" int,
    "looking_end_age"   int,
    PRIMARY KEY ("preference_id")
);
CREATE UNIQUE INDEX "member_id_pref_uq" ON "public"."member_preference" USING BTREE ("member_id");

CREATE TABLE "public"."member_personal_information"
(
    "personal_information_id" varchar(36),
    "member_id"               varchar(36),
    "name"                    varchar(50),
    "bod"                     date,
    "gender"                  varchar(1),
    "photo_path"              varchar(150),
    "self_description"        varchar(250),
    PRIMARY KEY ("personal_information_id")
);
CREATE UNIQUE INDEX "member_id_info_uq" ON "public"."member_personal_information" USING BTREE ("member_id");

CREATE TABLE "public"."member_contact_information"
(
    "contact_information_id" varchar(36),
    "member_id"              varchar(36),
    "mobile_phone"           varchar(50),
    "instagram_id"           varchar(75),
    "twitter_id"             varchar(75),
    "email"                  varchar(75),
    PRIMARY KEY ("contact_information_id")
);
CREATE UNIQUE INDEX "member_id_contact_uq" ON "public"."member_contact_information" USING BTREE ("member_id");

CREATE TABLE "public"."member_address_information"
(
    "address_information_id" varchar(36),
    "member_id"              varchar(36),
    "address"                varchar(150),
    "city"                   varchar(75),
    "postal_code"            varchar(7),
    PRIMARY KEY ("address_information_id")
);
CREATE UNIQUE INDEX "member_id_address_uq" ON "public"."member_address_information" USING BTREE ("member_id");
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
drop table member_access;
drop table member_preference;
drop table member_personal_information;
drop table member_contact_information;
drop table member_address_information;

