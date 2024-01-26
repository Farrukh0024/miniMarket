package api

import (
	"connected/api/handler"
	"connected/storage"

	"github.com/gin-gonic/gin"
)

func New(store storage.IStorage) *gin.Engine {
	h := handler.New(store)

	r := gin.New()

	//branch
	r.POST("/branch", h.CreateBranch)
	r.GET("/branch/:id", h.GetByID)
	r.GET("/branchs", h.GetList)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.Delete)

	//stafftarif
	r.POST("staff_tarif", h.CreateStaff_Tarif)
	r.GET("staff_tarif/:id", h.GetByIdStaff_Tarif)
	r.GET("/staff_tarifs", h.GetListStaff_Tarif)
	r.PUT("/staff_tarif/:id", h.UpdateStaff_Tarif)
	r.DELETE("/staff_tarif/:id", h.DeleteStaff_Tarif)

	//staf
	r.POST("/staff", h.CreateStaff)
	r.GET("/staff/:id", h.GetByIdStaff)
	r.GET("/staffs", h.GetListStaff)
	r.PUT("/staff/:id", h.UpdateStaff)
	r.DELETE("/staff/:id", h.DeleteStaff)

	//sale
	r.POST("/sale", h.CreateSale)
	r.GET("/sale/:id", h.GetByIDSales)
	r.GET("/sales", h.GetListSales)
	r.PUT("/sale/:id", h.UpdateSales)
	r.DELETE("/sale/:id", h.DeleteSales)

	return r
}
