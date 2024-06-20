CREATE TABLE IF NOT EXISTS projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` TEXT NOT NULL,
    `description` TEXT,
    `start_date` DATE,
    `end_date` DATE,
    `status` TEXT NOT NULL
);