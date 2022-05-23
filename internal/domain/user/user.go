package user

type User struct {
	Id   int64  `db:"id,omitempty"`
	Name string `db:"name"`
}
