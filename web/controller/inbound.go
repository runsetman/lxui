package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"x-ui/database/model"
	"x-ui/logger"
	"x-ui/web/global"
	"x-ui/web/service"
	"x-ui/web/session"

	"github.com/gin-gonic/gin"
)

type InboundController struct {
	inboundService service.InboundService
	xrayService    service.XrayService
}

func NewInboundController(g *gin.RouterGroup) *InboundController {
	a := &InboundController{}
	a.initRouter(g)
	a.startTask()
	return a
}

func (a *InboundController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/inbound")

	g.POST("/list", a.getInbounds)
	g.POST("/add", a.addInbound)
	g.POST("/del/:id", a.delInbound)
	g.POST("/update/:id", a.updateInbound)
	g.POST("/addClient", a.addInboundClient)
	g.POST("/:id/delClient/:clientId", a.delInboundClient)
	g.POST("/updateClient/:clientId", a.updateInboundClient)
	g.POST("/:id/resetClientTraffic/:email", a.resetClientTraffic)
	g.POST("/resetAllTraffics", a.resetAllTraffics)
	g.POST("/resetAllClientTraffics/:id", a.resetAllClientTraffics)
	g.POST("/delDepletedClients/:id", a.delDepletedClients)

}

func (a *InboundController) startTask() {
	webServer := global.GetWebServer()
	c := webServer.GetCron()
	c.AddFunc("@every 10s", func() {
		if a.xrayService.IsNeedRestartAndSetFalse() {
			err := a.xrayService.RestartXray(false)
			if err != nil {
				logger.Error("restart xray failed:", err)
			}
		}
	})
}

func (a *InboundController) getInbounds(c *gin.Context) {
	user := session.GetLoginUser(c)
	inbounds, err := a.inboundService.GetInbounds(user.Id)
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.toasts.obtain"), err)
		return
	}
	fmt.Println(inbounds)
	jsonObj(c, inbounds, nil)
}

func (a *InboundController) getInboundStats(c *gin.Context) {
	name := c.Param("name")
	inbound, err := a.inboundService.GetInboundByName(name)

	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.toasts.obtain"), err)
		return
	}

	upload := int64(0)
	download := int64(0)
	totalTraffic := int64(0)
	expiry := int64(0)

	settings := map[string][]model.Client{}
	json.Unmarshal([]byte(inbound.Settings), &settings)
	if settings != nil {
		// jsonMsg(c, I18n(c, "settings is nil"), err)
		// return
		clients := settings["clients"]
		if clients != nil {
			// jsonMsg(c, I18n(c, "no clients"), err)
			// return
			totalTraffic = clients[0].TotalGB
		}
	}

	fmt.Println(inbound)

	if len(inbound.ClientStats) != 0 {
		clientStat := inbound.ClientStats[0]
		upload = clientStat.Up
		download = clientStat.Down
		expiry = clientStat.ExpiryTime
	}

	inboundSimpleStats := struct {
		Up           int64 `json:"up"`
		Down         int64 `json:"down"`
		TotalTraffic int64 `json:"total"`
		Expiry       int64 `json:"expiry"`
	}{
		Up:           upload,
		Down:         download,
		Expiry:       expiry,
		TotalTraffic: totalTraffic,
	}

	jsonObj(c, inboundSimpleStats, nil)
}

func (a *InboundController) getInbound(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonMsg(c, I18n(c, "get"), err)
		return
	}
	inbound, err := a.inboundService.GetInbound(id)
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.toasts.obtain"), err)
		return
	}
	jsonObj(c, inbound, nil)
}
func (a *InboundController) getClientTraffics(c *gin.Context) {
	email := c.Param("email")
	clientTraffics, err := a.inboundService.GetClientTrafficByEmail(email)
	if err != nil {
		jsonMsg(c, "Error getting traffics", err)
		return
	}
	jsonObj(c, clientTraffics, nil)
}

func (a *InboundController) addInbound(c *gin.Context) {
	// jsonData, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// Handle error
	// }
	// fmt.Println(string(jsonData))
	// fmt.Println("nima1")
	inbound := &model.Inbound{}
	err := c.ShouldBind(inbound)
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.addTo"), err)
		return
	}
	// user := session.GetLoginUser(c)
	inbound.UserId = 1
	inbound.Enable = true
	inbound.Tag = fmt.Sprintf("inbound-%v", inbound.Port)
	inbound, err = a.inboundService.AddInbound(inbound)
	jsonMsgObj(c, I18n(c, "pages.inbounds.addTo"), inbound, err)
	if err == nil {
		a.xrayService.SetToNeedRestart()
	}
}

func (a *InboundController) addSingleInbound(c *gin.Context) {
	// jsonData, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// Handle error
	// }
	// fmt.Println(string(jsonData))
	inbound := &model.Inbound{}
	err := c.ShouldBind(inbound)
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.addTo"), err)
		return
	}
	// user := session.GetLoginUser(c)
	inbound.UserId = 1
	inbound.Enable = true
	inbound.Tag = fmt.Sprintf("inbound-%v", inbound.Port)
	inbound, err = a.inboundService.AddSingleInbound(inbound)
	jsonMsgObj(c, I18n(c, "pages.inbounds.addTo"), inbound.Port, err)
	if err == nil {
		a.xrayService.SetToNeedRestart()
	}
}

func (a *InboundController) delInboundByName(c *gin.Context) {
	name := c.Param("name")
	err := a.inboundService.DelInboundByName(name)
	jsonMsgObj(c, I18n(c, "delete"), name, err)
	if err == nil {
		a.xrayService.SetToNeedRestart()
	}
}

func (a *InboundController) delInbound(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonMsg(c, I18n(c, "delete"), err)
		return
	}
	err = a.inboundService.DelInbound(id)
	jsonMsgObj(c, I18n(c, "delete"), id, err)
	if err == nil {
		a.xrayService.SetToNeedRestart()
	}
}

func (a *InboundController) updateInbound(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.revise"), err)
		return
	}
	inbound := &model.Inbound{
		Id: id,
	}
	err = c.ShouldBind(inbound)
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.revise"), err)
		return
	}
	inbound, err = a.inboundService.UpdateInbound(inbound)
	jsonMsgObj(c, I18n(c, "pages.inbounds.revise"), inbound, err)
	if err == nil {
		a.xrayService.SetToNeedRestart()
	}
}

func (a *InboundController) addInboundClient(c *gin.Context) {
	data := &model.Inbound{}
	err := c.ShouldBind(data)
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.revise"), err)
		return
	}

	err = a.inboundService.AddInboundClient(data)
	if err != nil {
		jsonMsg(c, "Something went wrong!", err)
		return
	}
	jsonMsg(c, "Client(s) added", nil)
	if err == nil {
		a.xrayService.SetToNeedRestart()
	}
}

func (a *InboundController) delInboundClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.revise"), err)
		return
	}
	clientId := c.Param("clientId")

	err = a.inboundService.DelInboundClient(id, clientId)
	if err != nil {
		jsonMsg(c, "Something went wrong!", err)
		return
	}
	jsonMsg(c, "Client deleted", nil)
	if err == nil {
		a.xrayService.SetToNeedRestart()
	}
}

func (a *InboundController) updateInboundClient(c *gin.Context) {
	clientId := c.Param("clientId")

	inbound := &model.Inbound{}
	err := c.ShouldBind(inbound)
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.revise"), err)
		return
	}

	err = a.inboundService.UpdateInboundClient(inbound, clientId)
	if err != nil {
		jsonMsg(c, "Something went wrong!", err)
		return
	}
	jsonMsg(c, "Client updated", nil)
	if err == nil {
		a.xrayService.SetToNeedRestart()
	}
}

func (a *InboundController) resetClientTraffic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.revise"), err)
		return
	}
	email := c.Param("email")

	err = a.inboundService.ResetClientTraffic(id, email)
	if err != nil {
		jsonMsg(c, "Something went wrong!", err)
		return
	}
	jsonMsg(c, "traffic reseted", nil)
	if err == nil {
		a.xrayService.SetToNeedRestart()
	}
}

func (a *InboundController) resetAllTraffics(c *gin.Context) {
	err := a.inboundService.ResetAllTraffics()
	if err != nil {
		jsonMsg(c, "Something went wrong!", err)
		return
	}
	jsonMsg(c, "All traffics reseted", nil)
}

func (a *InboundController) resetAllClientTraffics(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.revise"), err)
		return
	}

	err = a.inboundService.ResetAllClientTraffics(id)
	if err != nil {
		jsonMsg(c, "Something went wrong!", err)
		return
	}
	jsonMsg(c, "All traffics of client reseted", nil)
}

func (a *InboundController) delDepletedClients(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonMsg(c, I18n(c, "pages.inbounds.revise"), err)
		return
	}
	err = a.inboundService.DelDepletedClients(id)
	if err != nil {
		jsonMsg(c, "Something went wrong!", err)
		return
	}
	jsonMsg(c, "All delpeted clients are deleted", nil)
}
