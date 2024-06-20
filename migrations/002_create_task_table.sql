CREATE TABLE IF NOT EXISTS tasks (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` TEXT NOT NULL,
    `description` TEXT,
    `responsible` TEXT,
    `start_date` DATE,
    `end_date` DATE,
    `status` TEXT NOT NULL,
    `project_id` INTEGER NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);
-- Indexes to optimize frequent queries
CREATE INDEX IF NOT EXISTS idx_project_id ON tasks (project_id);
CREATE INDEX IF NOT EXISTS idx_status ON tasks (status);