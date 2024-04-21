package controllers

import (
	"errors"
	"time"

	"blogy/config"
	"blogy/database"
	"blogy/models"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func getUserByUsername(u string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Username: u}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func getUserByEmail(e string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Email: e}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func login(c *fiber.ctx) error {
	type InputData struct {
		Indentity string
		Password  string
	}

	input = new(InputData)
	user, err := new(models.User), *new(error)

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on login request",
			"data":    err,
		})
	}

	if utils.isEmail(input.Indentity) {
		user, err = getUserByUsername(input.Indentity)
	} else {
		user, err = getUserByUsername(input.Indentity)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error",
			"data":    err,
		})
	} else if user == nil {
		CheckPasswordHash(pass, "")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid identity or password",
			"data":    err,
		})
	} else {
		if !CheckPasswordHash(input.Password, user.Password) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid identity or password",
				"data":    nil,
			})
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	jwtToken, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success login",
		"data":    jwtToken,
	})
}

func register(c *fiber.ctx) error {
	type InputData struct {
		Username string
		Name     string
		Email    string
		Password string
	}

	input = new(InputData)
	user, err := new(models.User), *new(error)

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on register request",
			"data":    err,
		})
	}
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error",
			"data": err,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 0)
	
	user.CreateUser({
		"username": input.Username,
		"password": hashedPassword,
	})
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	jwtToken, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success login",
		"data": jwtToken,
	})
}
