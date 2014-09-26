package main

import (
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
	//"os"
	// _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

type ComapnyType struct {
	Id          int
	ComapnyName string
}

func main() {
	db := mysql.New("tcp", "", "devserver1:3306", "root", "lotus", "appserver_core")

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	// rows, _, err := db.Query("SELECT id,company_name FROM companys")
	// if err != nil {
	// 	panic(err)
	// }

	// for _, row := range rows {
	// 	for _, col := range row {
	// 		if col == nil {
	// 			// col has NULL value
	// 		} else {
	// 			// Do something with text in col (type []byte)
	// 			fmt.Print(string(col.([]byte)))
	// 		}
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Println("")
	// }

	res, err := db.Start("SELECT id,company_name FROM companys")
	if err != nil {
		panic(err)
	}

	// Get result from first select
	for {
		row, err := res.GetRow()
		if err != nil {
			panic(err)
		}
		if row == nil {
			break
		}
		fmt.Print(row.Int(0), " ")
		fmt.Print(row.Str(1))
		fmt.Println("")

	}

	// You can get specific value from a row
	//val1 := row[1].([]byte)

	// You can use it directly if conversion isn't needed
	//os.Stdout.Write(val1)

	// // You can get converted value
	// number := row.Int(0)      // Zero value
	// str := row.Str(1)         // First value
	// bignum := row.MustUint(2) // Second value

	// // You may get values by column name
	// first := res.Map("FirstColumn")
	// second := res.Map("SecondColumn")
	// val1, val2 := row.Int(first), row.Str(second)
	//}
}
