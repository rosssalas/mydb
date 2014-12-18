package mydb

import (
	"fmt"
)

var Db int

func Init() {
	Db = 23
	fmt.Println("Initialized:",Db)
}

