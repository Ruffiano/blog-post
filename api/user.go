package api

import (
	"net/http"
	db "ruffiano/blog-post/db/sqlc"
	"ruffiano/blog-post/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

/*
createUser - >signUp
*/
type userRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Username string `json:"username" binding:"required,alphanum"`
}

type userResponse struct {
	Email             string    `json:"email"`
	Username          string    `json:"username"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Email:             user.Email,
		Username:          user.Username,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req userRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := newUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}

/*
	login - >signIn
*/

// type loinUserRequest struct {
// 	Username string `json:"username" binding:"required, alphanum"`
// 	Password string `json: "password" binding:"required, min=6"`
// }

// type loginResponse struct {
// 	SessionID             uuid.UUID    `json:"session_id"`
// 	AccessToken           string       `json:"access_token`
// 	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
// 	RefreshToken          string       `json:"refresh_token"`
// 	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
// 	User                  userResponse `json:"user"`
// }

// func (server *Server) loginUser(ctx *gin.Context) {
// 	var req loginUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 	}

// 	user, err := server.store.GetUser(ctx, req.Username)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	err = util.CheckPassword(req.Password, user.HashedPassword)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
// 		return
// 	}

// 	accessToken, aceessPayload, err := server.tokenMaker.CreateToken(
// 		user.Username,
// 		server.configAccessTokenDuration,
// 	)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// }
