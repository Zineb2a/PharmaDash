package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	db "pharmaDashServer/db/sqlc"
	"pharmaDashServer/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

// route: /api/user/register
func (server *Server) RegisterNewAccount(c *gin.Context) {
	ctx := context.Background()

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

	createParams := db.CreateUserParams{
		Name:        payload.UserData.Name,
		LastName:    payload.UserData.LastName,
		Password:    payload.UserData.Password,
		PhoneNumber: payload.UserData.PhoneNumber,
		Email:       payload.UserData.Email,
		Address:     payload.UserData.Address,
	}

	//acquire connection from connection pool
	conn, err := server.pool.Acquire(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	defer conn.Release()
	query := db.New(conn)

	//make sure all the necessary info is there
	if createParams.Name == "" || createParams.LastName == "" || createParams.Password == "" || createParams.Email.String == "" || createParams.Address.String == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid data"})
		return
	}

	//check if user already exists
	if strings.Contains(createParams.Email.String, "@") {
		_, err := query.GetUserByEmail(ctx, createParams.Email)
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Email is already in use."})
			return
		}
	}

	//if not hash password
	hash, err := util.HashPassword(createParams.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Could not create account, try later. Hash"})
		return
	}
	createParams.Password = hash

	_, err = query.CreateUser(ctx, createParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Could not create account, try later. DB connection"})
		return
	}

	//send success
	c.JSON(http.StatusOK, gin.H{"status": "account created"})
}

// route: /api/user/login
func (server *Server) Login(c *gin.Context) {
	ctx := context.Background()

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

	//acquire connection from connection pools
	conn, err := server.pool.Acquire(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	defer conn.Release()
	query := db.New(conn)

	//make sure all the necessary info is there
	if payload.Email == "" || payload.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid data"})
		return
	}

	//check if user exists
	var dbUser db.Account
	dbEmailUser, err := query.GetUserByEmail(ctx, pgtype.Text{String: payload.Email, Valid: true})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "User does not exist"})
		return
	}
	dbUser = dbEmailUser

	if !util.CheckPasswordHash(payload.Password, dbUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Authentication failed."})
		return
	} else {
		token, err := server.maker.CreateToken(payload.Email, dbUser.Authlevel, time.Hour)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Authentication failed due to internal error."})
			return
		}
		//modify for local storage
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{"status": "Authentication successful"})
		return
	}
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
