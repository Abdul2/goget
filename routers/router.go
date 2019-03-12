package routers

import (
	//"s3/controllers"

	"github.com/Abdul2/goget/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/list", &controllers.ListController{})
	beego.Router("/offtimes", &controllers.OfftimesController{})

}
