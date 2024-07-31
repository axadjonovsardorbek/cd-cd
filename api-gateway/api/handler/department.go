package handler

import (
	"context"
	cp "gateway/genprotos"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DepartmentCreate handles the creation of a new department.
// @Summary Create department
// @Description Create a new department
// @Tags department
// @Accept json
// @Produce json
// @Param department body cp.DepartmentCreateReq true "Department data"
// @Success 200 {object} cp.DepartmentRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /department [post]
func (h *Handler) DepartmentCreate(c *gin.Context) {
	var req cp.DepartmentCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.srvs.Department.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// DepartmentGetById handles the get a department.
// @Summary Get department
// @Description Get a department
// @Tags department
// @Accept json
// @Produce json
// @Param id path string true "Department ID"
// @Success 200 {object} cp.DepartmentGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /department/{id} [get]
func (h *Handler) DepartmentGetById(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	res, err := h.srvs.Department.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get department", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DepartmentGetAll handles getting all departments.
// @Summary Get all departments
// @Description Get all departments
// @Tags department
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Param company_id query string false "CompanyId"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} cp.DepartmentGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /department/all [get]
func (h *Handler) DepartmentGetAll(c *gin.Context) {
	req := cp.DepartmentGetAllReq{
		Name: c.Query("name"),
		CompanyId: c.Query("company_id"),
		Filter: &cp.Pagination{},
	}

	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	var limit, offset int
	var err error
	if limitStr == "" {
		limit = 0
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}
	if offsetStr == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
	}

	filter := cp.Pagination{
		Limit:  int64(limit),
		Offset: int64(offset),
	}

	req.Filter.Limit = filter.Limit
	req.Filter.Offset = filter.Offset
	
	res, err := h.srvs.Department.GetAll(context.Background(), &req)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get departments", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DepartmentUpdate handles updating an existing department.
// @Summary Update department
// @Description Update an existing department
// @Tags department
// @Accept json
// @Produce json
// @Param id path string true "Department ID"
// @Param department body cp.DepartmentCreateReq true "Updated department data"
// @Success 200 {object} cp.DepartmentRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Department not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /department/{id} [put]
func (h *Handler) DepartmentUpdate(c *gin.Context) {
	id := c.Param("id")
	var req cp.DepartmentUpdateReq
	var department cp.DepartmentCreateReq

	if err := c.BindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = id
	req.Name = department.Name
	req.CompanyId = department.CompanyId
	res, err := h.srvs.Department.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update department", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DepartmentDelete handles deleting a department by ID.
// @Summary Delete department
// @Description Delete a department by ID
// @Tags department
// @Accept json
// @Produce json
// @Param id path string true "Department ID"
// @Success 200 {object} string "Department deleted"
// @Failure 400 {object} string "Invalid department ID"
// @Failure 404 {object} string "Department not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /department/{id} [delete]
func (h *Handler) DepartmentDelete(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	_, err := h.srvs.Department.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete department", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Department deleted"})
}
