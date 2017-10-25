package Handlers

import (
	input "github.com/NirmalVatsyayan/FuzzyWuzzyGo/Input"
	helpers "github.com/NirmalVatsyayan/FuzzyWuzzyGo/Helpers"
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
)



func AbortWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
	c.Abort()
}

func GetFuzzyScore(c *gin.Context){

	var input input.InputData

	if c.BindJSON(&input) != nil {
		AbortWithError(c, http.StatusBadRequest, "Missing data")
		return
	}

	token_set_ratio, token_sort_ratio, partial_ratio := helpers.CalculateTokenRatio(input.Data1, input.Data2)

	c.JSON(200, gin.H{"token_set_ratio":token_set_ratio, "token_sort_ratio":token_sort_ratio, "partial_ratio":partial_ratio})

}
