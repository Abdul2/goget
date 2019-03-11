package controllers

import (
	"github.com/astaxie/beego"
)

type OfftimesController struct {
	beego.Controller
}

func (c *OfftimesController) Get() {
	c.TplName = "offtimes.tpl"
}
