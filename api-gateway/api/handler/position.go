package handler

import (
	"context"
	cp "gateway/genprotos"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PositionCreate handles the creation of a new position.
// @Summary Create position
// @Description Create a new position
// @Tags position
// @Accept json
// @Produce json
// @Param position body cp.PositionCreateReq true "Position data"
// @Success 200 {object} cp.PositionRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /position [post]
func (h *Handler) PositionCreate(c *gin.Context) {
	var req cp.PositionCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.srvs.Position.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// PositionGetById handles the get a position.
// @Summary Get position
// @Description Get a position
// @Tags position
// @Accept json
// @Produce json
// @Param id path string true "Position ID"
// @Success 200 {object} cp.PositionGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /position/{id} [get]
func (h *Handler) PositionGetById(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	res, err := h.srvs.Position.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get position", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PositionGetAll handles getting all positions.
// @Summary Get all positions
// @Description Get all positions
// @Tags position
// @Accept json
// @Produce json
// @Param department_id query string false "DepartmentId"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} cp.PositionGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /position/all [get]
func (h *Handler) PositionGetAll(c *gin.Context) {
	req := cp.PositionGetAllReq{
		DepartmentId: c.Query("department_id"),
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
	
	res, err := h.srvs.Position.GetAll(context.Background(), &req)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get positions", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PositionUpdate handles updating an existing position.
// @Summary Update position
// @Description Update an existing position
// @Tags position
// @Accept json
// @Produce json
// @Param id path string true "Position ID"
// @Param position body cp.PositionCreateReq true "Updated position data"
// @Success 200 {object} cp.PositionRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Position not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /position/{id} [put]
func (h *Handler) PositionUpdate(c *gin.Context) {
	id := c.Param("id")
	var req cp.PositionUpdateReq
	var position cp.PositionCreateReq

	if err := c.BindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = id
	req.Title = position.Title
	req.Description = position.DepartmentId
	res, err := h.srvs.Position.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update position", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PositionDelete handles deleting a position by ID.
// @Summary Delete position
// @Description Delete a position by ID
// @Tags position
// @Accept json
// @Produce json
// @Param id path string true "Position ID"
// @Success 200 {object} string "Position deleted"
// @Failure 400 {object} string "Invalid position ID"
// @Failure 404 {object} string "Position not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /position/{id} [delete]
func (h *Handler) PositionDelete(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	_, err := h.srvs.Position.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete position", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Position deleted"})
}

// PositionGetByDepartment handles getting all positions.
// @Summary Get position by department
// @Description Get position by department
// @Tags department
// @Accept json
// @Produce json
// @Param department_id query string false "DepartmentId"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} cp.PositionGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /department/{id}/position [get]
func (h *Handler) PositionGetByDepartment(c *gin.Context) {
	req := cp.PositionGetAllReq{
		DepartmentId: c.Query("department_id"),
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
	
	res, err := h.srvs.Position.GetAll(context.Background(), &req)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get positions", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}