package entity

type CustomEntity struct {
	ID       string
	Custom   string
	FullName string
	Email    string
}

func (CustomEntity) TableName() string {
	return "custom"
}
