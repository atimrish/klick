package postgres

func UpUsersTable() string {
	query :=
		`CREATE TABLE IF NOT EXISTS users (
    		id SERIAL PRIMARY KEY,
    		surname VARCHAR(30),
    		name VARCHAR(30),
    		login VARCHAR(20) UNIQUE,
    		password VARCHAR(255),
    		email VARCHAR(40) UNIQUE,
    		photo VARCHAR(255),
    		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP NULL DEFAULT NULL
	)`

	return query
}

func DownUsersTable() string {
	query := `DROP TABLE IF EXISTS users`
	return query
}
