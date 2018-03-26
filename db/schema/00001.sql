
-- First version password storage:
-- ps_hash = SHA256(password .. pw_salt)

CREATE DOMAIN bytes32 AS BYTEA CHECK(OCTET_LENGTH(VALUE) = 32);

CREATE TABLE users(
    id SERIAL NOT NULL PRIMARY KEY,
    username VARCHAR(32) NOT NULL UNIQUE,
    pw_hash BYTEA NOT NULL,
    pw_salt BYTEA NOT NULL,
    is_enabled BOOL NOT NULL,
    is_admin BOOL NOT NULL,
    min_sess_id INT NOT NULL DEFAULT 0,
    next_sess_id INT NOT NULL DEFAULT 1
);

-- Single row config table.
CREATE TABLE config(
    schema_ver INT NOT NULL
);

-- Currently, only the schema version is stored.
INSERT INTO config VALUES(1);
