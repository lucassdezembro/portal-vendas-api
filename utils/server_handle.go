package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber"
)

func HandleSuccessData(c *fiber.Ctx, data any, status int) {
	if data == nil {
		c.Status(fiber.StatusNoContent)
		return
	} else {

		b, err := json.Marshal(data)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		mappedData := map[string]interface{}{}
		err = json.Unmarshal(b, &mappedData)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		c.Status(status).JSON(fiber.Map{
			"data": mappedData,
		})
	}
}

func HandleErrorData(c *fiber.Ctx, err error, status int) {
	c.Status(status).JSON(fiber.Map{
		"status": "error",
		"error":  err.Error(),
	})
}
