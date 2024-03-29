CREATE TABLE IF NOT EXISTS public.client (
	uid varchar NOT NULL,
	birthday date NOT NULL,
	sex varchar NOT NULL,
	"name" varchar NOT NULL,
	CONSTRAINT client_pk PRIMARY KEY (uid)
);
CREATE INDEX IF NOT EXISTS client_birthday_idx ON public.client USING btree (birthday);
CREATE INDEX IF NOT EXISTS client_name_idx ON public.client USING btree (name);