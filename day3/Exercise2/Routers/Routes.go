package Routes

import (
	"Exercise2/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/student-api")
	{
		grp1.GET("student", Controllers.GetStudents)
		grp1.POST("student", Controllers.AddStudentDetails)
		grp1.GET("student/:id", Controllers.GetStudentByID)
		grp1.PUT("student/:id", Controllers.UpdateStudentDetails)
		grp1.DELETE("student/:id", Controllers.DeleteStudentDetails)
	}
	return r
}