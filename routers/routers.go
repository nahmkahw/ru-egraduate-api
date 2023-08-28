package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	"RU-Smart-Workspace/ru-smart-api/handlers/studenth"
	"RU-Smart-Workspace/ru-smart-api/logger"
	"RU-Smart-Workspace/ru-smart-api/repositories/studentr"
	"RU-Smart-Workspace/ru-smart-api/services/students"

	"RU-Smart-Workspace/ru-smart-api/middlewares"
)

func Setup(router *gin.Engine, oracle_db *sqlx.DB, redis_cache *redis.Client) {

	jsonFileLogger, err := logger.NewJSONFileLogger("./logger/app.log")
	if err != nil {
		// Handle error
	}

	router.Use(logger.ErrorLogger(jsonFileLogger))

	router.Use(middlewares.NewCorsAccessControl().CorsAccessControl())

	router.GET("/healthz", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "The service works normally...",
		})
	})

	googleAuth := router.Group("/google")
	{
		studentRepo := studentr.NewStudentRepo(oracle_db)
		studentService := students.NewStudentServices(studentRepo, redis_cache)
		studentHandler := studenth.NewStudentHandlers(studentService)

		googleAuth.POST("/authorization", middlewares.GoogleAuth, studentHandler.Authentication)
		googleAuth.POST("/authorization-test", studentHandler.AuthenticationTest)
		googleAuth.POST("/authorization-service", studentHandler.AuthenticationService)
		googleAuth.POST("/authorization-redirect", studentHandler.AuthenticationRedirect)

	}

	student := router.Group("/student")
	{
		studentRepo := studentr.NewStudentRepo(oracle_db)
		studentService := students.NewStudentServices(studentRepo, redis_cache)
		studentHandler := studenth.NewStudentHandlers(studentService)

		student.GET("/profile/:std_code", middlewares.Authorization(redis_cache), studentHandler.GetStudentProfile)

		student.GET("/", studentHandler.GetStudentAll)
	}

	PORT := viper.GetString("ruConnext.port")

	if err := router.Run(PORT); err != nil {
		jsonFileLogger.Error("Failed to start server", err, nil)
	}

}

// func errorLogger(log *logrus.Logger) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Continue to the next middleware or route handler
// 		c.Next()

// 		// Check if any errors occurred during the request handling
// 		err := c.Errors.Last()
// 		if err != nil {
// 			// Log the error
// 			log.WithField("status", c.Writer.Status()).
// 				WithField("method", c.Request.Method).
// 				WithField("path", c.Request.URL.Path).
// 				Error(err.Err)
// 		}
// 	}
// }
