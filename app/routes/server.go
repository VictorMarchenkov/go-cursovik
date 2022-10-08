package routes

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-cursovic/database"
	"go-cursovic/models"
	"go-cursovic/utils"
	"strconv"
)

func redirect(c *fiber.Ctx) error {
	shUrl := c.Params("redirect")
	shurl, err := database.QueryFindByShurlUrl(shUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find short url in DB " + err.Error(),
		})
	}
	// grab any stats you want...
	shurl.Clicked += 1
	err = database.QueryUpdateShurls(shurl)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return c.Redirect(shurl.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllShurls(c *fiber.Ctx) error {
	su, err := database.QueryGetAllShurles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all short urls links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(su)
}

func getShurl(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}

	su, err := database.QueryGetShurl(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retreive short urls from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(su)
}

func createShurl(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var su models.Shurls
	fmt.Println(su)
	err := c.BodyParser(&su)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if su.Random {
		su.Shurl = utils.RandomURL(8)
	}

	err = database.QueryCreateShurl(su)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create short url in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(su)

}

func updateShurl(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var shurl models.Shurls

	err := c.BodyParser(&shurl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse json " + err.Error(),
		})
	}

	err = database.QueryUpdateShurls(shurl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update short url link in DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(shurl)
}

func deleteShurl(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse id from url " + err.Error(),
		})
	}

	err = database.QueryDeleteShurl(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not delete from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "short url deleted.",
	})
}

func SetupAndListen() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PATCH, HEAD, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", redirect)

	router.Get("/shurl", getAllShurls)
	router.Get("/shurl/:id", getShurl)
	router.Post("/shurl", createShurl)
	router.Patch("/shurl", updateShurl)
	router.Delete("/shurl/:id", deleteShurl)

	router.Listen(":3000")

}
