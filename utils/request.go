package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// ReqToQueryValues converts request struct to ready to encode query values
func ReqToQueryValues(in interface{}) (*url.Values, error) {
	values := url.Values{}
	var vals map[string]interface{}

	reqJson, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(reqJson, &vals); err != nil {
		return nil, err
	}

	for k, v := range vals {
		values.Set(k, fmt.Sprintf("%v", v))
	}

	return &values, nil
}
