package controller

import (
	"github.com/gin-gonic/gin"
)

type APIController struct {
	BaseController
	inboundController *InboundController
	serverController  *ServerController
}

func NewAPIController(g *gin.RouterGroup) *APIController {
	a := &APIController{}
	a.serverController = NewServerControllerForApi()
	a.initRouter(g)
	return a
}

func (a *APIController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/xui/API/inbounds")
	g.Use(a.validate)

	g.GET("/", a.inbounds)
	// g.GET("/get/:id", a.inbound)
	g.POST("/get/:name", a.getInbound)
	// g.GET("/getClientTraffics/:email", a.getClientTraffics)
	g.POST("/add", a.addSingleInbound)
	// g.POST("/del/:id", a.delInbound)
	g.POST("/del/:name", a.delInboundByName)
	g.POST("/update/:name", a.updateInboundByName)
	// g.POST("/addClient", a.addInboundClient)
	// g.POST("/:id/delClient/:clientId", a.delInboundClient)
	g.POST("/updateClient/:clientId", a.updateInboundClient)
	// g.POST("/:id/resetClientTraffic/:email", a.resetClientTraffic)
	// g.POST("/resetAllTraffics", a.resetAllTraffics)
	// g.POST("/resetAllClientTraffics/:id", a.resetAllClientTraffics)
	// g.POST("/delDepletedClients/:id", a.delDepletedClients)
	g.POST("/status", a.trafficStatus)

	a.inboundController = NewInboundController(g)
}

func (a *APIController) trafficStatus(c *gin.Context) {
	a.serverController.trafficStatus(c)
}

func (a *APIController) inbounds(c *gin.Context) {
	a.inboundController.getInbounds(c)
}
func (a *APIController) getInbound(c *gin.Context) {
	a.inboundController.getInboundStats(c)
}
func (a *APIController) inbound(c *gin.Context) {
	a.inboundController.getInbound(c)
}
func (a *APIController) getClientTraffics(c *gin.Context) {
	a.inboundController.getClientTraffics(c)
}
func (a *APIController) addInbound(c *gin.Context) {
	a.inboundController.addInbound(c)
}

func (a *APIController) addSingleInbound(c *gin.Context) {
	a.inboundController.addSingleInbound(c)
}

func (a *APIController) delInboundByName(c *gin.Context) {
	a.inboundController.delInboundByName(c)
}

func (a *APIController) delInbound(c *gin.Context) {
	a.inboundController.delInbound(c)
}
func (a *APIController) updateInboundByName(c *gin.Context) {
	a.inboundController.updateInboundByName(c)
}
func (a *APIController) addInboundClient(c *gin.Context) {
	a.inboundController.addInboundClient(c)
}
func (a *APIController) delInboundClient(c *gin.Context) {
	a.inboundController.delInboundClient(c)
}
func (a *APIController) updateInboundClient(c *gin.Context) {
	a.inboundController.updateInboundClient(c)
}
func (a *APIController) resetClientTraffic(c *gin.Context) {
	a.inboundController.resetClientTraffic(c)
}
func (a *APIController) resetAllTraffics(c *gin.Context) {
	a.inboundController.resetAllTraffics(c)
}
func (a *APIController) resetAllClientTraffics(c *gin.Context) {
	a.inboundController.resetAllClientTraffics(c)
}
func (a *APIController) delDepletedClients(c *gin.Context) {
	a.inboundController.delDepletedClients(c)
}
