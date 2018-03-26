-- Performs a hard reset of the public schema.

DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
