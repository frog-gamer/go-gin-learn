package user

import (
	"myapp/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ErrorResponse(c, "400", "Invalid input")
		return
	}

	if err := h.useCase.Register(&user); err != nil {
		response.ErrorResponse(c, "500", "Registration failed")
		return
	}

	response.SuccessResponse(c, gin.H{"email": user.Email})
}

func (h *Handler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, "400", "Invalid input")
		return
	}

	token, err := h.useCase.Login(input.Email, input.Password)
	if err != nil {
		response.ErrorResponse(c, "401", "Invalid credentials")
		return
	}

	response.SuccessResponse(c, gin.H{"token": token})
}

func (h *Handler) GetUsers(c *gin.Context) {
	// Parse page and limit from query parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Fetch paginated users
	paginatedData, err := h.useCase.GetUsers(page, limit)
	if err != nil {
		response.ErrorResponse(c, "500", "Failed to fetch users")
		return
	}

	// Return the paginated data
	response.SuccessResponse(c, paginatedData)
}
