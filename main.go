package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// func main() {
// 	r := gin.Default()
// 	store, err := redisStore.NewStore(1024, "tcp", "127.0.0.1:6379", "112233", []byte("secret"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	r.Use(sessions.Sessions("mysession", store))

// 	r.GET("/hello", func(c *gin.Context) {
// 		session := sessions.Default(c)

// 		if session.Get("hello") != "world" {
// 			session.Set("hello", "world")
// 			session.Save()
// 		}

// 		c.JSON(200, gin.H{"hello": session.Get("hello")})
// 	})
// 	r.Run(":8000")
// }

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8000")
}
