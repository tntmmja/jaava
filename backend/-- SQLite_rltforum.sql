-- SQLite
-- kasu kaivitamiseks parem klops - run selected query
/*
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
*/


/*
create table user (
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
*/

-- kustutab tabeli posts
/*drop table user
*/

ALTER TABLE user RENAME COLUMN username TO nickname;