-- Create database
CREATE DATABASE IF NOT EXISTS Proyecto1;
-- Use database
USE Proyecto1;
-- Create table ram 
CREATE TABLE IF NOT EXISTS ram (
  id INT AUTO_INCREMENT PRIMARY KEY,
  date_time DATETIME NOT NULL,
  free INT NOT NULL,
  total INT NOT NULL
  in_use INT NOT NULL
  percentage_use INT NOT NULL
);

-- Create table cpu
CREATE TABLE IF NOT EXISTS cpu (
  id INT AUTO_INCREMENT PRIMARY KEY,
  date_time DATETIME NOT NULL,
  percentage_use INT NOT NULL
);