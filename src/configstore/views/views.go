package views

import (
	"configstore/models"

	"fmt"
)

type ConfigView struct {
	Store models.ConfigStore
}

type GetRequest struct {
	Type string
	Data string
}

func (cv *ConfigView) Get(input interface{}) (interface{}, error) {

	// step 0 -- input parameters
	req, ok := input.(*GetRequest)
	if !ok || req == nil {
		return nil, fmt.Errorf("Bad type")
	}

	// step 1 -- create output struct
	config, err := models.NewConfig(req.Type)
	if err != nil {
		return nil, err
	}

	// step 2 -- read it
	err = cv.Store.Get(req.Type, req.Data, config)
	if err != nil {
		return nil, err
	}

	// step 3 -- output
	return config, nil
}
