// Code generated by MockGen. DO NOT EDIT.
// Source: routemonitor_supplement.go

// Package routemonitor is a generated GoMock package.
package routemonitor

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift/api/route/v1"
	v1alpha1 "github.com/openshift/route-monitor-operator/api/v1alpha1"
	reconcile "github.com/openshift/route-monitor-operator/pkg/util/reconcile"
	reflect "reflect"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

// MockRouteMonitorReconcilerInterface is a mock of RouteMonitorReconcilerInterface interface
type MockRouteMonitorReconcilerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRouteMonitorReconcilerInterfaceMockRecorder
}

// MockRouteMonitorReconcilerInterfaceMockRecorder is the mock recorder for MockRouteMonitorReconcilerInterface
type MockRouteMonitorReconcilerInterfaceMockRecorder struct {
	mock *MockRouteMonitorReconcilerInterface
}

// NewMockRouteMonitorReconcilerInterface creates a new mock instance
func NewMockRouteMonitorReconcilerInterface(ctrl *gomock.Controller) *MockRouteMonitorReconcilerInterface {
	mock := &MockRouteMonitorReconcilerInterface{ctrl: ctrl}
	mock.recorder = &MockRouteMonitorReconcilerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouteMonitorReconcilerInterface) EXPECT() *MockRouteMonitorReconcilerInterfaceMockRecorder {
	return m.recorder
}

// EnsureRouteMonitorAndDependenciesAbsent mocks base method
func (m *MockRouteMonitorReconcilerInterface) EnsureRouteMonitorAndDependenciesAbsent(ctx context.Context, routeMonitor v1alpha1.RouteMonitor) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureRouteMonitorAndDependenciesAbsent", ctx, routeMonitor)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureRouteMonitorAndDependenciesAbsent indicates an expected call of EnsureRouteMonitorAndDependenciesAbsent
func (mr *MockRouteMonitorReconcilerInterfaceMockRecorder) EnsureRouteMonitorAndDependenciesAbsent(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureRouteMonitorAndDependenciesAbsent", reflect.TypeOf((*MockRouteMonitorReconcilerInterface)(nil).EnsureRouteMonitorAndDependenciesAbsent), ctx, routeMonitor)
}

// EnsurePrometheusRuleResourceExists mocks base method
func (m *MockRouteMonitorReconcilerInterface) EnsurePrometheusRuleResourceExists(ctx context.Context, routeMonitor v1alpha1.RouteMonitor) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsurePrometheusRuleResourceExists", ctx, routeMonitor)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsurePrometheusRuleResourceExists indicates an expected call of EnsurePrometheusRuleResourceExists
func (mr *MockRouteMonitorReconcilerInterfaceMockRecorder) EnsurePrometheusRuleResourceExists(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsurePrometheusRuleResourceExists", reflect.TypeOf((*MockRouteMonitorReconcilerInterface)(nil).EnsurePrometheusRuleResourceExists), ctx, routeMonitor)
}

// EnsureServiceMonitorResourceExists mocks base method
func (m *MockRouteMonitorReconcilerInterface) EnsureServiceMonitorResourceExists(ctx context.Context, routeMonitor v1alpha1.RouteMonitor) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureServiceMonitorResourceExists", ctx, routeMonitor)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureServiceMonitorResourceExists indicates an expected call of EnsureServiceMonitorResourceExists
func (mr *MockRouteMonitorReconcilerInterfaceMockRecorder) EnsureServiceMonitorResourceExists(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureServiceMonitorResourceExists", reflect.TypeOf((*MockRouteMonitorReconcilerInterface)(nil).EnsureServiceMonitorResourceExists), ctx, routeMonitor)
}

// EnsureRouteURLExists mocks base method
func (m *MockRouteMonitorReconcilerInterface) EnsureRouteURLExists(ctx context.Context, route v1.Route, routeMonitor v1alpha1.RouteMonitor) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureRouteURLExists", ctx, route, routeMonitor)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureRouteURLExists indicates an expected call of EnsureRouteURLExists
func (mr *MockRouteMonitorReconcilerInterfaceMockRecorder) EnsureRouteURLExists(ctx, route, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureRouteURLExists", reflect.TypeOf((*MockRouteMonitorReconcilerInterface)(nil).EnsureRouteURLExists), ctx, route, routeMonitor)
}

// EnsureFinalizerSet mocks base method
func (m *MockRouteMonitorReconcilerInterface) EnsureFinalizerSet(ctx context.Context, routeMonitor v1alpha1.RouteMonitor) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureFinalizerSet", ctx, routeMonitor)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureFinalizerSet indicates an expected call of EnsureFinalizerSet
func (mr *MockRouteMonitorReconcilerInterfaceMockRecorder) EnsureFinalizerSet(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureFinalizerSet", reflect.TypeOf((*MockRouteMonitorReconcilerInterface)(nil).EnsureFinalizerSet), ctx, routeMonitor)
}

// EnsureFinalizerAbsent mocks base method
func (m *MockRouteMonitorReconcilerInterface) EnsureFinalizerAbsent(ctx context.Context, routeMonitor v1alpha1.RouteMonitor) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureFinalizerAbsent", ctx, routeMonitor)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureFinalizerAbsent indicates an expected call of EnsureFinalizerAbsent
func (mr *MockRouteMonitorReconcilerInterfaceMockRecorder) EnsureFinalizerAbsent(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureFinalizerAbsent", reflect.TypeOf((*MockRouteMonitorReconcilerInterface)(nil).EnsureFinalizerAbsent), ctx, routeMonitor)
}

// MockRouteMonitorSupplementInterface is a mock of RouteMonitorSupplementInterface interface
type MockRouteMonitorSupplementInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRouteMonitorSupplementInterfaceMockRecorder
}

// MockRouteMonitorSupplementInterfaceMockRecorder is the mock recorder for MockRouteMonitorSupplementInterface
type MockRouteMonitorSupplementInterfaceMockRecorder struct {
	mock *MockRouteMonitorSupplementInterface
}

// NewMockRouteMonitorSupplementInterface creates a new mock instance
func NewMockRouteMonitorSupplementInterface(ctrl *gomock.Controller) *MockRouteMonitorSupplementInterface {
	mock := &MockRouteMonitorSupplementInterface{ctrl: ctrl}
	mock.recorder = &MockRouteMonitorSupplementInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouteMonitorSupplementInterface) EXPECT() *MockRouteMonitorSupplementInterfaceMockRecorder {
	return m.recorder
}

// GetRoute mocks base method
func (m *MockRouteMonitorSupplementInterface) GetRoute(ctx context.Context, routeMonitor v1alpha1.RouteMonitor) (v1.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoute", ctx, routeMonitor)
	ret0, _ := ret[0].(v1.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoute indicates an expected call of GetRoute
func (mr *MockRouteMonitorSupplementInterfaceMockRecorder) GetRoute(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoute", reflect.TypeOf((*MockRouteMonitorSupplementInterface)(nil).GetRoute), ctx, routeMonitor)
}

// GetRouteMonitor mocks base method
func (m *MockRouteMonitorSupplementInterface) GetRouteMonitor(ctx context.Context, req controllerruntime.Request) (v1alpha1.RouteMonitor, reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouteMonitor", ctx, req)
	ret0, _ := ret[0].(v1alpha1.RouteMonitor)
	ret1, _ := ret[1].(reconcile.Result)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRouteMonitor indicates an expected call of GetRouteMonitor
func (mr *MockRouteMonitorSupplementInterfaceMockRecorder) GetRouteMonitor(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouteMonitor", reflect.TypeOf((*MockRouteMonitorSupplementInterface)(nil).GetRouteMonitor), ctx, req)
}

// UpdateRouteMonitor mocks base method
func (m *MockRouteMonitorSupplementInterface) UpdateRouteMonitor(ctx context.Context, routeMonitor *v1alpha1.RouteMonitor) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRouteMonitor", ctx, routeMonitor)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRouteMonitor indicates an expected call of UpdateRouteMonitor
func (mr *MockRouteMonitorSupplementInterfaceMockRecorder) UpdateRouteMonitor(ctx, routeMonitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRouteMonitor", reflect.TypeOf((*MockRouteMonitorSupplementInterface)(nil).UpdateRouteMonitor), ctx, routeMonitor)
}
