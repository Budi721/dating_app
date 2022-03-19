-- Write your migrate up statements here
CREATE TABLE "public"."member_partner"
(
    "member_id"   varchar(36),
    "partner_id"   varchar(36),
    PRIMARY KEY ("member_id", "partner_id")
);
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
drop table member_partner;
