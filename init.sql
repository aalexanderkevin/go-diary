CREATE DATABASE IF NOT EXISTS tap_talk;
USE tap_talk;
CREATE TABLE IF NOT EXISTS user (
    username varchar(50) NOT NULL PRIMARY KEY,
    password char(255) NOT NULL,
    email varchar(50) NOT NULL,
    name varchar(255) NOT NULL,
    birthday date NOT NULL
);

CREATE TABLE IF NOT EXISTS diary (
  user_id varchar(50) NOT NULL,
  content text NOT NULL,
  date date NOT NULL,
  PRIMARY KEY (user_id, date),
  CONSTRAINT fk_username_user
    FOREIGN KEY (user_id) REFERENCES user (username)
);