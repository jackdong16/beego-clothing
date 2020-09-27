package controllers

import (
	"beego-clothing/models"

	"github.com/astaxie/beego"
)


type StoreController struct {
	beego.Controller
}


// func (u *UserController) GetAll() {
//     users := models.GetAllUsers()
//     u.Data["json"] = users
//     u.ServeJSON()
// }

// func (c *StoreController) List() {

// 	res :=struct{ Stores[]*models.Stores}{models.storeList.All()}
// 	c.Data["json"] = res
// 	c.ServeJSON()

// 	// var stores models.Store
// 	// models.

// 	// c.TplName = "store/list.tpl"
// }

func (c *StoreController) Create() {
	models.Init()
	c.TplName = "store/list.tpl"

	flash := beego.NewFlash()
	flash.Notice("stores added!")
	flash.Store(&c.Controller)
	c.Redirect("/store/list", 302)
}