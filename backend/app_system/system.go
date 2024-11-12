package appsystem

import (
	"pharmaDashServer/token"
	"pharmaDashServer/util"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type System struct {
	accounts Accounts
}

func InitializeSystem() *System {
	return &System{}
}

func (system *System) RegisterNewAccount(payload *util.RegisterRequest, p *pgxpool.Pool) (httpStatus int, jsonBody gin.H) {
	return system.accounts.RegisterNewAccount(payload, p)
}

func (system *System) Login(payload *util.LoginRequest, p *pgxpool.Pool, maker *token.PasetoMaker, c *gin.Context) (httpStatus int, jsonBody gin.H) {
	return system.accounts.Login(payload, p, maker, c)
}
