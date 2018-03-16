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
// Source: github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients/aws/cloudwatchlogs (interfaces: LogClientFactory)

package mock_cloudwatchlogs

import (
	cloudwatchlogs "github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients/aws/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
)

// Mock of LogClientFactory interface
type MockLogClientFactory struct {
	ctrl     *gomock.Controller
	recorder *_MockLogClientFactoryRecorder
}

// Recorder for MockLogClientFactory (not exported)
type _MockLogClientFactoryRecorder struct {
	mock *MockLogClientFactory
}

func NewMockLogClientFactory(ctrl *gomock.Controller) *MockLogClientFactory {
	mock := &MockLogClientFactory{ctrl: ctrl}
	mock.recorder = &_MockLogClientFactoryRecorder{mock}
	return mock
}

func (_m *MockLogClientFactory) EXPECT() *_MockLogClientFactoryRecorder {
	return _m.recorder
}

func (_m *MockLogClientFactory) Get(_param0 string) cloudwatchlogs.Client {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(cloudwatchlogs.Client)
	return ret0
}

func (_mr *_MockLogClientFactoryRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}
