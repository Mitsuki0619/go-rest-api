package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func getLoginUser(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	rawUserId := claims["user_id"]
	userId := uint(rawUserId.(float64))
	return userId
}

func judgeResponseOkOrInternalServerError(c echo.Context, err error, res interface{}) error {
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if res == nil {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, res)
}

func getTaskIdFromParam(c echo.Context) uint {
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	return uint(taskId)
}

func getTaskFromRequestBody(c echo.Context) (model.Task, error) {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	userId := getLoginUser(c)
	tasksRes, err := tc.tu.GetAllTasks(userId)
	return judgeResponseOkOrInternalServerError(c, err, tasksRes)
}

func (tc *taskController) GetTaskById(c echo.Context) error {
	userId := getLoginUser(c)
	taskId := getTaskIdFromParam(c)
	taskRes, err := tc.tu.GetTaskById(userId, taskId)
	return judgeResponseOkOrInternalServerError(c, err, taskRes)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	userId := getLoginUser(c)
	task, err := getTaskFromRequestBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserId = userId
	taskRes, err := tc.tu.CreateTask(task)
	return judgeResponseOkOrInternalServerError(c, err, taskRes)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	userId := getLoginUser(c)
	taskId := getTaskIdFromParam(c)
	task, err := getTaskFromRequestBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserId = userId
	taskRes, err := tc.tu.UpdateTask(task, userId, taskId)
	return judgeResponseOkOrInternalServerError(c, err, taskRes)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	userId := getLoginUser(c)
	taskId := getTaskIdFromParam(c)
	err := tc.tu.DeleteTask(userId, taskId)
	return judgeResponseOkOrInternalServerError(c, err, nil)
}
