package user_repository

import (
	"github.com/gin-gonic/gin"
	"github.com/pg06pg/stay-sorted/models/entities"
)

type UserRepository interface {
	CreateUser(ctx *gin.Context, user entities.User) (entities.User, error)
	GetUserByPid(ctx *gin.Context, userPid string) (entities.User, error)
	GetUserByEmail(ctx *gin.Context, email string) (entities.User, error)
}
