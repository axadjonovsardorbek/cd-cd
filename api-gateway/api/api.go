package api

import (
	"gateway/api/handler"
	"gateway/grpc/clients"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "gateway/docs"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(connC *clients.GrpcClients) *gin.Engine {
	h := handler.NewHandler(connC)

	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	company := router.Group("/company")
	company.POST("/", h.CompanyCreate)
	company.GET("/:id", h.CompanyGetById)
	company.GET("/all", h.CompanyGetAll)
	company.PUT("/:id", h.CompanyUpdate)
	company.DELETE("/:id", h.CompanyDelete)

	department := router.Group("/department")
	department.POST("/", h.DepartmentCreate)
	department.GET("/:id", h.DepartmentGetById)
	department.GET("/all", h.DepartmentGetAll)
	department.PUT("/:id", h.DepartmentUpdate)
	department.DELETE("/:id", h.DepartmentDelete)
	department.GET("/:id/position", h.PositionGetByDepartment)

	position := router.Group("/position")
	position.POST("/", h.PositionCreate)
	position.GET("/:id", h.PositionGetById)
	position.GET("/all", h.PositionGetAll)
	position.PUT("/:id", h.PositionUpdate)
	position.DELETE("/:id", h.PositionDelete)

	resume := router.Group("/resume")
	resume.POST("/", h.ResumeCreate)
	resume.GET("/:id", h.ResumeGetById)
	resume.GET("/all", h.ResumeGetAll)
	resume.PUT("/:id", h.ResumeUpdate)
	resume.DELETE("/:id", h.ResumeDelete)

	return router
}
