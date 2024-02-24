package notification

import (
	"database/sql"
)

type Create struct {
	UserID        string         `db:"user_id"`
	Title         string         `db:"title"`
	Body          string         `db:"body"`
	LinkUrl       sql.NullString `db:"link_url"`
	PhotoPath     sql.NullString `db:"photo_path"`
	PhotoPathType sql.NullString `db:"photo_path_type"`
	ExpiredAt     sql.NullTime   `db:"expired_at"`
	IdType        sql.NullString `db:"id_type"`
	Type          string         `db:"type"`
	Data          string         `db:"data"`
}
