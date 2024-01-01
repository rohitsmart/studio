// controller/authcontroller.go

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rohitsmart/studio/model"
	"github.com/rohitsmart/studio/service"
	"github.com/rohitsmart/studio/util"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// func (c *AuthController) SignUp(ctx *gin.Context) {
// 	var user model.User

// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		util.NewErrorResponse(ctx, "Invalid request", "Invalid request format", "4001", 400)
// 		return
// 	}

// 	if err := c.authService.SignUp(&user); err != nil {
// 		util.NewErrorResponse(ctx, "Internal Server Error", "Failed to register user", "5001", 500)
// 		return
// 	}

// 	ctx.JSON(200, gin.H{"message": "User registered successfully"})
// }

func (c *AuthController) SignUp(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		util.NewErrorResponse(ctx, "Invalid request", "Invalid request format", "4001", 400)
		return
	}

	if err := c.authService.SignUp(&user); err != nil {
		switch err {
		case service.ErrDuplicateUsername:
			util.NewErrorResponse(ctx, "Duplicate username", "Username is already taken", "4002", 400)
		case util.ErrInvalidCredentials:
			util.NewErrorResponse(ctx, "Invalid credentials", "Invalid password", "4003", 400)
		default:
			util.NewErrorResponse(ctx, "Internal Server Error", "Failed to register user", "5001", 500)
		}
		return
	}

	ctx.JSON(200, gin.H{"message": "User registered successfully"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var inputUser model.User

	if err := ctx.ShouldBindJSON(&inputUser); err != nil {
		util.NewErrorResponse(ctx, "Invalid request", "Invalid request format", "4001", 400)
		return
	}

	if err := c.authService.Login(&inputUser); err != nil {
		switch err {
		case service.ErrUserNotFound:
			util.NewErrorResponse(ctx, "User not found", "User not found", "4011", 401)
		case service.ErrInvalidCredentials:
			util.NewErrorResponse(ctx, "Invalid credentials", "Invalid username or password", "4012", 401)
		default:
			util.NewErrorResponse(ctx, "Internal Server Error", "Internal Server Error", "5001", 500)
		}
		return
	}

	ctx.JSON(200, gin.H{"message": "Login successful"})
}
