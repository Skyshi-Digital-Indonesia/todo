package data

import (
	"time"

	"todo/fitur/todos"

	"gorm.io/gorm"
)

type Todo struct {
	Todo_id           uint `gorm:"primarykey"`
	Title             string
	Priority          string
	IsActive          bool
	Created_at        time.Time
	Updated_at        time.Time
	Deleted_at        gorm.DeletedAt `gorm:"index"`
	Activity_Group_id uint
}

func Todata(data todos.TodoEntities) Todo {
	return Todo{

		Todo_id:           data.ID,
		Title:             data.Title,
		Priority:          data.Priority,
		IsActive:          data.IsActive,
		Activity_Group_id: data.ActivitiesID,
		Created_at:        data.Createdat,
		Updated_at:        data.Updatedat,
	}
}

func (data *Todo) ModelsToCore() todos.TodoEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return todos.TodoEntities{
		ID:           data.Todo_id,
		Title:        data.Title,
		Priority:     data.Priority,
		IsActive:     data.IsActive,
		Createdat:    data.Created_at,
		Updatedat:    data.Updated_at,
		ActivitiesID: data.Activity_Group_id,
	}
}

func ToCore(data Todo) todos.TodoEntities {
	return todos.TodoEntities{
		ID:           data.Todo_id,
		Title:        data.Title,
		Priority:     data.Priority,
		IsActive:     data.IsActive,
		Updatedat:    data.Updated_at,
		ActivitiesID: data.Activity_Group_id,
	}
}

func ListModelTOCore(dataModel []Todo) []todos.TodoEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []todos.TodoEntities
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
