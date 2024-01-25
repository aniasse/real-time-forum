
--Creation de la base de donnée
--CREATE DATABASE IF NOT EXISTS db;

--Utilisation de la base de donnée
--USE db;

-- Creation d'une table "utilisateur"
CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nickname TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    email TEXT NOT NULL,
    gender TEXT NOT NULL,
    age INTEGER NOT NULL,
    password VARCHAR(254) NOT NULL
)
