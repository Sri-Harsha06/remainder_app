package main
import (
	"context"
	"testing"
	"remainder_app/functions"
	"github.com/tj/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"remainder_app/model"
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

func TestInsertData(t *testing.T) {
	mockCol := &mockCollection{}
	res, err := functions.InsertData(mockCol, model.Event{Id:"5", Name:"you", Event:"walking", Date:"27-12-2022", Time:"12:36", CreatedAt:"12:37", UpdatedAt:"12:37", CreatedBy:"harsha", UpdatedBy:"harsha"})
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