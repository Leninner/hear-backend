package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/classroom/domain"
)

type Handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateClassroom(c *fiber.Ctx) error {
	var input struct {
		Name        string  `json:"name"`
		Building    string  `json:"building"`
		Floor       int     `json:"floor"`
		Capacity    int     `json:"capacity"`
		LocationLat float64 `json:"locationLat"`
		LocationLng float64 `json:"locationLng"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	classroom, err := h.useCase.CreateClassroom(input.Name, input.Building, input.Floor, input.Capacity, input.LocationLat, input.LocationLng)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create classroom",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(classroom)
}

func (h *Handler) GetClassroom(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid classroom ID",
		})
	}

	classroom, err := h.useCase.GetClassroom(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Classroom not found",
		})
	}

	return c.JSON(classroom)
}

func (h *Handler) GetAllClassrooms(c *fiber.Ctx) error {
	classrooms, err := h.useCase.GetAllClassrooms()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch classrooms",
		})
	}

	return c.JSON(classrooms)
}

func (h *Handler) GetClassroomsByLocation(c *fiber.Ctx) error {
	var input struct {
		Lat    float64 `query:"lat"`
		Lng    float64 `query:"lng"`
		Radius float64 `query:"radius"`
	}

	if err := c.QueryParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid location parameters",
		})
	}

	classrooms, err := h.useCase.GetClassroomsByLocation(input.Lat, input.Lng, input.Radius)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch classrooms by location",
		})
	}

	return c.JSON(classrooms)
}

func (h *Handler) UpdateClassroom(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid classroom ID",
		})
	}

	var classroom domain.Classroom
	if err := c.BodyParser(&classroom); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	classroom.ID = id
	if err := h.useCase.UpdateClassroom(&classroom); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update classroom",
		})
	}

	return c.JSON(classroom)
}

func (h *Handler) DeleteClassroom(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid classroom ID",
		})
	}

	if err := h.useCase.DeleteClassroom(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete classroom",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
} 