-- create schema
create database if not exists letter;


-- create table
create table if not exists letter.messages
(
    id           int unsigned auto_increment
    primary key,
    from_user_id int unsigned default '0'               not null,
    to_user_id   int unsigned default '0'               not null,
    content      text                                   null,
    create_at    timestamp    default CURRENT_TIMESTAMP not null
);

create index messages_create_at_index
    on letter.messages (create_at);

create index messages_from_user_id_index
    on letter.messages (from_user_id);

create index messages_to_user_id_index
    on letter.messages (to_user_id);

create table if not exists letter.user
(
    uid       int unsigned auto_increment
    primary key,
    name      varchar(20)  default ''                not null,
    password  varchar(32)  default ''                not null,
    avatar    varchar(256) default ''                not null,
    create_at timestamp    default CURRENT_TIMESTAMP not null,
    update_at timestamp    default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint user_name_uindex
    unique (name)
    );

create table if not exists letter.user_friend
(
    id        int unsigned auto_increment
    primary key,
    user_id   int unsigned default '0' not null,
    friend_id int unsigned default '0' null
);

create index index_user_id
    on letter.user_friend (user_id);

