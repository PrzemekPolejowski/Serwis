package controllers

import (
	"example.com/serw/models"
	"github.com/gin-gonic/gin"
)

type CreateStudentData struct {
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Age     uint   `json:"age" binding:"required"`
}

type UpdateStudentData struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint   `json:"age"`
}

func ShowAllStudents() []models.Student {
	var Students []models.Student
	models.DB.Find(&Students)
	return Students
	
}

func UpdateStudent(ctx *gin.Context) {
	var student models.Student
	if err := models.DB.Where("id=?", ctx.Param("id")).First(&student).Error; err != nil {
		ctx.JSON(400, gin.H{"error": "Student doesn't exist"})
		return
	}
	var input UpdateStudentData
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&student).Updates(input)
	ctx.JSON(200, gin.H{"data": student})
}

func CreateStudent(ctx *gin.Context) {
	var input CreateStudentData
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	student := models.Student{Name: input.Name, Surname: input.Surname, Age: input.Age}
	models.DB.Create(&student)
	ctx.JSON(200, gin.H{"data": student})
}

func AllStudents(ctx *gin.Context) {
	var Students []models.Student
	models.DB.Find(&Students)
	ctx.JSON(200, gin.H{"data": Students})
}
func FindStudent(ctx *gin.Context) {
	var student models.Student
	if err := models.DB.Where("id=?", ctx.Param("id")).First(&student).Error; err != nil {
		ctx.JSON(400, gin.H{"error": "Cant find student"})
		return
	}
	ctx.JSON(200, gin.H{"data": student})
}
func DeleteStudent(ctx *gin.Context) {
	var student models.Student
	if err := models.DB.Where("id=?", ctx.Param("id")).First(&student).Error; err != nil {
		ctx.JSON(400, gin.H{"error": "Cant find student"})
		return
	}
	models.DB.Delete(&student)
	ctx.JSON(200, gin.H{"data": true})
}
