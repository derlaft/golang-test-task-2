package views

import (
	"fmt"
	"testing"

	"configstore/mocks"
	"configstore/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetView(t *testing.T) {

	mock := gomock.NewController(t)
	defer mock.Finish()

	configStore := mocks.NewMockConfigStore(mock)
	views := &ConfigView{
		Store: configStore,
	}

	// test #1: pass some shit as a type name
	res, err := views.Get(&GetRequest{
		Type: "ucucuga",
		Data: "does not matter",
	})

	// should fail
	assert.Nil(t, res, "Got result for unexistent view")
	assert.NotNil(t, err, "Got non-nil error while it should fail")

	// test #2: pass nil request. shoild result in no panic
	res, err = views.Get(nil)
	assert.Nil(t, res, "Got result for unexistent view")
	assert.NotNil(t, err, "Got non-nil error while it should fail")

	// test #3: data not found
	configStore.EXPECT().Get("child", "_404", gomock.Any()).
		Return(fmt.Errorf("Non-existent")).
		Times(1)

	res, err = views.Get(&GetRequest{
		Type: "child",
		Data: "_404",
	})
	assert.Nil(t, res, "Got result for unexistent view")
	assert.NotNil(t, err, "Got non-nil error while it should fail")

	// test #4: data is returned
	configStore.EXPECT().Get("child", "existentSample", gomock.Any()).
		Do(func(Type, ID string, result interface{}) {
			res, typeOk := result.(*models.ChildModel)
			assert.True(t, typeOk, "incorrect type")
			assert.NotNil(t, res)
			res.Magic = 31337
		}).
		Return(nil).
		Times(1)

	// do requesting
	res, err = views.Get(&GetRequest{
		Type: "child",
		Data: "existentSample",
	})
	assert.NotNil(t, res, "Got nil result for existent config")
	assert.Nil(t, err, "Got non-nil error while it should not fail")

	// convert result to model, check it
	resChild, typeOk := res.(*models.ChildModel)
	assert.True(t, typeOk)
	assert.Equal(t, 31337, resChild.Magic, "Incorrect returned value")
}
