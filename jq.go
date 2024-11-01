package main

import (
	"encoding/json"

	"github.com/itchyny/gojq"
	log "github.com/sirupsen/logrus"
)

func jq(query string, input interface{}) (interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(input.(string)), &data); err != nil {
		return nil, err
	}
	gojquery, err := gojq.Parse(query)
	if err != nil {
		log.Fatalln(err)
	}
	log.Debugf("jq: %v", data)
	iter := gojquery.Run(data) // or gojquery.RunWithContext
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return nil, err
		}
		return v, nil
	}
	return nil, nil
}
