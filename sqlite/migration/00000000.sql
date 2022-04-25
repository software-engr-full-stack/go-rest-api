CREATE TABLE users (
	id         INTEGER PRIMARY KEY AUTOINCREMENT,
	name       TEXT NOT NULL,
	email      TEXT UNIQUE,
	created_at TEXT NOT NULL,
	updated_at TEXT NOT NULL
);
