package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego-api-server/controllers:MartabakController"] = append(beego.GlobalControllerRouter["beego-api-server/controllers:MartabakController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-api-server/controllers:MartabakController"] = append(beego.GlobalControllerRouter["beego-api-server/controllers:MartabakController"],
        beego.ControllerComments{
            Method: "FindAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-api-server/controllers:MartabakController"] = append(beego.GlobalControllerRouter["beego-api-server/controllers:MartabakController"],
        beego.ControllerComments{
            Method: "Find",
            Router: "/:martabakID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-api-server/controllers:MartabakController"] = append(beego.GlobalControllerRouter["beego-api-server/controllers:MartabakController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:martabakID",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-api-server/controllers:MartabakController"] = append(beego.GlobalControllerRouter["beego-api-server/controllers:MartabakController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:martabakID",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
