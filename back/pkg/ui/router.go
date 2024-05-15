package ui

import (
	"cram-school-reserve-server/back/infra"
	"cram-school-reserve-server/back/pkg/handler"
	"cram-school-reserve-server/back/pkg/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const StudentPath = "/student"
const TeacherPath = "/teacher"

func DBContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := infra.ConnectDB()
		if err != nil {
			panic("failed to connect database")
		}
		c.Set("db", db)
		c.Next()
		err = db.Close()
		if err != nil {
			return
		}
	}
}

func SetupRouter() {
	r := gin.Default()

	r.Use(cors.Default())
	r.Use(DBContext())
	r.Use(util.SessionMiddleware(r))
	//first page
	r.GET("/", handler.Home)
	//signup or login or logout
	r.GET("/signup", handler.GetSignup)
	r.POST("/signup", handler.PostSignup)
	r.GET("/login", handler.GetLogin)
	r.POST("/login", handler.PostLogin)
	r.GET("/logout", handler.GetLogout)
	//student
	r.GET(StudentPath, handler.GetStudent)
	r.GET(StudentPath+"/reserve", handler.GetReserve)
	r.POST(StudentPath+"/reserve", handler.PostReserve)
	//teacher
	r.GET(TeacherPath, handler.GetTeacher)
	r.POST(TeacherPath+"/delete_student", handler.PostDeleteStudent)
	r.POST(TeacherPath+"/edit_student_class", handler.PostEditStudentClass)
	r.POST(TeacherPath+"/update_student_class", handler.PostUpdateStudentClass)
	r.GET(TeacherPath+"/edit_student_name_password", handler.GetEditStudentNameAndPassword)
	r.POST(TeacherPath+"/edit_student_name_password", handler.PostEditStudentNameAndPassword)
	r.GET(TeacherPath+"/delete_teacher", handler.GetDeleteTeacher)
	r.GET(TeacherPath+"/edit_teacher_name_password", handler.GetEditTeacherNameAndPassword)
	r.POST(TeacherPath+"/edit_teacher_name_password", handler.PostEditTeacherNameAndPassword)
	r.GET(TeacherPath+"/edit_form", handler.GetEditForm)
	r.POST(TeacherPath+"/edit_form", handler.PostEditForm)
	r.GET(TeacherPath+"/create_form", handler.GetCreateForm)
	r.POST(TeacherPath+"/create_form", handler.PostCreateForm)
	r.GET(TeacherPath+"/delete_form", handler.GetDeleteForm)
	//output
	r.GET(TeacherPath+"/download_reservations", handler.GetCreateCSVForOneForm)
	r.GET(TeacherPath+"/download_all_reservations", handler.GetCreateCSVForAllForms)
	err := r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
	if err != nil {
		return
	}
}
