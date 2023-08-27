CREATE SCHEMA IF NOT EXISTS content;

CREATE TABLE IF NOT EXISTS content.users
(
    id            serial primary key,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS content.notes
(
    id      serial primary key,
    title   varchar(255)                                        not null,
    text    text                                                not null,
    user_id int references content.users (id) on delete cascade not null
);