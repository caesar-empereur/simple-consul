package routers

import (
	"github.com/astaxie/beego"
	"nginx-gray-go/controllers"
)

func init() {

	beego.Router("/v1/agent/service/register", &controllers.RegisterController{}, "put:Register")
	beego.Router("/v1/agent/services", &controllers.RegisterController{}, "get:AgentList")
	beego.Router("/v1/catalog/services", &controllers.RegisterController{}, "get:CatalogList")
	beego.Router("/v1/event/list", &controllers.RegisterController{}, "get:EventList")
	beego.Router("/v1/health/service/:name", &controllers.RegisterController{}, "get:HealthService")

}
