package models

import (
	"fmt"

	"errors"
	"gopkg.in/mgo.v2/bson"

	"strconv"
	"strings"
)

type Alert struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name    string        `json:"name"`
	Compare string        `json:"compare"`
	Value   float64       `json:"value"`
}

func (a *Alert) AsMap() map[string]string {
	alert := make(map[string]string)

	alert["id"] = string(a.ID)
	alert["name"] = a.Name
	alert["value"] = fmt.Sprintf("%v", a.Value)

	return alert
}

// example input "p SBD > 0.000020"
func NewAlert(input string) (Alert, error) {
	var alert Alert
	values := strings.Split(input, " ")

	if len(values) != 4 {
		err := errors.New("input not valid")
		return alert, err
	}

	name := fmt.Sprintf("BTC_%s", strings.ToUpper(values[1]))
	compare := values[2]
	value, err := strconv.ParseFloat(values[3], 64)

	if err != nil {
		return alert, err
	}

	alert = Alert{
		ID:      bson.NewObjectId(),
		Name:    name,
		Compare: compare,
		Value:   value,
	}

	return alert, nil
}

func RemoveAlert(user User, alert Alert) error {
	return nil
}
