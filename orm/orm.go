package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"GoExamples/orm/model"
)


func main()  {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/gotest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&model.Workflow{}, model.WorkflowTask{})

	db.Create(&model.Workflow{Name:"hi"})
	db.Create(&model.WorkflowTask{Name:"test",SubType:"st", Type:"t",WorkflowID:2})

	var workflow model.Workflow
	var workflowTask model.WorkflowTask

	db.First(&workflowTask, 1)

	db.Model(&workflowTask).Related(&workflow)

}