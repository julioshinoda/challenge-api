
CREATE SEQUENCE IF NOT EXISTS users_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	CACHE 1
	NO CYCLE;

CREATE TABLE IF NOT EXISTS users (
	id bigint NOT NULL DEFAULT nextval('users_id_seq'::regclass),
	username varchar NULL,
	secret varchar NULL,
	created_at date NULL,
	CONSTRAINT username_un UNIQUE (username),
	CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE SEQUENCE IF NOT EXISTS expressions_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	CACHE 1
	NO CYCLE;

CREATE TABLE IF NOT EXISTS expressions (
	id bigint NOT NULL DEFAULT nextval('expressions_id_seq'::regclass),
	expression TEXT NOT NULL,
	created_at date NULL,
    CONSTRAINT expressions_pk PRIMARY KEY (id)
);


INSERT INTO public.users
(id, username, secret, created_at)
VALUES(1, 'test', 'secret', NULL);
