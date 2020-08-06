package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"nginx-gray-go/utils"
	"nginx-gray-go/vo"
)

var serviceMap = make(map[string]vo.RegisterRequest)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Register() {

	var registerRequest vo.RegisterRequest

	json.Unmarshal(this.Ctx.Input.RequestBody, &registerRequest)

	//if err != nil {
	//	panic("请求参数缺失")
	//}

	serviceMap[registerRequest.Id] = registerRequest

}

func (this *RegisterController) AgentList() {

	resMap := make(map[string]vo.RegisterResponse)

	for key, value := range serviceMap {
		regisResp := vo.RegisterResponse{}

		utils.CopyFields(&value, &regisResp)
		regisResp.CreateIndex = 0
		regisResp.ModifyIndex = 0
		regisResp.EnableTagOverride = false
		regisResp.ProxyDestination = ""

		resMap[key] = regisResp
	}

	this.Data["json"] = resMap
	this.ServeJSON()
}

func (this *RegisterController) CatalogList() {

	resMap := make(map[string]interface{})

	var list [1]string
	list[0] = "secure=false"

	for key, value := range serviceMap {
		key = key
		resMap[value.Name] = list
	}

	var consul [0]interface{}
	resMap["consul"] = consul

	this.Data["json"] = resMap

	this.ServeJSON()
}

func (this *RegisterController) EventList() {

	var list [0]interface{}
	this.Data["json"] = list
	this.ServeJSON()
}

func (this *RegisterController) HealthService() {

	var serviceParam = this.GetString(":name")

	var node vo.Node

	node.Node = "GRS-yyy"
	node.ID = "258e15eb-954b-1cc8-063a-718efdf29613"
	node.Address = "127.0.0.1"
	node.Datacenter = "dc1"
	node.CreateIndex = 0
	node.ModifyIndex = 0
	taggedAddresses := make(map[string]string)
	taggedAddresses["lan"] = "127.0.0.1"
	taggedAddresses["wan"] = "127.0.0.1"
	node.TaggedAddresses = taggedAddresses

	var regisRequest vo.RegisterRequest
	for key, value := range serviceMap {
		key = key
		if value.Name == serviceParam {
			regisRequest = value
		}
	}

	var service vo.Service
	utils.CopyFields(&regisRequest, &service)

	var checkNode vo.CheckNode
	checkNode.Node = "GRS-yyy"
	checkNode.CheckID = "service:" + regisRequest.Id
	checkNode.Name = "Service " + serviceParam + "check"
	checkNode.Status = "passing"
	checkNode.Notes = ""
	checkNode.Output = "HTTP GET http://10.10.100.41:8081/actuator/health: 200  Output: {\"status\":\"UP\"}"
	checkNode.ServiceID = regisRequest.Id
	checkNode.ServiceName = regisRequest.Name
	checkNode.CreateIndex = 0
	checkNode.ModifyIndex = 0

	var checkList [1]interface{}
	checkList[0] = checkNode

	response := make(map[string]interface{})
	response["Node"] = node
	response["Service"] = service
	response["Checks"] = checkList

	var respList [1]interface{}
	respList[0] = response
	this.Data["json"] = respList
	this.ServeJSON()
}
