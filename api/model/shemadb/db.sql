CREATE USER user0 WITH PASSWORD 'pass';
CREATE DATABASE mydb0 OWNER user0;

\c mydb0

CREATE TABLE events (
    id        uuid PRIMARY KEY,
    calendar_id       int,
    title         text,
    descr   text,
    start_time        timestamp,
    end_time         timestamp
);

GRANT ALL PRIVILEGES ON DATABASE "mydb0" to user0;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO user0;