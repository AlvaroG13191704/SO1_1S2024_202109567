-- Create database
CREATE DATABASE IF NOT EXISTS Proyecto1;
-- Use database
USE Proyecto1;
-- Create table ram
CREATE TABLE IF NOT EXISTS RAM (
  id INT AUTO_INCREMENT PRIMARY KEY,
  total VARCHAR(10) NOT NULL,
  free VARCHAR(10) NOT NULL,
  used_ram VARCHAR(10) NOT NULL,
  percentage_use VARCHAR(10) NOT NULL,
  date_time DATETIME NOT NULL
);

-- Create table cpu
CREATE TABLE IF NOT EXISTS CPU (
  id INT AUTO_INCREMENT PRIMARY KEY,
  total_cpu VARCHAR(10) NOT NULL,
  percentage_use VARCHAR(10) NOT NULL,
  date_time DATETIME NOT NULL
  processes JSON NOT NULL
);