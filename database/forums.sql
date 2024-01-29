-- Table des utilisateurs
CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nickname TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    email TEXT NOT NULL,
    gender TEXT NOT NULL,
    age INTEGER NOT NULL,
    session TEXT NOT NULL,
    password VARCHAR(254) NOT NULL
);

-- Table des publications (Post)
CREATE TABLE IF NOT EXISTS Post (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    title VARCHAR(254) NULL,
    content VARCHAR(254) NULL,
    image VARCHAR(254) NULL,
    date_created TIMESTAMP NULL
);

-- Table de liaison entre les publications et les catégories (Post_Category)
CREATE TABLE IF NOT EXISTS Post_Category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    cat_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL
);

-- Table des sessions
CREATE TABLE IF NOT EXISTS Session (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email VARCHAR(250),
    end_date TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Commentary (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    com_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    content VARCHAR(250) NULL,
    date_created TIMESTAMP NULL
);