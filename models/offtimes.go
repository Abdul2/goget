package models

import (
	"github.com/astaxie/beego/orm"
)

type Offtimes struct {
	ID   int64
	Day  string
	Moth string
	Year string
	No   string
}

func init() {

	orm.RegisterModel(new(Offtimes))

}
