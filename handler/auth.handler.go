package handler

import (
	"fiber/backend/database"
	"fiber/backend/model/entity"
	"fiber/backend/model/request"
	"fiber/backend/utils"
	"fiber/backend/validation"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// @Summary Login user
// @Description login with email,password
// @ID login-user
// @Accept json
// @Produce json
// @Param data body request.LoginRequest true "Login data (default value: Email=test@email.com,Password=12345)"
// @Success 200 {object} entity.Auth
// @Failure 400 {object} utils.ErrorResponseSwagger
// @Failure 404 {object} utils.ErrorResponseSwagger
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	req := new(request.LoginRequest)

	if err := c.BodyParser(req); err != nil {
		log.Println(err)
	}

	err := validation.ValidateStruct(req)
	if len(err) > 0 {
		return utils.JSONResponse(c,fiber.StatusBadRequest, "request error", err,nil)
	}
	var user entity.Auth
	if err := database.DB.Debug().First(&user, "email = ?", req.Email).Error; err != nil {
		return utils.JSONResponse(c,fiber.StatusNotFound,"OK","not found", nil)
	}

	checkPassword := utils.ComparePassword(req.Password,user.Password)
	if !checkPassword {
		return utils.JSONResponse(c,fiber.StatusUnauthorized,"OK","invalid email & password", nil)
	}

	token,errToken := utils.GenerateAccessToken(user.Email)
	if errToken != nil {
		return utils.JSONResponse(c,fiber.StatusUnauthorized,"OK","invalid access token", errToken.Error())
	}

	refreshToken,errRefreshToken := utils.GenerateRefreshToken(user.Email)
	if errRefreshToken != nil {
		return utils.JSONResponse(c,fiber.StatusUnauthorized,"OK","invalid access token", errRefreshToken.Error())
	}

	user.Password = ""

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"OK",
		"message":"Login success",
		"data":user,
		"token":token,
		"refresh_token":refreshToken,
	})
}

func RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.FormValue("refresh_token")
	token, err := utils.VerifyRefreshToken(refreshToken)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid refresh token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token claims",
		})
	}

	username := claims["username"].(string)
	newAccessToken, err := utils.GenerateAccessToken(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error generating new access token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": newAccessToken,
	})
}

// @Summary Create login user
// @Description Create login user with email,password
// @ID create-login-user
// @Accept json
// @Produce json
// @Param data body request.AuthRequest true "Create login data"
// @Success 200 {object} entity.Auth
// @Failure 400 {object} utils.ErrorResponseSwagger
// @Failure 404 {object} utils.ErrorResponseSwagger
// @Router /auth-create [post]
func AuthCreate(c *fiber.Ctx) error {
	user := new(request.AuthRequest)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
	}

	err := validation.ValidateStruct(user)
	log.Println(err)
	if len(err) > 0 {
		return utils.JSONResponse(c,fiber.StatusBadRequest, "request error", err,nil)
	}
	var findEmail entity.Auth
	if err := database.DB.Debug().Where("email = ?", user.Email).First(&findEmail).Error; err == nil {
		var email = user.Email
		return utils.JSONResponse(c,fiber.StatusBadRequest,"request error","user " + email + " already exist", email)
	}

	hashPassword,errHash := utils.HashPassword(user.Password)

	if errHash != nil {
		return utils.JSONResponse(c,fiber.StatusBadRequest, "request error", err,nil)
	}

	create := entity.Auth {
		Email: user.Email,
		Password: hashPassword,
		CreatedAt: time.Now(),
	}

	if err := database.DB.Debug().Create(&create).Error; err != nil {
		return utils.JSONResponse(c,fiber.StatusInternalServerError,"internal server error",err.Error(), nil)
	}

	create.Password = ""

	return utils.JSONResponse(c,fiber.StatusCreated,"OK","Created", create)
}