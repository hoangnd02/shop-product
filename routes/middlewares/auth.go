package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/params"
	"github.com/hoanggggg5/shopproduct/usecase"
	"github.com/hoanggggg5/shopproduct/utils"
)

func CheckRequest(userService usecase.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jwt_auth, err := utils.ParserJWT(strings.Replace(c.Get("Authorization"), "Bearer ", "", -1))
		if err != nil {
			return params.FailedToParseJWT
		}

		user := &models.User{
			UID:      jwt_auth.UID,
			Username: jwt_auth.Username,
			Email:    jwt_auth.Email,
			Role:     jwt_auth.Role,
			State:    jwt_auth.State,
		}

		if _, err := userService.Get(jwt_auth.UID); err != nil {
			if _, err := userService.Create(user); err != nil {
				return err
			}
		} else {
			if _, err := userService.Update(user); err != nil {
				return err
			}
		}

		c.Locals("current_user", user)

		return c.Next()
	}
}
