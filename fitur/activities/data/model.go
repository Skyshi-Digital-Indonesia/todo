package data

import (
	"time"
	"todo/fitur/activities"
	"todo/fitur/todos/data"

	"gorm.io/gorm"
)

type Activities struct {
	Activity_id uint   `gorm:"primarykey"`
	Title       string `gorm:"type:char(50);not null"`
	Email       string `gorm:"type:varchar(50);unique;not null"`
	Created_at  time.Time
	Updated_at  time.Time
	Deleted_at  gorm.DeletedAt `gorm:"index"`
	Todos       []data.Todo    `gorm:"foreignKey:Activity_Group_id;references:Activity_id"`
}

// register
func FromEntities(dataCore activities.ActivitiesEntities) Activities { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	return Activities{
		Activity_id: dataCore.ID,
		Email:       dataCore.Email,
		Title:       dataCore.Title,
		Created_at:  dataCore.Createdat,
		Updated_at:  dataCore.Updatedat,
	}

}

// profile user
func (dataModel *Activities) ModelsToCore() activities.ActivitiesEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return activities.ActivitiesEntities{
		ID:        dataModel.Activity_id,
		Title:     dataModel.Title,
		Email:     dataModel.Email,
		Updatedat: dataModel.Updated_at,
		Createdat: dataModel.Created_at,
	}
}

func ListModelEntities(datamodel []Activities) []activities.ActivitiesEntities {
	var entities []activities.ActivitiesEntities

	for _, val := range datamodel {
		entities = append(entities, val.ModelsToCore())
	}
	return entities

}

func ToCore(model Activities) activities.ActivitiesEntities {
	return activities.ActivitiesEntities{
		ID:        model.Activity_id,
		Title:     model.Title,
		Email:     model.Email,
		Updatedat: model.Updated_at,
	}

}
