package routers

import (
	"example.com/serw/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("./sites/*")

	v1 := r.Group("/v1")
	{
		v1.GET("/students", controllers.AllStudents)
		v1.POST("/students", controllers.CreateStudent)
		v1.GET("/students/:id", controllers.FindStudent)
		v1.PATCH("/students/:id", controllers.UpdateStudent)
		v1.DELETE("/students/:id", controllers.DeleteStudent)
	}
	viewPage := r.Group("/viewPage")
	{
		viewPage.GET("/", func(ctx *gin.Context) {
			ctx.HTML(200, "index.tmpl", gin.H{
				"title": "Database",
				"desc":  "JJ",
				"studs": controllers.ShowAllStudents(),
			})
		})
	}
	return r
}
