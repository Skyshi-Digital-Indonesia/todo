package handler

import (
	"time"
	"todo/fitur/activities"
)

type ActivitiesRequest struct {
	ID          uint
	Title       string `json:"title" form:"title"`
	Email       string `json:"email" form:"email"`
	Created_at  time.Time
	Updated_at  time.Time
	Deleteed_at time.Time
}

type UpdateRequest struct {
	ID         uint
	Title      string `json:"title" form:"title"`
	Email      string `json:"email" form:"email"`
	Updated_at time.Time
}

func ActivitiesRequestToUserCore(data ActivitiesRequest) activities.ActivitiesEntities {
	return activities.ActivitiesEntities{
		ID:        data.ID,
		Title:     data.Title,
		Email:     data.Email,
		Createdat: data.Created_at,
		Updatedat: data.Updated_at,
	}
}
func ActivityRequestToUserCore(data UpdateRequest) activities.ActivitiesEntities {
	return activities.ActivitiesEntities{
		ID:        data.ID,
		Title:     data.Title,
		Email:     data.Email,
		Updatedat: data.Updated_at,
	}
}
