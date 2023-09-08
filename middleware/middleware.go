package middleware

import (
	"adv/models"

	"github.com/gin-gonic/gin"
)

func ValidateIdParam(c *gin.Context) {
	var currentId models.IdPath
	if err := c.ShouldBindUri(&currentId); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Ошибка валидации url-пути": "id должен быть целым числом"})
		return
	}
	return
}

func AuthRequiredCheck(c *gin.Context) {

}
