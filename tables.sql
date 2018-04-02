CREATE SCHEMA IF NOT EXISTS SnatchHQ CHARACTER SET 'utf8';

CREATE USER 'Snatch'@'localhost' IDENTIFIED BY 'AndHide';
GRANT ALL PRIVILEGES ON SnatchHQ . * TO 'Snatch'@'localhost';

USE SnatchHQ;

CREATE TABLE IF NOT EXISTS Users (
    id BIGINT AUTO_INCREMENT,
    username VARCHAR(12) NOT NULL UNIQUE,
    phone_number CHAR(11) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (id)
) 
ENGINE = InnoDB; 

CREATE TABLE IF NOT EXISTS Locations (
    id BIGINT AUTO_INCREMENT,
    longitude FLOAT(10, 6) NOT NULL,
    latitude FLOAT(10, 6) NOT NULL,
    user_id BIGINT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES Users(id)
) 
ENGINE = InnoDB;
