package entity

type Exmpl struct {
	ID    uint    // Standard field for the primary key
	Name  string  // A regular string field
	Email *string // A pointer to a string, allowing for null values
	Age   uint8
}

func (Exmpl) TableName() string {
	return "exmpl" // Specify the singular form here
}
