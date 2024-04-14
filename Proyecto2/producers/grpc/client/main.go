package main

import (
	"context"
	"fmt"
	"log"

	pb "proyecto2/clientgRPC/proto"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ctx = context.Background()

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func insertData(c *fiber.Ctx) error {

	var body Data
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// connection to grpc
	conn, err := grpc.NewClient("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Name:  body.Name,
		Album: body.Album,
		Year:  body.Year,
		Rank:  body.Rank,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Respuesta del server " + ret.GetInfo())

	return c.Status(200).JSON(fiber.Map{
		"message": "Data inserted",
	})
}

func main() {
	app := fiber.New()
	app.Post("/grcp", insertData)

	err := app.Listen(":3000")
	if err != nil {
		log.Println(err)
		return
	}
}
