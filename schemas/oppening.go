package schemas

import (
	"gorm.io/gorm"
)

type Oppening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Ling     string
	Salary   int64
}
