// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	storage "github.com/kyma-project/test-infra/development/tools/pkg/gcscleaner/storage"
	mock "github.com/stretchr/testify/mock"
)

// BucketIterator is an autogenerated mock type for the BucketIterator type
type BucketIterator struct {
	mock.Mock
}

// Next provides a mock function with given fields:
func (_m *BucketIterator) Next() (storage.BucketAttrs, error) {
	ret := _m.Called()

	var r0 storage.BucketAttrs
	if rf, ok := ret.Get(0).(func() storage.BucketAttrs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.BucketAttrs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}