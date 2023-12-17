package postgres

func UpFriendsTable() string {
	query :=
		`CREATE TABLE IF NOT EXISTS friends (
    		id SERIAL PRIMARY KEY,
    		user_id INT NOT NULL,
    		friend_id INT NOT NULL,
    		status VARCHAR(20),
    		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    		updated_at TIMESTAMP NULL DEFAULT NULL,
    		FOREIGN KEY (user_id) REFERENCES users (id),
    		FOREIGN KEY (friend_id) REFERENCES users (id)
        )
        `

	return query
}

func DownFriendsTable() string {
	query := `DROP TABLE IF EXISTS friends`

	return query
}
