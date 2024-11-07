package server

import (
	"context"
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
	//loginMode := c.PostForm("loginMode")
	createParams := db.CreateUserParams{
		Name:        c.PostForm("name"),
		LastName:    c.PostForm("lastname"),
		Password:    c.PostForm("password"),
		PhoneNumber: pgtype.Text{String: c.PostForm("phonenumber"), Valid: true},
		Email:       pgtype.Text{String: c.PostForm("email"), Valid: true},
		Address:     pgtype.Text{String: c.PostForm("address"), Valid: true},
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
	if createParams.Name == "" || createParams.LastName == "" || createParams.Password == "" || createParams.Email.String == "" {
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
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Could not create account, try later."})
		return
	}
	createParams.Password = hash

	//for later
	// if loginMode == "client" {

	// } else if loginMode == "driver" {

	// } else {

	// }

	_, err = query.CreateUser(ctx, createParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Could not create account, try later."})
		return
	}

	//send success
	c.JSON(http.StatusOK, gin.H{"status": "account created"})
}

// route: /api/user/login
func (server *Server) Login(c *gin.Context) {
	ctx := context.Background()

	email := c.PostForm("emailOrUsername")
	password := c.PostForm("password")

	//acquire connection from connection pools
	conn, err := server.pool.Acquire(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	defer conn.Release()
	query := db.New(conn)

	//make sure all the necessary info is there
	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid data"})
		return
	}

	//check if user exists
	var dbUser db.Account
	dbEmailUser, err := query.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "User does not exist"})
		return
	}
	dbUser = dbEmailUser

	if !util.CheckPasswordHash(password, dbUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Authentication failed."})
		return
	} else {
		token, err := server.maker.CreateToken(email, time.Hour)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Authentication failed due to internal error."})
			return
		}
		//modify for local storage
		c.SetCookie("lingtoken", token, 3600, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{"status": "Authentication successful"})
		return
	}
}

// route: /api/user/logout
func (server *Server) LogOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "Successfully logged out."})
}
