
-- First version password storage:
-- ps_hash = SHA256(password .. pw_salt)

CREATE TYPE sched_type AS ENUM (
    'oneshot',
    'interval',
    'weekly',
    'month_day',
    'month_weekday',
    'annual'
);

CREATE TABLE users(
    id SERIAL NOT NULL PRIMARY KEY,
    username VARCHAR(32) NOT NULL UNIQUE,
    pw_hash BYTEA NOT NULL CHECK(OCTET_LENGTH(pw_hash) = 32),
    pw_salt BYTEA NOT NULL,
    is_enabled BOOL NOT NULL,
    is_admin BOOL NOT NULL,
    min_sess_id INT NOT NULL DEFAULT 0,
    next_sess_id INT NOT NULL DEFAULT 1
);

-- TBD: How to give a group ownership of an object?

CREATE TABLE groups(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(64) NOT NULL UNIQUE
);

-- Many-to-many relation on users belonging to groups.
CREATE TABLE group_users(
    user_id INT NULL,
    group_id INT NOT NULL,
    is_owner BOOL NOT NULL,
    is_admin BOOL NOT NULL,
    PRIMARY KEY(user_id, group_id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(group_id) REFERENCES groups(id)
);

-- Create indexes to speed retrieving membership by group or user id.
CREATE INDEX ON group_users (user_id);
CREATE INDEX ON group_users (group_id);

CREATE TABLE task_defs(
    id SERIAL NOT NULL PRIMARY KEY,
    description VARCHAR(280) NOT NULL,
    owner_id INT NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users(id)
);

CREATE TABLE task_instances(
    id SERIAL NOT NULL PRIMARY KEY,
    task_id INT NOT NULL,
    FOREIGN KEY (task_id) REFERENCES task_defs(id)
);

CREATE TABLE task_schedules(
    id SERIAL NOT NULL PRIMARY KEY,
    task_id INT NOT NULL,
    owner_id INT NOT NULL,

    type sched_type NOT NULL,
    exact_only BOOL NOT NULL,
    sched_time INT NOT NULL,
    sched_weekday INT NULL,
    next_due BIGINT NOT NULL,
    is_active BOOL NOT NULL,

    FOREIGN KEY (task_id) REFERENCES task_defs(id),
    FOREIGN KEY (owner_id) REFERENCES users(id)
);

-- Single row config table.
CREATE TABLE config(
    schema_ver INT NOT NULL
);

-- Currently, only the schema version is stored.
INSERT INTO config VALUES(1);
