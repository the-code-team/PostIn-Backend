package controllers

import "github.com/gin-gonic/gin"

func GetPropouses(context *gin.Context) {

}

func CreatePropouse(context *gin.Context) {

}

func AcceptPropose(context *gin.Context) {

}

func RejectPropose(context *gin.Context) {

}

func ProposesController(r *gin.Engine) {
	proposes := r.Group("/proposes")
	{
		proposes.GET("/", GetPropouses)
		proposes.POST("/create", CreatePropouse)
		proposes.POST("/accept/:id", AcceptPropose)
		proposes.POST("/reject/:id", RejectPropose)
	}
}
