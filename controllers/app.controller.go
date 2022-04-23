package controllers

import (
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"log"
	"quizbe/models"
	"quizbe/services"
	"quizbe/utils"
)

// CreateApp func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]
func CreateApp(c *fiber.Ctx) error {

	resData := utils.ResponseData{}
	param := models.CreateAppParam{}

	if err := c.BodyParser(&param); err != nil {
		return err
	}

	if err := services.CreateApp(param); err != nil {
		defer sentry.CaptureException(err)
		resData.Signal = 0
		resData.ErrorCode = 0
		resData.Message = "loi"
		resData.Data = nil

		return c.Status(201).JSON(resData)
	}

	resData.Signal = 1
	resData.ErrorCode = 1
	resData.Message = "thanh cong"
	resData.Data = nil
	return c.Status(201).JSON(resData)
}

// UpdateApp func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]
func UpdateApp(c *fiber.Ctx) error {
	resData := utils.ResponseData{}
	param := models.UpdateAppParam{}

	if err := c.BodyParser(&param); err != nil {
		return err
	}

	appIdparam := param.AppId
	updateParam := map[string]interface{}{
		"name":        param.Name,
		"description": param.Description,
		"extend":      param.Extend,
	}

	app, err := services.UpdateApp(appIdparam, updateParam)

	if app.ID == 0 || err != nil {
		resData.Signal = 0
		resData.ErrorCode = 0
		resData.Message = "khong update duoc"
		resData.Data = nil
		defer sentry.CaptureException(err)
		return c.Status(201).JSON(resData)
	}

	resData.Signal = 1
	resData.ErrorCode = 1
	resData.Message = "thanh cong"
	resData.Data = app
	return c.Status(201).JSON(resData)
}

// DeleteApp func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]
func DeleteApp(c *fiber.Ctx) error {

	resData := utils.ResponseData{}
	param := models.DeleteAppParam{}

	if err := c.BodyParser(&param); err != nil {
		return err
	}

	err := services.DeleleApp(param.AppId)
	if err != nil {
		resData.Signal = 0
		resData.ErrorCode = 0
		resData.Message = "loi"
		resData.Data = nil

		return c.Status(201).JSON(resData)
	}

	resData.Signal = 1
	resData.ErrorCode = 0
	resData.Message = "xoa thanh cong"
	resData.Data = nil

	return c.Status(201).JSON(resData)
}

// GetApps func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]
func GetApps(c *fiber.Ctx) error {
	resData := utils.ResponseData{}

	apps, err := services.GetApps()

	if err != nil {
		log.Printf("%v\n", err)
	}

	if len(apps) == 0 {
		resData.Signal = 0
		resData.ErrorCode = 0
		resData.Message = "khong thay ket qua"
		resData.Data = nil
		return c.Status(201).JSON(resData)
	}

	resData.Signal = 1
	resData.ErrorCode = 0
	resData.Message = "thanh cong"
	resData.Data = apps
	return c.Status(201).JSON(resData)

}

// GetApp func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]
func GetApp(c *fiber.Ctx) error {

	resData := utils.ResponseData{}
	param := models.GetAppParam{}

	if err := c.BodyParser(&param); err != nil {
		return err
	}

	app, err := services.GetApp(param.AppId)

	if err != nil {
		resData.Signal = 0
		resData.ErrorCode = 0
		resData.Message = "khong thay ket qua"
		resData.Data = nil
		return c.Status(201).JSON(resData)
	}

	resData.Signal = 1
	resData.ErrorCode = 0
	resData.Message = "thanh cong"
	resData.Data = app
	return c.Status(201).JSON(resData)
}
