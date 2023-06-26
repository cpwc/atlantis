// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"

	models "github.com/runatlantis/atlantis/server/models"
)

func AnyModelsVCSHostType() models.VCSHostType {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.VCSHostType))(nil)).Elem()))
	var nullValue models.VCSHostType
	return nullValue
}

func EqModelsVCSHostType(value models.VCSHostType) models.VCSHostType {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.VCSHostType
	return nullValue
}

func NotEqModelsVCSHostType(value models.VCSHostType) models.VCSHostType {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue models.VCSHostType
	return nullValue
}

func ModelsVCSHostTypeThat(matcher pegomock.ArgumentMatcher) models.VCSHostType {
	pegomock.RegisterMatcher(matcher)
	var nullValue models.VCSHostType
	return nullValue
}