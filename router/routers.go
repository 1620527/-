package router

import (
	"fmt"

	"ginStudy/config"
	"ginStudy/controllers"

	"ginStudy/logger"

	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Usergroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("usergroup middleware")
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	r, err := logger.NewLoggerConfig("log.log", []string{"/skip"})
	if err != nil {
		fmt.Printf("初始化日志配置失败: %v\n", err)
		return nil
	}
	r.Use(CORSMiddleware())

	store, _ := sessions_redis.NewStore(10, "tcp", config.Conf.Redis.RedisAddr, config.Conf.Redis.RedisPass, []byte("secret"))

	r.Use(sessions.Sessions("mysession", store))

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin",
			"status":  "success",
		})
	})

	user := r.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
	}

	player := r.Group("/player")
	{
		player.POST("/list", controllers.PlayerController{}.GetPlayers)
	}

	vote := r.Group("/vote")
	{
		vote.POST("/add", controllers.VoteController{}.AddVote)
	}

	r.POST("ranking", controllers.PlayerController{}.GetRanking)

	return r
}
