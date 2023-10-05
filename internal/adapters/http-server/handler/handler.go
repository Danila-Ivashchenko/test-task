package handler

import (
	"context"
	"fmt"
	"go-kafka/internal/domain/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type config interface {
	GetTimeLimit() time.Duration
}

type userService interface {
	AddUser(ctx context.Context, user *model.UserInsertRequest) error
	GetUsers(ctx context.Context, dto *model.GetUsersDTO) (*model.Users, error)
	GetUserById(ctx context.Context, dto *model.IdDTO) (*model.User, error)
	DeleteUser(ctx context.Context, dto *model.IdDTO) error
	UpdateUser(ctx context.Context, dto *model.User) error
}

type handler struct {
	service   userService
	timeLimit time.Duration
}

func New(cfg config, s userService) *handler {
	return &handler{
		service:   s,
		timeLimit: cfg.GetTimeLimit(),
	}
}

func (h handler) AddUser(c *gin.Context) {
	request := &model.UserInsertRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	err = h.service.AddUser(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, successResponse{Success: true})
}

func (h handler) GetUsers(c *gin.Context) {
	request := &model.GetUsersDTO{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	request.Name = c.Query("name")
	request.Surname = c.Query("suname")
	request.Patronymic = c.Query("patronymic")
	ageStr := c.Query("age")
	if ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, badResponse{Message: fmt.Sprintf("invalid value of age, got: %s, must be: int", ageStr)})
			return
		}
		request.Age = uint(age)
	}

	request.Gender = c.Query("gender")
	request.Nationality = c.Query("nationality")

	offsetStr := c.Query("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, badResponse{Message: fmt.Sprintf("invalid value of offset, got: %s, must be: int", offsetStr)})
			return
		}
		request.Offset = offset
	}
	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, badResponse{Message: fmt.Sprintf("invalid value of limit, got: %s, must be: int", limitStr)})
			return
		}
		request.Limit = limit
	}
	fmt.Println(request)
	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	result, err := h.service.GetUsers(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h handler) GetUserById(c *gin.Context) {
	request := &model.IdDTO{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	idStr := c.Query("id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, badResponse{Message: fmt.Sprintf("invalid value of id, got: %s, must be: int", idStr)})
			return
		}
		request.Id = int64(id)
	} else {
		c.JSON(http.StatusBadRequest, badResponse{Message: "no id in request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	result, err := h.service.GetUserById(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h handler) UpdateUser(c *gin.Context) {
	request := &model.User{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	err = h.service.UpdateUser(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, successResponse{Success: true})
}

func (h handler) DeleteUser(c *gin.Context) {
	request := &model.IdDTO{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	err = h.service.DeleteUser(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, successResponse{Success: true})
}
