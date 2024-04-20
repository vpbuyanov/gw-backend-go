package common

const (
	CreateComment = `INSERT INTO "comment" (creator_id, topic_id) VALUES (?, ?) RETURNING *`
	CreateTopic   = `INSERT INTO "forum_topic" (title, creator_id) VALUES(?, ?) RETURNING *`
	DeleteTopic   = `DELETE FROM "forum_topic" WHERE id=? RETURNING *`
	DeleteComment = `DELETE FROM "comment" WHERE id=? RETURNING *`
	UpdateTopic   = `UPDATE "forum_topic" SET title=? WHERE id=? RETURNING *`
	UpdateComment = `UPDATE "comment" SET text=? WHERE id=? RETURNING *`
	SelectTopic   = `SELECT * FROM "forum_topic" WHERE id=?`
	SelectComment = `SELECT * FROM "comment" WHERE id=?`
)
