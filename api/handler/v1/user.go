package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Asliddin3/api-gateway/genproto/user"
	l "github.com/Asliddin3/api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary create user with info
// @Description this func create user
// @Tags user
// @Accept json
// @Produce json
// @Param user body user.UserRequest true "User"
// @Success 201 {object} user.UserResponse
// @Router /user [post]
func (h *handlerV1) Create(c *gin.Context) {
	var (
		body        user.UserRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.UserService().CreateUser(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// @BasePath /api/v1
// @Summary get all users
// @Description this func get all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} user.UsersResponse
// @Router /customer/list [get]
func (h *handlerV1) FindAll(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.UserService().GetAllUsers(ctx, &user.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get list users", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @BasePath /api/v1

// @Summary get user info
// @Description this func get user info
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} user.UserResponse
// @Router /user/{id} [get]
func (h *handlerV1) FindByID(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	idStr := c.Param("id")
	fmt.Println("id str----", idStr)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert string to int", l.Error(err))
		return
	}
	body := &user.UserId{
		Id: id,
	}
	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(h.cfg.CtxTimeout))

	defer cancel()

	response, err := h.serviceManager.UserService().GetUser(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user info", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}
