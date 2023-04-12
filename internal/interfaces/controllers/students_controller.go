package controllers

import (
	"CBI/internal/domain/entity"
	"CBI/internal/domain/usecase"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentsController struct {
	StudentsUseCase usecase.StudentsUseCase
}

func (ctrl *StudentsController) GetAll(c *gin.Context) {
	students, err := ctrl.StudentsUseCase.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve students",
		})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (uc *StudentsController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid student ID"})
		return
	}

	student, err := uc.StudentsUseCase.FindById(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}
func (uc *StudentsController) Create(c *gin.Context) {
	var student entity.Students
	if err := c.ShouldBindJSON(&student); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if err := uc.StudentsUseCase.Register(c.Request.Context(), student); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create students successfully"})

}

func (uc *StudentsController) UpdateById(c *gin.Context) {
	{
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid student ID"})
			return
		}

		var student entity.Students
		if err := c.ShouldBindJSON(&student); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
			return
		}

		if err := uc.StudentsUseCase.EditById(c.Request.Context(), id, student); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "student not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "student update successfully"})
	}
}
func (uc *StudentsController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := uc.StudentsUseCase.Unregister(c.Request.Context(), id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "student deleted successfully"})

}

func NewStudentsController(useCase usecase.StudentsUseCase) *StudentsController {
	return &StudentsController{StudentsUseCase: useCase}
}
