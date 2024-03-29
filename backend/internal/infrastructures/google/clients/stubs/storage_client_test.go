// Code generated by mockery. DO NOT EDIT.

package stubs

import (
	context "context"
	io "io"

	iampb "cloud.google.com/go/iam/apiv1/iampb"

	mock "github.com/stretchr/testify/mock"

	storage "cloud.google.com/go/storage"
)

// storageClient is an autogenerated mock type for the storageClient type
type storageClient struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *storageClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ComposeObject provides a mock function with given fields: ctx, req, opts
func (_m *storageClient) ComposeObject(ctx context.Context, req *storage.composeObjectRequest, opts ...storage.storageOption) (*storage.ObjectAttrs, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.ObjectAttrs
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.composeObjectRequest, ...storage.storageOption) (*storage.ObjectAttrs, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.composeObjectRequest, ...storage.storageOption) *storage.ObjectAttrs); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ObjectAttrs)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.composeObjectRequest, ...storage.storageOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBucket provides a mock function with given fields: ctx, project, bucket, attrs, opts
func (_m *storageClient) CreateBucket(ctx context.Context, project string, bucket string, attrs *storage.BucketAttrs, opts ...storage.storageOption) (*storage.BucketAttrs, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, project, bucket, attrs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.BucketAttrs
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *storage.BucketAttrs, ...storage.storageOption) (*storage.BucketAttrs, error)); ok {
		return rf(ctx, project, bucket, attrs, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *storage.BucketAttrs, ...storage.storageOption) *storage.BucketAttrs); ok {
		r0 = rf(ctx, project, bucket, attrs, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.BucketAttrs)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *storage.BucketAttrs, ...storage.storageOption) error); ok {
		r1 = rf(ctx, project, bucket, attrs, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateHMACKey provides a mock function with given fields: ctx, project, serviceAccountEmail, opts
func (_m *storageClient) CreateHMACKey(ctx context.Context, project string, serviceAccountEmail string, opts ...storage.storageOption) (*storage.HMACKey, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, project, serviceAccountEmail)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.HMACKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...storage.storageOption) (*storage.HMACKey, error)); ok {
		return rf(ctx, project, serviceAccountEmail, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...storage.storageOption) *storage.HMACKey); ok {
		r0 = rf(ctx, project, serviceAccountEmail, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.HMACKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...storage.storageOption) error); ok {
		r1 = rf(ctx, project, serviceAccountEmail, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNotification provides a mock function with given fields: ctx, bucket, n, opts
func (_m *storageClient) CreateNotification(ctx context.Context, bucket string, n *storage.Notification, opts ...storage.storageOption) (*storage.Notification, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, n)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.Notification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *storage.Notification, ...storage.storageOption) (*storage.Notification, error)); ok {
		return rf(ctx, bucket, n, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *storage.Notification, ...storage.storageOption) *storage.Notification); ok {
		r0 = rf(ctx, bucket, n, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *storage.Notification, ...storage.storageOption) error); ok {
		r1 = rf(ctx, bucket, n, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucket provides a mock function with given fields: ctx, bucket, conds, opts
func (_m *storageClient) DeleteBucket(ctx context.Context, bucket string, conds *storage.BucketConditions, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, conds)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *storage.BucketConditions, ...storage.storageOption) error); ok {
		r0 = rf(ctx, bucket, conds, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBucketACL provides a mock function with given fields: ctx, bucket, entity, opts
func (_m *storageClient) DeleteBucketACL(ctx context.Context, bucket string, entity storage.ACLEntity, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, entity)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, storage.ACLEntity, ...storage.storageOption) error); ok {
		r0 = rf(ctx, bucket, entity, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDefaultObjectACL provides a mock function with given fields: ctx, bucket, entity, opts
func (_m *storageClient) DeleteDefaultObjectACL(ctx context.Context, bucket string, entity storage.ACLEntity, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, entity)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, storage.ACLEntity, ...storage.storageOption) error); ok {
		r0 = rf(ctx, bucket, entity, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteHMACKey provides a mock function with given fields: ctx, project, accessID, opts
func (_m *storageClient) DeleteHMACKey(ctx context.Context, project string, accessID string, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, project, accessID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...storage.storageOption) error); ok {
		r0 = rf(ctx, project, accessID, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNotification provides a mock function with given fields: ctx, bucket, id, opts
func (_m *storageClient) DeleteNotification(ctx context.Context, bucket string, id string, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...storage.storageOption) error); ok {
		r0 = rf(ctx, bucket, id, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteObject provides a mock function with given fields: ctx, bucket, object, gen, conds, opts
func (_m *storageClient) DeleteObject(ctx context.Context, bucket string, object string, gen int64, conds *storage.Conditions, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, object, gen, conds)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64, *storage.Conditions, ...storage.storageOption) error); ok {
		r0 = rf(ctx, bucket, object, gen, conds, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteObjectACL provides a mock function with given fields: ctx, bucket, object, entity, opts
func (_m *storageClient) DeleteObjectACL(ctx context.Context, bucket string, object string, entity storage.ACLEntity, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, object, entity)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, storage.ACLEntity, ...storage.storageOption) error); ok {
		r0 = rf(ctx, bucket, object, entity, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBucket provides a mock function with given fields: ctx, bucket, conds, opts
func (_m *storageClient) GetBucket(ctx context.Context, bucket string, conds *storage.BucketConditions, opts ...storage.storageOption) (*storage.BucketAttrs, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, conds)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.BucketAttrs
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *storage.BucketConditions, ...storage.storageOption) (*storage.BucketAttrs, error)); ok {
		return rf(ctx, bucket, conds, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *storage.BucketConditions, ...storage.storageOption) *storage.BucketAttrs); ok {
		r0 = rf(ctx, bucket, conds, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.BucketAttrs)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *storage.BucketConditions, ...storage.storageOption) error); ok {
		r1 = rf(ctx, bucket, conds, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHMACKey provides a mock function with given fields: ctx, project, accessID, opts
func (_m *storageClient) GetHMACKey(ctx context.Context, project string, accessID string, opts ...storage.storageOption) (*storage.HMACKey, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, project, accessID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.HMACKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...storage.storageOption) (*storage.HMACKey, error)); ok {
		return rf(ctx, project, accessID, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...storage.storageOption) *storage.HMACKey); ok {
		r0 = rf(ctx, project, accessID, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.HMACKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...storage.storageOption) error); ok {
		r1 = rf(ctx, project, accessID, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIamPolicy provides a mock function with given fields: ctx, resource, version, opts
func (_m *storageClient) GetIamPolicy(ctx context.Context, resource string, version int32, opts ...storage.storageOption) (*iampb.Policy, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, resource, version)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *iampb.Policy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int32, ...storage.storageOption) (*iampb.Policy, error)); ok {
		return rf(ctx, resource, version, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int32, ...storage.storageOption) *iampb.Policy); ok {
		r0 = rf(ctx, resource, version, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iampb.Policy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int32, ...storage.storageOption) error); ok {
		r1 = rf(ctx, resource, version, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObject provides a mock function with given fields: ctx, bucket, object, gen, encryptionKey, conds, opts
func (_m *storageClient) GetObject(ctx context.Context, bucket string, object string, gen int64, encryptionKey []byte, conds *storage.Conditions, opts ...storage.storageOption) (*storage.ObjectAttrs, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, object, gen, encryptionKey, conds)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.ObjectAttrs
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64, []byte, *storage.Conditions, ...storage.storageOption) (*storage.ObjectAttrs, error)); ok {
		return rf(ctx, bucket, object, gen, encryptionKey, conds, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64, []byte, *storage.Conditions, ...storage.storageOption) *storage.ObjectAttrs); ok {
		r0 = rf(ctx, bucket, object, gen, encryptionKey, conds, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ObjectAttrs)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64, []byte, *storage.Conditions, ...storage.storageOption) error); ok {
		r1 = rf(ctx, bucket, object, gen, encryptionKey, conds, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServiceAccount provides a mock function with given fields: ctx, project, opts
func (_m *storageClient) GetServiceAccount(ctx context.Context, project string, opts ...storage.storageOption) (string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, project)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...storage.storageOption) (string, error)); ok {
		return rf(ctx, project, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...storage.storageOption) string); ok {
		r0 = rf(ctx, project, opts...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...storage.storageOption) error); ok {
		r1 = rf(ctx, project, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBucketACLs provides a mock function with given fields: ctx, bucket, opts
func (_m *storageClient) ListBucketACLs(ctx context.Context, bucket string, opts ...storage.storageOption) ([]storage.ACLRule, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []storage.ACLRule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...storage.storageOption) ([]storage.ACLRule, error)); ok {
		return rf(ctx, bucket, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...storage.storageOption) []storage.ACLRule); ok {
		r0 = rf(ctx, bucket, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]storage.ACLRule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...storage.storageOption) error); ok {
		r1 = rf(ctx, bucket, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBuckets provides a mock function with given fields: ctx, project, opts
func (_m *storageClient) ListBuckets(ctx context.Context, project string, opts ...storage.storageOption) *storage.BucketIterator {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, project)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.BucketIterator
	if rf, ok := ret.Get(0).(func(context.Context, string, ...storage.storageOption) *storage.BucketIterator); ok {
		r0 = rf(ctx, project, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.BucketIterator)
		}
	}

	return r0
}

// ListDefaultObjectACLs provides a mock function with given fields: ctx, bucket, opts
func (_m *storageClient) ListDefaultObjectACLs(ctx context.Context, bucket string, opts ...storage.storageOption) ([]storage.ACLRule, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []storage.ACLRule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...storage.storageOption) ([]storage.ACLRule, error)); ok {
		return rf(ctx, bucket, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...storage.storageOption) []storage.ACLRule); ok {
		r0 = rf(ctx, bucket, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]storage.ACLRule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...storage.storageOption) error); ok {
		r1 = rf(ctx, bucket, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHMACKeys provides a mock function with given fields: ctx, project, serviceAccountEmail, showDeletedKeys, opts
func (_m *storageClient) ListHMACKeys(ctx context.Context, project string, serviceAccountEmail string, showDeletedKeys bool, opts ...storage.storageOption) *storage.HMACKeysIterator {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, project, serviceAccountEmail, showDeletedKeys)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.HMACKeysIterator
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool, ...storage.storageOption) *storage.HMACKeysIterator); ok {
		r0 = rf(ctx, project, serviceAccountEmail, showDeletedKeys, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.HMACKeysIterator)
		}
	}

	return r0
}

// ListNotifications provides a mock function with given fields: ctx, bucket, opts
func (_m *storageClient) ListNotifications(ctx context.Context, bucket string, opts ...storage.storageOption) (map[string]*storage.Notification, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[string]*storage.Notification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...storage.storageOption) (map[string]*storage.Notification, error)); ok {
		return rf(ctx, bucket, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...storage.storageOption) map[string]*storage.Notification); ok {
		r0 = rf(ctx, bucket, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*storage.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...storage.storageOption) error); ok {
		r1 = rf(ctx, bucket, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListObjectACLs provides a mock function with given fields: ctx, bucket, object, opts
func (_m *storageClient) ListObjectACLs(ctx context.Context, bucket string, object string, opts ...storage.storageOption) ([]storage.ACLRule, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, object)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []storage.ACLRule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...storage.storageOption) ([]storage.ACLRule, error)); ok {
		return rf(ctx, bucket, object, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...storage.storageOption) []storage.ACLRule); ok {
		r0 = rf(ctx, bucket, object, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]storage.ACLRule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...storage.storageOption) error); ok {
		r1 = rf(ctx, bucket, object, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListObjects provides a mock function with given fields: ctx, bucket, q, opts
func (_m *storageClient) ListObjects(ctx context.Context, bucket string, q *storage.Query, opts ...storage.storageOption) *storage.ObjectIterator {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, q)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.ObjectIterator
	if rf, ok := ret.Get(0).(func(context.Context, string, *storage.Query, ...storage.storageOption) *storage.ObjectIterator); ok {
		r0 = rf(ctx, bucket, q, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ObjectIterator)
		}
	}

	return r0
}

// LockBucketRetentionPolicy provides a mock function with given fields: ctx, bucket, conds, opts
func (_m *storageClient) LockBucketRetentionPolicy(ctx context.Context, bucket string, conds *storage.BucketConditions, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, conds)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *storage.BucketConditions, ...storage.storageOption) error); ok {
		r0 = rf(ctx, bucket, conds, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRangeReader provides a mock function with given fields: ctx, params, opts
func (_m *storageClient) NewRangeReader(ctx context.Context, params *storage.newRangeReaderParams, opts ...storage.storageOption) (*storage.Reader, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.Reader
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.newRangeReaderParams, ...storage.storageOption) (*storage.Reader, error)); ok {
		return rf(ctx, params, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.newRangeReaderParams, ...storage.storageOption) *storage.Reader); ok {
		r0 = rf(ctx, params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.Reader)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.newRangeReaderParams, ...storage.storageOption) error); ok {
		r1 = rf(ctx, params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenWriter provides a mock function with given fields: params, opts
func (_m *storageClient) OpenWriter(params *storage.openWriterParams, opts ...storage.storageOption) (*io.PipeWriter, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *io.PipeWriter
	var r1 error
	if rf, ok := ret.Get(0).(func(*storage.openWriterParams, ...storage.storageOption) (*io.PipeWriter, error)); ok {
		return rf(params, opts...)
	}
	if rf, ok := ret.Get(0).(func(*storage.openWriterParams, ...storage.storageOption) *io.PipeWriter); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*io.PipeWriter)
		}
	}

	if rf, ok := ret.Get(1).(func(*storage.openWriterParams, ...storage.storageOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RewriteObject provides a mock function with given fields: ctx, req, opts
func (_m *storageClient) RewriteObject(ctx context.Context, req *storage.rewriteObjectRequest, opts ...storage.storageOption) (*storage.rewriteObjectResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.rewriteObjectResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.rewriteObjectRequest, ...storage.storageOption) (*storage.rewriteObjectResponse, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.rewriteObjectRequest, ...storage.storageOption) *storage.rewriteObjectResponse); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.rewriteObjectResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.rewriteObjectRequest, ...storage.storageOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetIamPolicy provides a mock function with given fields: ctx, resource, policy, opts
func (_m *storageClient) SetIamPolicy(ctx context.Context, resource string, policy *iampb.Policy, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, resource, policy)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *iampb.Policy, ...storage.storageOption) error); ok {
		r0 = rf(ctx, resource, policy, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TestIamPermissions provides a mock function with given fields: ctx, resource, permissions, opts
func (_m *storageClient) TestIamPermissions(ctx context.Context, resource string, permissions []string, opts ...storage.storageOption) ([]string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, resource, permissions)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, ...storage.storageOption) ([]string, error)); ok {
		return rf(ctx, resource, permissions, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, ...storage.storageOption) []string); ok {
		r0 = rf(ctx, resource, permissions, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string, ...storage.storageOption) error); ok {
		r1 = rf(ctx, resource, permissions, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBucket provides a mock function with given fields: ctx, bucket, uattrs, conds, opts
func (_m *storageClient) UpdateBucket(ctx context.Context, bucket string, uattrs *storage.BucketAttrsToUpdate, conds *storage.BucketConditions, opts ...storage.storageOption) (*storage.BucketAttrs, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, uattrs, conds)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.BucketAttrs
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *storage.BucketAttrsToUpdate, *storage.BucketConditions, ...storage.storageOption) (*storage.BucketAttrs, error)); ok {
		return rf(ctx, bucket, uattrs, conds, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *storage.BucketAttrsToUpdate, *storage.BucketConditions, ...storage.storageOption) *storage.BucketAttrs); ok {
		r0 = rf(ctx, bucket, uattrs, conds, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.BucketAttrs)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *storage.BucketAttrsToUpdate, *storage.BucketConditions, ...storage.storageOption) error); ok {
		r1 = rf(ctx, bucket, uattrs, conds, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBucketACL provides a mock function with given fields: ctx, bucket, entity, role, opts
func (_m *storageClient) UpdateBucketACL(ctx context.Context, bucket string, entity storage.ACLEntity, role storage.ACLRole, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, entity, role)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, storage.ACLEntity, storage.ACLRole, ...storage.storageOption) error); ok {
		r0 = rf(ctx, bucket, entity, role, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDefaultObjectACL provides a mock function with given fields: ctx, bucket, entity, role, opts
func (_m *storageClient) UpdateDefaultObjectACL(ctx context.Context, bucket string, entity storage.ACLEntity, role storage.ACLRole, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, entity, role)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, storage.ACLEntity, storage.ACLRole, ...storage.storageOption) error); ok {
		r0 = rf(ctx, bucket, entity, role, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateHMACKey provides a mock function with given fields: ctx, project, serviceAccountEmail, accessID, attrs, opts
func (_m *storageClient) UpdateHMACKey(ctx context.Context, project string, serviceAccountEmail string, accessID string, attrs *storage.HMACKeyAttrsToUpdate, opts ...storage.storageOption) (*storage.HMACKey, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, project, serviceAccountEmail, accessID, attrs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.HMACKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *storage.HMACKeyAttrsToUpdate, ...storage.storageOption) (*storage.HMACKey, error)); ok {
		return rf(ctx, project, serviceAccountEmail, accessID, attrs, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *storage.HMACKeyAttrsToUpdate, ...storage.storageOption) *storage.HMACKey); ok {
		r0 = rf(ctx, project, serviceAccountEmail, accessID, attrs, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.HMACKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, *storage.HMACKeyAttrsToUpdate, ...storage.storageOption) error); ok {
		r1 = rf(ctx, project, serviceAccountEmail, accessID, attrs, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateObject provides a mock function with given fields: ctx, bucket, object, uattrs, gen, encryptionKey, conds, opts
func (_m *storageClient) UpdateObject(ctx context.Context, bucket string, object string, uattrs *storage.ObjectAttrsToUpdate, gen int64, encryptionKey []byte, conds *storage.Conditions, opts ...storage.storageOption) (*storage.ObjectAttrs, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, object, uattrs, gen, encryptionKey, conds)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage.ObjectAttrs
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *storage.ObjectAttrsToUpdate, int64, []byte, *storage.Conditions, ...storage.storageOption) (*storage.ObjectAttrs, error)); ok {
		return rf(ctx, bucket, object, uattrs, gen, encryptionKey, conds, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *storage.ObjectAttrsToUpdate, int64, []byte, *storage.Conditions, ...storage.storageOption) *storage.ObjectAttrs); ok {
		r0 = rf(ctx, bucket, object, uattrs, gen, encryptionKey, conds, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ObjectAttrs)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *storage.ObjectAttrsToUpdate, int64, []byte, *storage.Conditions, ...storage.storageOption) error); ok {
		r1 = rf(ctx, bucket, object, uattrs, gen, encryptionKey, conds, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateObjectACL provides a mock function with given fields: ctx, bucket, object, entity, role, opts
func (_m *storageClient) UpdateObjectACL(ctx context.Context, bucket string, object string, entity storage.ACLEntity, role storage.ACLRole, opts ...storage.storageOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bucket, object, entity, role)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, storage.ACLEntity, storage.ACLRole, ...storage.storageOption) error); ok {
		r0 = rf(ctx, bucket, object, entity, role, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newStorageClient creates a new instance of storageClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newStorageClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *storageClient {
	mock := &storageClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
