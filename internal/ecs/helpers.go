package ecs

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

func Fill(thingToFill reflect.Value, compsToFillWith map[reflect.Type]Component) {
	v := thingToFill.Elem()

	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		for _, comp := range compsToFillWith {
			if reflect.TypeOf(comp) == f.Type() {
				logrus.Trace("system field match, setting")
				f.Set(reflect.ValueOf(comp))
			}
		}
	}
}
