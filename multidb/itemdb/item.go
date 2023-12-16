package itemdb

type Item struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func NewItem(id int, name string) *Item {
	return &Item{
		ID:   id,
		Name: name,
	}
}
