-- name: CreateArticle :one
INSERT INTO article (
    user_id,
    title,
    content,
    updated_at
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetArticle :one
SELECT * FROM article
WHERE user_id =$1 LIMIT 1;

-- name: UpdateArticle :one
UPDATE article
SET
   title = COALESCE(sqlc.narg(title), title),
   content = COALESCE(sqlc.narg(content), content),
   updated_at = COALESCE(sqlc.narg(updated_at), updated_at)
WHERE
   user_id = sqlc.arg(user_id)
RETURNING *;
