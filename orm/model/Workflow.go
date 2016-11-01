package model

import "github.com/jinzhu/gorm"

type Workflow struct {
	gorm.Model
	Name string
	Tasks []WorkflowTask
}

type WorkflowTask struct {
	gorm.Model
	Name string
	Type string
	SubType string
	Workflow Workflow
	WorkflowID int
}

