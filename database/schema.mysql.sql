USE piglatin;

-- DROP TABLE IF EXISTS pig_latin;
CREATE TABLE pig_latin (
    id INT NOT NULL auto_increment primary key,
    text TEXT NOT NULL,
    translation TEXT NOT NULL
-- index not necessary
);

