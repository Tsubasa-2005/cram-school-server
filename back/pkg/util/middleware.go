package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware(r *gin.Engine) gin.HandlerFunc {
	// セッションストアを作成します。
	store := cookie.NewStore([]byte("secret"))

	// セッションミドルウェアを設定します。
	r.Use(sessions.Sessions("session_name", store))
	return func(c *gin.Context) {
		// 新しいセッションを作成するか、既存のセッションを取得します。
		// 実際の実装は使用しているセッションライブラリに依存します。
		session := sessions.Default(c)

		// セッションをコンテキストに設定します。
		c.Set("session", session)

		// 次のハンドラ関数に進みます。
		c.Next()
	}
}
