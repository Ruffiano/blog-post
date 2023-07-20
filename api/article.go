package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/lib/pq"
	db "github.com/ruffiano/blog-post/db/sqlc"
	"github.com/ruffiano/blog-post/token"

	"github.com/gin-gonic/gin"
)

/*
createArticleRequest ->POST
*/
type createArticleRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required,min=6"`
}

type createArticleResonse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int64     `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (server *Server) createArticle(ctx *gin.Context) {
	var req createArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateArticleParams{
		UserID:    authPayload.UserID,
		Title:     req.Title,
		Content:   req.Content,
		UpdatedAt: time.Now(),
	}

	article, err := server.store.CreateArticle(ctx, arg)
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
	resp := createArticleResonse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		UserID:    article.UserID,
		UpdatedAt: article.CreatedAt,
		CreatedAt: article.CreatedAt,
	}
	ctx.JSON(http.StatusOK, resp)
}

type updateArticleRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required,min=6"`
}

/*
updateArticleRequest ->POST
*/
func (server *Server) updateArticle(ctx *gin.Context) {
	var req updateArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.UpdateArticleParams{
		UserID:    authPayload.UserID,
		Title:     sql.NullString{String: req.Title, Valid: true},
		Content:   sql.NullString{String: req.Content, Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	article, err := server.store.UpdateArticle(ctx, arg)
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

	resp := createArticleResonse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		UserID:    article.UserID,
		UpdatedAt: article.CreatedAt,
		CreatedAt: article.CreatedAt,
	}
	ctx.JSON(http.StatusOK, resp)
}
