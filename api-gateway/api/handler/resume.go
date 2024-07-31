package handler

import (
	"context"
	cp "gateway/genprotos"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ResumeCreate handles the creation of a new resume.
// @Summary Create resume
// @Description Create a new resume
// @Tags resume
// @Accept json
// @Produce json
// @Param resume body cp.ResumeCreateReq true "Resume data"
// @Success 200 {object} cp.ResumeRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /resume [post]
func (h *Handler) ResumeCreate(c *gin.Context) {
	var req cp.ResumeCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.srvs.Resume.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// ResumeGetById handles the get a resume.
// @Summary Get resume
// @Description Get a resume
// @Tags resume
// @Accept json
// @Produce json
// @Param id path string true "Resume ID"
// @Success 200 {object} cp.ResumeGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /resume/{id} [get]
func (h *Handler) ResumeGetById(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	res, err := h.srvs.Resume.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get resume", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ResumeGetAll handles getting all resumes.
// @Summary Get all resumes
// @Description Get all resumes
// @Tags resume
// @Accept json
// @Produce json
// @Param user_id query string false "UserId"
// @Param education query string false "Education"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} cp.ResumeGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /resume/all [get]
func (h *Handler) ResumeGetAll(c *gin.Context) {
	req := cp.ResumeGetAllReq{
		UserId: c.Query("user_id"),
		Education: c.Query("education"),
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
	
	res, err := h.srvs.Resume.GetAll(context.Background(), &req)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get resumes", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ResumeUpdate handles updating an existing resume.
// @Summary Update resume
// @Description Update an existing resume
// @Tags resume
// @Accept json
// @Produce json
// @Param id path string true "Resume ID"
// @Param position body cp.ResumeCreateReq true "Updated resume data"
// @Success 200 {object} cp.ResumeRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Resume not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /resume/{id} [put]
func (h *Handler) ResumeUpdate(c *gin.Context) {
	id := c.Param("id")
	var req cp.ResumeUpdateReq
	var resume cp.ResumeCreateReq

	if err := c.BindJSON(&resume); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = id
	req.Education = resume.Education
	res, err := h.srvs.Resume.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update resume", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ResumeDelete handles deleting a resume by ID.
// @Summary Delete resume
// @Description Delete a resume by ID
// @Tags resume
// @Accept json
// @Produce json
// @Param id path string true "Resume ID"
// @Success 200 {object} string "Resume deleted"
// @Failure 400 {object} string "Invalid resume ID"
// @Failure 404 {object} string "Resume not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /resume/{id} [delete]
func (h *Handler) ResumeDelete(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	_, err := h.srvs.Resume.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete resume", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Resume deleted"})
}
