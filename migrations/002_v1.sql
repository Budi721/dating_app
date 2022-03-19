-- Write your migrate up statements here
CREATE TABLE "public"."master_interest"
(
    "interest_id" varchar(36),
    "interest"    varchar(75),
    PRIMARY KEY ("interest_id")
);
INSERT INTO "public"."master_interest"(interest_id, interest)
VALUES ('a8127c2f-f878-40ab-8903-917954ab814z', 'Coffee');
INSERT INTO "public"."master_interest"(interest_id, interest)
VALUES ('b9127c2f-f878-40ab-8903-917954ab814x', 'Live Music');
INSERT INTO "public"."master_interest"(interest_id, interest)
VALUES ('c1027c2f-f878-40ab-8903-917954ab814y', 'Dog');
INSERT INTO "public"."master_interest"(interest_id, interest)
VALUES ('d1127c2f-f878-40ab-8903-917954ab814z', 'Gym');
INSERT INTO "public"."master_interest"(interest_id, interest)
VALUES ('e1227c2f-f878-40ab-8903-917954ab814w', 'Beer');


---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
drop table master_interest;

