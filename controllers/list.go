package controllers

import (
	"github.com/astaxie/beego"
)

type ListController struct {
	beego.Controller
}

func (c *ListController) Get() {
	c.TplName = "list.tpl"
}
