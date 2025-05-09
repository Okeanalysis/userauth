CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) UNIQUE NOT NULL,
	password_hash VARCHAR(255) NOT NUll,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	
);

CREATE INDEX idx_user_email ON users(email);



CREATE TABLE project (
	id SERIAL PRIMARY KEY,
	user_id INT REFERENCES users(id) ON DELETE CASCADE,
	title VARCHAR(255) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP

);



CREATE TABLE TODO (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL, 
	project_id INT REFERENCES project(id) ON DELETE CASCADE,
	completed Boolean DEFAULT FALSE,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);