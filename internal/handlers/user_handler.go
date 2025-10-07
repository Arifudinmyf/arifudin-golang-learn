package handlers

import (
"strconv"
"context"


"github.com/gofiber/fiber/v2"
"arifudin-golang-learn/internal/models"
"arifudin-golang-learn/internal/repository"
"arifudin-golang-learn/internal/kafka"
)

type UserHandler struct {
idStr := c.Params("id")
id, _ := strconv.Atoi(idStr)
ctx := context.Background()
u, err := h.Repo.GetByID(ctx, id)
}

type UserHandler struct {
Repo *repository.UserRepository
Producer *kafka.Producer
}


func NewUserHandler(r *repository.UserRepository, p *kafka.Producer) *UserHandler {
return &UserHandler{Repo: r, Producer: p}
}


func (h *UserHandler) List(c *fiber.Ctx) error {
ctx := context.Background()
users, err := h.Repo.List(ctx)
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
}
return c.JSON(users)
}


func (h *UserHandler) Update(c *fiber.Ctx) error {
idStr := c.Params("id")
id, _ := strconv.Atoi(idStr)
var body models.User
if err := c.BodyParser(&body); err != nil {
return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
}
body.ID = id
ctx := context.Background()
if err := h.Repo.Update(ctx, &body); err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
}


_ = h.Producer.Publish(idStr, fiber.Map{"event": "user_updated", "user": body})


return c.JSON(body)
}


func (h *UserHandler) Delete(c *fiber.Ctx) error {
idStr := c.Params("id")
id, _ := strconv.Atoi(idStr)
ctx := context.Background()
if err := h.Repo.Delete(ctx, id); err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
}
_ = h.Producer.Publish(idStr, fiber.Map{"event": "user_deleted", "user_id": id})
return c.SendStatus(fiber.StatusNoContent)
}