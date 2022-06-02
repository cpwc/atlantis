// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	github "github.com/google/go-github/v31/github"
)

func AnySliceOfPtrToGithubCheckRun() []*github.CheckRun {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*([]*github.CheckRun))(nil)).Elem()))
	var nullValue []*github.CheckRun
	return nullValue
}

func EqSliceOfPtrToGithubCheckRun(value []*github.CheckRun) []*github.CheckRun {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue []*github.CheckRun
	return nullValue
}

func NotEqSliceOfPtrToGithubCheckRun(value []*github.CheckRun) []*github.CheckRun {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue []*github.CheckRun
	return nullValue
}

func SliceOfPtrToGithubCheckRunThat(matcher pegomock.ArgumentMatcher) []*github.CheckRun {
	pegomock.RegisterMatcher(matcher)
	var nullValue []*github.CheckRun
	return nullValue
}
