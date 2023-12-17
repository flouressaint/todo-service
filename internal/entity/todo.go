package entity

type Todo struct {
	Id        int    `db:"id"`
	Title     string `db:"title"`
	Completed bool   `db:"completed"`
}
