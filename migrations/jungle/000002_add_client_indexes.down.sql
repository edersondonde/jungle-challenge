BEGIN;

ALTER TABLE client DROP CONSTRAINT client_pk;
DROP INDEX IF EXISTS client_birthday_idx;
DROP INDEX IF EXISTS client_name_idx;

COMMIT;