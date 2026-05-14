package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"stellarbill-backend/internal/service"
)

// NewGetStatementHandler returns a gin.HandlerFunc that retrieves a statement.
func NewGetStatementHandler(svc service.StatementService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
	}
}

// NewListStatementsHandler returns a gin.HandlerFunc that lists statements.
func NewListStatementsHandler(svc service.StatementService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"statements": []interface{}{}})
	}
}
