CREATE DATABASE app;
use app;
CREATE TABLE users (
    id int(11) unsigned not null auto_increment,
    name varchar(255) not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    primary key (id)
);
INSERT INTO
    users(name)
VALUES("test1");
INSERT INTO
    users(name)
VALUES("test2");