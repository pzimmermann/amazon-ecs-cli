// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients/aws/cloudwatchlogs (interfaces: Client)

package mock_cloudwatchlogs

import (
	cloudwatchlogs "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) CreateLogGroup(_param0 *string) error {
	ret := _m.ctrl.Call(_m, "CreateLogGroup", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) CreateLogGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateLogGroup", arg0)
}

func (_m *MockClient) FilterAllLogEvents(_param0 *cloudwatchlogs.FilterLogEventsInput, _param1 func([]*cloudwatchlogs.FilteredLogEvent)) error {
	ret := _m.ctrl.Call(_m, "FilterAllLogEvents", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) FilterAllLogEvents(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilterAllLogEvents", arg0, arg1)
}
