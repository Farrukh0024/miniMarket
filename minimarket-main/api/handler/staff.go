package handler

import (
	"connected/api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//create staff
func (h Handler) CreateStaff(c *gin.Context) {
	createStaff := models.CreateStaff{}

	if err := c.ShouldBindJSON(&createStaff); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Staff().CreateStaff(createStaff)
	if err != nil {
		handleResponse(c, "error while creating staff", http.StatusInternalServerError, err)
		return
	}

	staff, err := h.storage.Staff().GetByIdStaff(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting staff by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, staff)
}

//get by id  staff
func (h Handler) GetByIdStaff(c *gin.Context) {
	var err error

	uid := c.Param("id")

	staff, err := h.storage.Staff().GetByIdStaff(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting staff by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, staff)
}

//getlist staff
func (h Handler) GetListStaff(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)

	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	resp, err := h.storage.Staff().GetListStaff(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting staff", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

//create staff
func (h Handler) UpdateStaff(c *gin.Context) {
	updateStaff := models.UpdateStaff{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateStaff.ID = uid

	if err := c.ShouldBindJSON(&updateStaff); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Staff().UpdateStaffs(updateStaff)
	if err != nil {
		handleResponse(c, "error while updating staff", http.StatusInternalServerError, err.Error())
		return
	}

	branch, err := h.storage.Staff().GetByIdStaff(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting staff by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, branch)
}

//delete staff
func (h Handler) DeleteStaff(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Staff().DeleteStaff(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting staff by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}
