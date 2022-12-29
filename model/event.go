package model

type Event struct {
	Id        string `json:"id,omitempty" bson:"id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Event     string `json:"event,omitempty" bson:"event,omitempty"`
	Date      string `json:"date,omitempty" bson:"date,omitempty"`
	Time      string `json:"time,omitempty" bson:"time,omitempty"`
	CreatedAt string `json:"createdat,omitempty" bson:"createdat,omitempty"`
	UpdatedAt string `json:"updatedat,omitempty" bson:"updatedat,omitempty"`
	CreatedBy string `json:"createdby,omitempty" bson:"createdby,omitempty"`
	UpdatedBy string `json:"updatedby,omitempty" bson:"updatedby,omitempty"`
}