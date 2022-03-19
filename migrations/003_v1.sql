-- Write your migrate up statements here
CREATE TABLE "public"."member_interest"
(
    "interest_id" varchar(36),
    "member_id"   varchar(36),
    PRIMARY KEY ("interest_id", "member_id")
);
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
drop table member_interest;
