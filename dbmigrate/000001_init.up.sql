CREATE table users
(
    id            serial       not null unique,
    username      varchar(255) not null unique,
    name          varchar(50)  not null,
    surname       varchar(50)  not null,
    email         varchar(250) not null,
    phone         varchar(11),
    password_hash varchar(255) not null
);

create table posts(
    post_id     serial       not null unique,
    header      varchar(255) not null,
    body        varchar(1000),
    author_id int references users(id) on delete cascade not null
);