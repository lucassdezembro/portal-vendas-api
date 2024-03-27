package utils

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber"
	errors_utils "github.com/lucassdezembro/portal-vendas-api/utils/errors"
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

		var mappedData any
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

func HandleErrorData(c *fiber.Ctx, err errors_utils.Error, status int) {

	customSplittedError := strings.Split(err.Error(), ":")
	customMessage := err.Error()

	if len(customSplittedError) > 0 {

		regex := regexp.MustCompile(`^[0-9]+$`)

		if regex.MatchString(customSplittedError[0]) {
			parsedStatus, parseError := strconv.Atoi(customSplittedError[0])
			if parseError == nil {
				status = parsedStatus
				customMessage = strings.Replace(err.Error(), customSplittedError[0]+": ", "", 1)
			}
		}
	}

	if status == 0 {
		status = fiber.StatusInternalServerError
	}

	c.Status(status).JSON(fiber.Map{
		"status": "error",
		"error":  customMessage,
	})
}
