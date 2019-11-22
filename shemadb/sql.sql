drop database mydb;
 drop user myuser; 
create user myuser with encrypted password 'mypass312'; 
create database mydb owner myuser;

/c mydb
create table events (
id uuid primary key,
calendar_id int,
title text,
descr text,
start_time timestamp,
end_time timestamp
);

GRANT ALL PRIVILEGES ON DATABASE "mydb" to myuser;