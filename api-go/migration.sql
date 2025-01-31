CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(50) NOT NULL,
    user_email VARCHAR(50) UNIQUE NOT NULL,
    user_password VARCHAR(255) NOT NULL,
    user_created_at TIMESTAMP WITH TIME ZONE,
    user_updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS tasks (
    task_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    task_title VARCHAR(255) NOT NULL,
    task_status VARCHAR(50) NOT NULL DEFAULT 'on_going',
    task_created_at TIMESTAMP WITH TIME ZONE,
    task_updated_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) 
);

CREATE INDEX tasks_status_idx
ON tasks (task_status);