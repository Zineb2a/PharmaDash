package appsystem

import (
	"context"
	"net/http"
	db "pharmaDashServer/db/sqlc"
	"pharmaDashServer/token"
	"pharmaDashServer/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Accounts struct {
}

func (account *Accounts) RegisterNewAccount(payload *util.RegisterRequest, p *pgxpool.Pool) (httpStatus int, jsonBody gin.H) {
	ctx := context.Background()
	createParams := db.CreateUserParams{
		Name:        payload.UserData.Name,
		LastName:    payload.UserData.LastName,
		Password:    payload.UserData.Password,
		PhoneNumber: payload.UserData.PhoneNumber,
		Email:       payload.UserData.Email,
		Address:     payload.UserData.Address,
	}

	//acquire connection from connection pool
	conn, err := p.Acquire(ctx)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"status": "Server error."}
	}
	defer conn.Release()
	query := db.New(conn)

	//make sure all the necessary info is there
	if createParams.Name == "" || createParams.LastName == "" || createParams.Password == "" || createParams.Email.String == "" {
		return http.StatusBadRequest, gin.H{"status": "Invalid data"}
	}

	//check if user already exists
	if strings.Contains(createParams.Email.String, "@") {
		_, err := query.GetUserByEmail(ctx, createParams.Email)
		if err == nil {
			return http.StatusBadRequest, gin.H{"status": "Email is already in use."}
		}
	}

	//if not hash password
	hash, err := util.HashPassword(createParams.Password)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"status": "Could not create account, try later."}
	}
	createParams.Password = hash

	_, err = query.CreateUser(ctx, createParams)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"status": "Could not create account, try later."}
	}

	//send success
	return http.StatusOK, gin.H{"status": "account created"}
}

func (account *Accounts) Login(payload *util.LoginRequest, p *pgxpool.Pool, maker *token.PasetoMaker, c *gin.Context) (httpStatus int, jsonBody gin.H) {
	ctx := context.Background()
	conn, err := p.Acquire(ctx)
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
		token, err := maker.CreateToken(payload.Email, time.Hour)
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
