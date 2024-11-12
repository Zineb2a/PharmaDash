package server

import (
	"encoding/json"
	"io"
	"net/http"
	"pharmaDashServer/util"

	"github.com/gin-gonic/gin"
)

func (server *Server) RegisterNewAccount(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	var payload util.RegisterRequest
	err = json.Unmarshal(jsonData, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid json payload format."})
		return
	}
	status, jsonBody := server.system.RegisterNewAccount(&payload, server.pool)
	c.JSON(status, jsonBody)
}

// route: /api/user/login
func (server *Server) Login(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	var payload util.LoginRequest
	err = json.Unmarshal(jsonData, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid json payload format."})
		return
	}
	server.system.Login(&payload, server.pool, server.maker, c)
}

func (server *Server) mustAuthChecker(c *gin.Context) {
	tok, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "User not Authenticated"})
		return
	}
	//validation here
	payload, err := server.maker.VerifyToken(tok) //must cast to payload.(*token.Payload) in order to use in subsequent routes
	if err != nil {
		c.SetCookie("token", "", -1, "", "", false, true)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Authentication failed."})
		return
	}
	c.Set("auth_payload", payload)
}

// route: /api/user/logout
func (server *Server) LogOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "Successfully logged out."})
}
