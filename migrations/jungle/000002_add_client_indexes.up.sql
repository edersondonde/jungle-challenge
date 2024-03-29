BEGIN;

ALTER TABLE client ADD CONSTRAINT client_pk PRIMARY KEY (uid);
CREATE INDEX IF NOT EXISTS client_birthday_idx ON public.client USING btree (birthday);
CREATE INDEX IF NOT EXISTS client_name_idx ON public.client USING btree (name);

COMMIT;