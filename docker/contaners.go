package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/gofiber/fiber/v2"
)

// 1️⃣ Barcha konteynerlarni ko‘rish
func ListContainers(c *fiber.Ctx) error {
	ctx := context.Background()
	containers, err := DockerClient.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var result []fiber.Map
	for _, cnt := range containers {
		result = append(result, fiber.Map{
			"ID":    cnt.ID[:10],
			"Name":  cnt.Names[0],
			"State": cnt.State,
			"Image": cnt.Image,
		})
	}

	return c.JSON(result)
}

// 2️⃣ Yangi konteyner yaratish
func CreateContainer(c *fiber.Ctx) error {
	type Request struct {
		Image string `json:"image"`
		Name  string `json:"name"`
		Port  string `json:"port"`
	}
	req := new(Request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx := context.Background()
	resp, err := DockerClient.ContainerCreate(ctx, &container.Config{
		Image: req.Image,
		ExposedPorts: nat.PortSet{
			nat.Port(req.Port): struct{}{},
		},
	}, nil, nil, nil, req.Name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Container created", "ID": resp.ID})
}

// 3️⃣ Konteynerni ishga tushirish
func StartContainer(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	if err := DockerClient.ContainerStart(ctx, id, types.ContainerStartOptions{}); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Container started", "ID": id})
}

// 4️⃣ Konteynerni to‘xtatish
func StopContainer(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	if err := DockerClient.ContainerStop(ctx, id, nil); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Container stopped", "ID": id})
}

// 5️⃣ Konteynerni o‘chirish
func RemoveContainer(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	if err := DockerClient.ContainerRemove(ctx, id, types.ContainerRemoveOptions{Force: true}); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Container removed", "ID": id})
}
