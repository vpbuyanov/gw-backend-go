package postgres

const (
	CreateTopic   = `INSERT INTO "forum_topic" (title, creator_id) VALUES(?, ?)`
	CreateUser    = `INSERT INTO "user" (name, email, hash_pass) VALUES(?, ?, ?)`
	CreateComment = `INSERT INTO "comment" (creator_id, topic_id) VALUES (?, ?)`
)

const (
	DeleteTopic   = `DELETE FROM "forum_topic" WHERE id=?`
	DeleteUser    = `DELETE FROM "user" WHERE id=?`
	DeleteComment = `DELETE FROM "comment" WHERE id=?`
)

const (
	UpdateTopic   = `UPDATE "forum_topic" SET title=? WHERE id=?`
	UpdateUser    = `UPDATE "user" SET name=?, email=?, hash_pass=? WHERE id=?`
	UpdateComment = `UPDATE "comment" SET text=? WHERE id=?`
)

const (
	SelectTopic   = `SELECT * FROM "forum_topic" WHERE id=?`
	SelectUser    = `SELECT * FROM "user" WHERE id=?`
	SelectComment = `SELECT * FROM "comment" WHERE id=?`
)
