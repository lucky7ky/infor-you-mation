package routers

import (
	"github.com/astaxie/beego"
	"github.com/lucky7ky/infor-you-mation/controllers"
)

func init() {
	beego.Router("/cardlist", &controllers.CardListController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/carddetail", &controllers.CardDetailController{})
	beego.Router("/trend", &controllers.TrendController{})
}
