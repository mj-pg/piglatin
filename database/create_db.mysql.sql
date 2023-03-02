-- mysql -u root in terminal
-- execute the ff

CREATE DATABASE piglatin;
CREATE USER 'dbuser'@'localhost' IDENTIFIED BY 'dbp@ss123';
GRANT ALL PRIVILEGES ON piglatin.* TO 'dbuser'@'localhost';
