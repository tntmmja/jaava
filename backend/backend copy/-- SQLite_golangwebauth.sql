-- SQLite

CREATE TABLE IF NOT EXISTS user (
id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
firstname TEXT NOT NULL,
lastname TEXT NOT NULL,
age INTEGER NOT NULL,
gender VARCHAR NOT NULL,
username VARCHAR NOT NULL,
email TEXT NOT NULL,
password TEXT NOT NULL,
createdDate REAL,
sessionID TEXT);


/*create table posts (
id INTEGER PRIMARY KEY AUTOINCREMENT,
user_id INTEGER NOT NULL,
title VARCHAR NOT NULL,
dislike VARCHAR NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (user_id) REFERENCES user(id));
*/

create table likes (
id INTEGER PRIMARY KEY AUTOINCREMENT,
user_id INTEGER NOT NULL,
post_id INTEGER NOT NULL,
comment_id INTEGER,
likes VARCHAR NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (user_id) REFERENCES user(id),
FOREIGN KEY (post_id) REFERENCES posts(id),
FOREIGN KEY (comment_id) REFERENCES comments(id));

INSERT INTO likes (user_id, post_id, likes) VALUES (4, 6, 'L');

DROP TABLE likes;

DELETE FROM likes;

DELETE FROM user WHERE id = 2;

/*
create table comments (
id INTEGER PRIMARY KEY AUTOINCREMENT,
user_id INTEGER NOT NULL,
post_id VARCHAR INTEGER,
text VARCHAR NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
*/

-- kustutab tabeli posts
--drop table posts

/*
select * from post;

insert into post(title, text) values('title1', 'text 1');
insert into post(title, text) values('title2', 'text 2');
SELECT id, title, text FROM post WHERE id=3;
*/

/*
CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
	user_id INTEGER NOT NULL
	title VARCHAR NOT NULL,
	text VARCHAR NOT NULL
	created_at TIMESTAMP DEFAULT:CURRENT TIMESTAMP
	like VARCHAR
	dislike VARCHAR
	);*/

    --update user set sessionId = 'suvaliune' where id = 3;

    -- alter table user add sessionID text;