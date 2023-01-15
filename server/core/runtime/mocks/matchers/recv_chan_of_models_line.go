// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	models "github.com/runatlantis/atlantis/server/core/runtime/models"
)

func AnyRecvChanOfModelsLine() <-chan models.Line {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(<-chan models.Line))(nil)).Elem()))
	var nullValue <-chan models.Line
	return nullValue
}

func EqRecvChanOfModelsLine(value <-chan models.Line) <-chan models.Line {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue <-chan models.Line
	return nullValue
}

func NotEqRecvChanOfModelsLine(value <-chan models.Line) <-chan models.Line {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue <-chan models.Line
	return nullValue
}

func RecvChanOfModelsLineThat(matcher pegomock.ArgumentMatcher) <-chan models.Line {
	pegomock.RegisterMatcher(matcher)
	var nullValue <-chan models.Line
	return nullValue
}