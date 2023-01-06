package main

import (
	"context"
	"remainder_app/functions"
	"remainder_app/model"
	"testing"
	"github.com/tj/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mockCollection struct {
}

func (m *mockCollection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	c := &mongo.InsertOneResult{}
	return c, nil
}

func (m *mockCollection) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	c := &mongo.SingleResult{}
	return c
}

func (m *mockCollection) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	c := &mongo.Cursor{}
	return c, nil
}

func (m *mockCollection) ReplaceOne(ctx context.Context, filter interface{},
	replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	c := &mongo.UpdateResult{}
	return c, nil
}
func (m *mockCollection) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	c := &mongo.DeleteResult{}
	return c, nil
}

type answervalid struct {
	err bool
}

func TestInsertData(t *testing.T) {
	mockCol := &mockCollection{}
	sol := []answervalid{{
		err: false}, {
		err: true}, {
		err: true}, {
		err: true}, {
		err: true},
	}
	tests := []model.Event{{
		Id:        "1",
		Name:      "you",
		Event:     "walking",
		Date:      "2023-12-03",
		Time:      "12:36",
		CreatedAt: "12:37",
		UpdatedAt: "12:37",
		CreatedBy: "harsha",
		UpdatedBy: "harsha"},
		{
			Id:        "2",
			Name:      "",
			Event:     "walking",
			Date:      "2023-12-03",
			Time:      "12:36",
			CreatedAt: "12:37",
			UpdatedAt: "12:37",
			CreatedBy: "harsha",
			UpdatedBy: "harsha"},
		{
			Id:        "3",
			Name:      "you",
			Event:     "",
			Date:      "2023-12-03",
			Time:      "12:36",
			CreatedAt: "12:37",
			UpdatedAt: "12:37",
			CreatedBy: "harsha",
			UpdatedBy: "harsha"},
		{
			Id:        "4",
			Name:      "you",
			Event:     "walking",
			Date:      "2023-12-03",
			Time:      "12:36",
			CreatedAt: "12:37",
			UpdatedAt: "12:37",
			CreatedBy: "",
			UpdatedBy: "harsha"},
		{
			Id:        "5",
			Name:      "you",
			Event:     "walking",
			Date:      "2022-12-27",
			Time:      "12:36",
			CreatedAt: "12:37",
			UpdatedAt: "12:37",
			CreatedBy: "harsha",
			UpdatedBy: "harsha"}}

	for i, tt := range tests {
		res, err := functions.InsertData(mockCol, tt)
		assert.IsType(t, &mongo.InsertOneResult{}, res)
		if sol[i].err == true {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
	res, err := functions.InsertData(mockCol, model.Event{Id: "5", Name: "you", Event: "walking", Date: "27-12-2022", Time: "12:36", CreatedAt: "12:37", UpdatedAt: "12:37", CreatedBy: "harsha", UpdatedBy: "harsha"})
	res2 := functions.FindDataById(mockCol, model.Event{Id: "5"})
	res3, err2 := functions.GetData(mockCol, model.Event{})
	res4, err3 := functions.UpdateData(mockCol, model.Event{Id: "5"})
	res5, err4 := functions.DeleteData(mockCol, model.Event{Id: "5"})
	assert.IsType(t, &mongo.DeleteResult{}, res5)
	assert.Nil(t, err4)
	assert.IsType(t, &mongo.UpdateResult{}, res4)
	assert.Nil(t, err3)
	assert.IsType(t, &mongo.Cursor{}, res3)
	assert.Nil(t, err2)
	assert.IsType(t, &mongo.SingleResult{}, res2)
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
}
