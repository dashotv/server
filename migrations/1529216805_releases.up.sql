BEGIN;

CREATE TABLE releases (
	id uuid NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	raw text NOT NULL,
	title text NOT NULL,
	description text NOT NULL,
	season bigint NOT NULL,
	episode bigint NOT NULL,
	size bigint NOT NULL,
	guid text NOT NULL,
	resolution bigint NOT NULL,
	team text NOT NULL,
	author text NOT NULL,
	verified boolean NOT NULL,
	bluray boolean NOT NULL,
	uncensored boolean NOT NULL,
	checksum text NOT NULL,
	download text NOT NULL,
	source text NOT NULL,
	type text NOT NULL,
	published timestamptz NOT NULL
);


COMMIT;
