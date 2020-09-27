package models

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    // _ "github.com/go-sql-driver/mysql" // import your required driver
    _ "github.com/mattn/go-sqlite3"
)

// var storeList *StoreList

// Model Struct
type Store struct {
    Id   int
	Name string `orm:"size(100)"`
}

// func storeManager() *StoreList {
//     return &StoreList{}
// }


func Init() {
    // register model
    orm.RegisterModel(new(Store))

    // set default database
    // orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8", 30)
    orm.RegisterDataBase("default", "sqlite3", "data.db")

    Create()
}

func Create() {
    o := orm.NewOrm()

    store := Store{Name: "總倉"}
    // insert
    id, err := o.Insert(&store)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    store = Store{Name: "虹亭"}
    // insert
    id, err = o.Insert(&store)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    store = Store{Name: "彩意房"}
    // insert
    id, err = o.Insert(&store)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    store = Store{Name: "靖衣坊"}
    // insert
    id, err = o.Insert(&store)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    store = Store{Name: "逸琳春"}
    // insert
    id, err = o.Insert(&store)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    store = Store{Name: "唐庄"}
    // insert
    id, err = o.Insert(&store)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    store = Store{Name: "蓉蓉"}
    // insert
    id, err = o.Insert(&store)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    store = Store{Name: "半掩"}
    // insert
    id, err = o.Insert(&store)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    store = Store{Name: "名古屋"}
    // insert
    id, err = o.Insert(&store)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    store = Store{Name: "鳳凰豆"}
    // insert
    id, err = o.Insert(&store)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)
    
    // update
    // user.Name = "updateproductname"
    // num, err := o.Update(&product)
    // fmt.Printf("NUM: %d, ERR: %v\n", num, err)

    // // read one
    // u := Product{Id: product.Id}
    // err = o.Read(&u)
    // fmt.Printf("ERR: %v\n", err)

    // // delete
    // num, err = o.Delete(&u)
    // fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
