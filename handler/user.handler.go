package handler

import (
	"fiber/backend/database"
	"fiber/backend/model/entity"
	"fiber/backend/model/request"
	"fiber/backend/model/response"
	"fiber/backend/utils"
	"fiber/backend/validation"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get user
// @Description Get user
// @ID get-user
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit per page" default(10)
// @Param query query string false "Search query"
// @Success 200 {object} entity.User
// @Security Bearer
// @Router /user [get]
func GetAllUser(c *fiber.Ctx) error {
	var users []entity.User
	var totalRecord int64
	pageParam := c.Query("page","1")
	limitParam := c.Query("limit","10")
	query := c.Query("query")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	result := database.DB.Debug().Model(&users)
	if query != "" {
		result = result.Where("name = ?", query).Or("email = ?",query)
	}

	// for pagination
	if err := result.Count(&totalRecord).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return utils.JSONResponse(c,fiber.StatusInternalServerError,"error","unable to get users", err.Error())
	}

	return utils.JSONResponseWithPagination(c,fiber.StatusOK,"OK","Show all user", users,page,limit,totalRecord)
}

// @Summary Create User
// @Description Create a new user with name, email, phone, and address
// @ID create-user
// @Accept json
// @Produce json
// @Param data body request.UserCreateRequest true "User data"
// @Success 200 {object} entity.User
// @Failure 400 {object} utils.ErrorResponseSwagger
// @Security Bearer
// @Router /user [post]
func UserCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
	}

	err := validation.ValidateStruct(user)
	log.Println(err)
	if len(err) > 0 {
		return utils.JSONResponse(c,fiber.StatusBadRequest, "request error", err,nil)
	}
	var findEmail entity.User
	if err := database.DB.Debug().Where("email = ?", user.Email).First(&findEmail).Error; err == nil {
		var email = user.Email
		return utils.JSONResponse(c,fiber.StatusBadRequest,"request error","user " + email + " already exist", email)
	}

	create := entity.User {
		Name: user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Address: user.Address,
		CreatedAt: time.Now(),
	}

	if err := database.DB.Debug().Create(&create).Error; err != nil {
		return utils.JSONResponse(c,fiber.StatusInternalServerError,"internal server error",err.Error(), nil)
	}

	return utils.JSONResponse(c,fiber.StatusCreated,"OK","Created", create)
}

// @Summary Get user By ID
// @Description Get user By ID
// @ID get-user-by-id
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.User
// @Failure 404 {object} utils.ErrorResponseSwagger
// @Failure 400 {object} utils.ErrorResponseSwagger
// @Security Bearer
// @Router /user/{id} [get]
func GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	userId,err := strconv.Atoi(id)
	
	if err != nil {
		return utils.JSONResponse(c,fiber.StatusBadRequest,"Bad Request","user id must be an integer", id)
	}

	var user response.User
	if err := database.DB.Debug().First(&user, "id = ?", userId).Error; err != nil {
		return utils.JSONResponse(c,fiber.StatusNotFound,"OK","not found", nil)
	}

	return utils.JSONResponse(c,fiber.StatusOK,"OK","show by id", user)
}

// @Summary Update User
// @Description Update user with name, phone, and address
// @ID update-user
// @Accept json
// @Produce json
// @param id path int true "User ID"
// @Param data body request.UserUpdateRequest true "User data"
// @Success 200 {object} entity.User
// @Failure 400 {object} utils.ErrorResponseSwagger
// @Failure 404 {object} utils.ErrorResponseSwagger
// @Security Bearer
// @Router /user/{id} [patch]
func UserUpdate(c *fiber.Ctx) error {
	data := new(request.UserUpdateRequest)
	id := c.Params("id")
	userId,errParams := strconv.Atoi(id)
	var user entity.User

	if err := c.BodyParser(data); err != nil {
		log.Println(err)
	}
	
	errValidation := validation.ValidateStruct(data)
	if len(errValidation) > 0 {
		return utils.JSONResponse(c,fiber.StatusBadRequest, "request error", errValidation,nil)
	}
		
	if errParams != nil {
		return utils.JSONResponse(c,fiber.StatusBadRequest,"Bad Request","user id must be an integer", id)
	}

	if err := database.DB.Debug().First(&user, "id = ?", userId).Error; err != nil {
		return utils.JSONResponse(c,fiber.StatusNotFound,"OK","not found", nil)
	}

	if err := database.DB.Debug().Model(&user).Updates(map[string]interface{}{
		"Name":    data.Name,
    	"Address": data.Address,
    	"Phone":   data.Phone,
		}).Error; err != nil {
		return utils.JSONResponse(c,fiber.StatusNotFound,"error","internal server error", err.Error())
	}

	return utils.JSONResponse(c,fiber.StatusOK,"OK","updated", user)
}


// @Summary Delete User
// @Description delete user
// @ID delete-user
// @Produce json
// @param id path int true "User ID"
// @Success 204 {object} entity.User
// @Failure 400 {object} utils.ErrorResponseSwagger
// @Failure 404 {object} utils.ErrorResponseSwagger
// @Security Bearer
// @Router /user/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	userId,err := strconv.Atoi(id)
	
	if err != nil {
		return utils.JSONResponse(c,fiber.StatusBadRequest,"Bad Request","user id must be an integer", id)
	}

	var user response.User
	if err := database.DB.Debug().First(&user, "id = ?", userId).Error; err != nil {
		return utils.JSONResponse(c,fiber.StatusNotFound,"OK","not found", nil)
	}

	if err := database.DB.Debug().Delete(&user).Error; err != nil {
		return utils.JSONResponse(c,fiber.StatusInternalServerError,"internal server error",err.Error(), nil)
	}

	return utils.JSONResponse(c,fiber.StatusNoContent,"OK","Deleted",nil)
}