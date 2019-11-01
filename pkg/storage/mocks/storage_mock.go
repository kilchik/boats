package storage

// DO NOT EDIT!
// The code below was generated with http://github.com/gojuno/minimock (dev)

//go:generate minimock -i boats/pkg/storage.Storage -o ./storage_mock.go

import (
	"sync/atomic"
	"time"

	"boats/clients/nausys"
	"boats/pkg/storage"

	"context"

	"github.com/jmoiron/sqlx"

	"github.com/gojuno/minimock"
)

// StorageMock implements storage.Storage
type StorageMock struct {
	t minimock.Tester

	funcClearAll          func(ctx context.Context, querier sqlx.ExecerContext) (err error)
	afterClearAllCounter  uint64
	beforeClearAllCounter uint64
	ClearAllMock          mStorageMockClearAll

	funcFindBuildersByPrefix          func(ctx context.Context, prefix string, limit int) (sa1 []string, err error)
	afterFindBuildersByPrefixCounter  uint64
	beforeFindBuildersByPrefixCounter uint64
	FindBuildersByPrefixMock          mStorageMockFindBuildersByPrefix

	funcFindModelsByPrefix          func(ctx context.Context, prefix string, limit int) (sa1 []string, err error)
	afterFindModelsByPrefixCounter  uint64
	beforeFindModelsByPrefixCounter uint64
	FindModelsByPrefixMock          mStorageMockFindModelsByPrefix

	funcFindYachts          func(ctx context.Context, builderNamePrefix string, modelNamePrefix string, limit int, offset int) (yachts []storage.YachtInfo, total int64, err error)
	afterFindYachtsCounter  uint64
	beforeFindYachtsCounter uint64
	FindYachtsMock          mStorageMockFindYachts

	funcGetLastUpdateInfo          func(ctx context.Context) (t1 time.Time, err error)
	afterGetLastUpdateInfoCounter  uint64
	beforeGetLastUpdateInfoCounter uint64
	GetLastUpdateInfoMock          mStorageMockGetLastUpdateInfo

	funcInsertBuilders          func(ctx context.Context, querier sqlx.ExecerContext, builders []nausys.Builder) (err error)
	afterInsertBuildersCounter  uint64
	beforeInsertBuildersCounter uint64
	InsertBuildersMock          mStorageMockInsertBuilders

	funcInsertCharters          func(ctx context.Context, querier sqlx.ExecerContext, charters []nausys.Charter) (err error)
	afterInsertChartersCounter  uint64
	beforeInsertChartersCounter uint64
	InsertChartersMock          mStorageMockInsertCharters

	funcInsertModels          func(ctx context.Context, querier sqlx.ExecerContext, models []nausys.Model) (err error)
	afterInsertModelsCounter  uint64
	beforeInsertModelsCounter uint64
	InsertModelsMock          mStorageMockInsertModels

	funcInsertUpdateInfo          func(ctx context.Context, querier sqlx.ExecerContext) (err error)
	afterInsertUpdateInfoCounter  uint64
	beforeInsertUpdateInfoCounter uint64
	InsertUpdateInfoMock          mStorageMockInsertUpdateInfo

	funcInsertYachts          func(ctx context.Context, querier sqlx.ExecerContext, yachts []*nausys.Yacht) (err error)
	afterInsertYachtsCounter  uint64
	beforeInsertYachtsCounter uint64
	InsertYachtsMock          mStorageMockInsertYachts

	funcWithTransaction          func(ctx context.Context, f func(tx *sqlx.Tx) error) (err error)
	afterWithTransactionCounter  uint64
	beforeWithTransactionCounter uint64
	WithTransactionMock          mStorageMockWithTransaction
}

// NewStorageMock returns a mock for storage.Storage
func NewStorageMock(t minimock.Tester) *StorageMock {
	m := &StorageMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}
	m.ClearAllMock = mStorageMockClearAll{mock: m}
	m.FindBuildersByPrefixMock = mStorageMockFindBuildersByPrefix{mock: m}
	m.FindModelsByPrefixMock = mStorageMockFindModelsByPrefix{mock: m}
	m.FindYachtsMock = mStorageMockFindYachts{mock: m}
	m.GetLastUpdateInfoMock = mStorageMockGetLastUpdateInfo{mock: m}
	m.InsertBuildersMock = mStorageMockInsertBuilders{mock: m}
	m.InsertChartersMock = mStorageMockInsertCharters{mock: m}
	m.InsertModelsMock = mStorageMockInsertModels{mock: m}
	m.InsertUpdateInfoMock = mStorageMockInsertUpdateInfo{mock: m}
	m.InsertYachtsMock = mStorageMockInsertYachts{mock: m}
	m.WithTransactionMock = mStorageMockWithTransaction{mock: m}

	return m
}

type mStorageMockClearAll struct {
	mock               *StorageMock
	defaultExpectation *StorageMockClearAllExpectation
	expectations       []*StorageMockClearAllExpectation
}

// StorageMockClearAllExpectation specifies expectation struct of the Storage.ClearAll
type StorageMockClearAllExpectation struct {
	mock    *StorageMock
	params  *StorageMockClearAllParams
	results *StorageMockClearAllResults
	Counter uint64
}

// StorageMockClearAllParams contains parameters of the Storage.ClearAll
type StorageMockClearAllParams struct {
	ctx     context.Context
	querier sqlx.ExecerContext
}

// StorageMockClearAllResults contains results of the Storage.ClearAll
type StorageMockClearAllResults struct {
	err error
}

// Expect sets up expected params for Storage.ClearAll
func (m *mStorageMockClearAll) Expect(ctx context.Context, querier sqlx.ExecerContext) *mStorageMockClearAll {
	if m.mock.funcClearAll != nil {
		m.mock.t.Fatalf("StorageMock.ClearAll mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockClearAllExpectation{}
	}

	m.defaultExpectation.params = &StorageMockClearAllParams{ctx, querier}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Storage.ClearAll
func (m *mStorageMockClearAll) Return(err error) *StorageMock {
	if m.mock.funcClearAll != nil {
		m.mock.t.Fatalf("StorageMock.ClearAll mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockClearAllExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StorageMockClearAllResults{err}
	return m.mock
}

//Set uses given function f to mock the Storage.ClearAll method
func (m *mStorageMockClearAll) Set(f func(ctx context.Context, querier sqlx.ExecerContext) (err error)) *StorageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Storage.ClearAll method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Storage.ClearAll method")
	}

	m.mock.funcClearAll = f
	return m.mock
}

// When sets expectation for the Storage.ClearAll which will trigger the result defined by the following
// Then helper
func (m *mStorageMockClearAll) When(ctx context.Context, querier sqlx.ExecerContext) *StorageMockClearAllExpectation {
	if m.mock.funcClearAll != nil {
		m.mock.t.Fatalf("StorageMock.ClearAll mock is already set by Set")
	}

	expectation := &StorageMockClearAllExpectation{
		mock:   m.mock,
		params: &StorageMockClearAllParams{ctx, querier},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Storage.ClearAll return parameters for the expectation previously defined by the When method
func (e *StorageMockClearAllExpectation) Then(err error) *StorageMock {
	e.results = &StorageMockClearAllResults{err}
	return e.mock
}

// ClearAll implements storage.Storage
func (m *StorageMock) ClearAll(ctx context.Context, querier sqlx.ExecerContext) (err error) {
	atomic.AddUint64(&m.beforeClearAllCounter, 1)
	defer atomic.AddUint64(&m.afterClearAllCounter, 1)

	for _, e := range m.ClearAllMock.expectations {
		if minimock.Equal(*e.params, StorageMockClearAllParams{ctx, querier}) {
			atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.ClearAllMock.defaultExpectation != nil {
		atomic.AddUint64(&m.ClearAllMock.defaultExpectation.Counter, 1)
		want := m.ClearAllMock.defaultExpectation.params
		got := StorageMockClearAllParams{ctx, querier}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StorageMock.ClearAll got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.ClearAllMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StorageMock.ClearAll")
		}
		return (*results).err
	}
	if m.funcClearAll != nil {
		return m.funcClearAll(ctx, querier)
	}
	m.t.Fatalf("Unexpected call to StorageMock.ClearAll. %v %v", ctx, querier)
	return
}

// ClearAllAfterCounter returns a count of finished StorageMock.ClearAll invocations
func (m *StorageMock) ClearAllAfterCounter() uint64 {
	return atomic.LoadUint64(&m.afterClearAllCounter)
}

// ClearAllBeforeCounter returns a count of StorageMock.ClearAll invocations
func (m *StorageMock) ClearAllBeforeCounter() uint64 {
	return atomic.LoadUint64(&m.beforeClearAllCounter)
}

// MinimockClearAllDone returns true if the count of the ClearAll invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockClearAllDone() bool {
	for _, e := range m.ClearAllMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ClearAllMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterClearAllCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClearAll != nil && atomic.LoadUint64(&m.afterClearAllCounter) < 1 {
		return false
	}
	return true
}

// MinimockClearAllInspect logs each unmet expectation
func (m *StorageMock) MinimockClearAllInspect() {
	for _, e := range m.ClearAllMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.ClearAll with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ClearAllMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterClearAllCounter) < 1 {
		m.t.Errorf("Expected call to StorageMock.ClearAll with params: %#v", *m.ClearAllMock.defaultExpectation.params)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClearAll != nil && atomic.LoadUint64(&m.afterClearAllCounter) < 1 {
		m.t.Error("Expected call to StorageMock.ClearAll")
	}
}

type mStorageMockFindBuildersByPrefix struct {
	mock               *StorageMock
	defaultExpectation *StorageMockFindBuildersByPrefixExpectation
	expectations       []*StorageMockFindBuildersByPrefixExpectation
}

// StorageMockFindBuildersByPrefixExpectation specifies expectation struct of the Storage.FindBuildersByPrefix
type StorageMockFindBuildersByPrefixExpectation struct {
	mock    *StorageMock
	params  *StorageMockFindBuildersByPrefixParams
	results *StorageMockFindBuildersByPrefixResults
	Counter uint64
}

// StorageMockFindBuildersByPrefixParams contains parameters of the Storage.FindBuildersByPrefix
type StorageMockFindBuildersByPrefixParams struct {
	ctx    context.Context
	prefix string
	limit  int
}

// StorageMockFindBuildersByPrefixResults contains results of the Storage.FindBuildersByPrefix
type StorageMockFindBuildersByPrefixResults struct {
	sa1 []string
	err error
}

// Expect sets up expected params for Storage.FindBuildersByPrefix
func (m *mStorageMockFindBuildersByPrefix) Expect(ctx context.Context, prefix string, limit int) *mStorageMockFindBuildersByPrefix {
	if m.mock.funcFindBuildersByPrefix != nil {
		m.mock.t.Fatalf("StorageMock.FindBuildersByPrefix mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockFindBuildersByPrefixExpectation{}
	}

	m.defaultExpectation.params = &StorageMockFindBuildersByPrefixParams{ctx, prefix, limit}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Storage.FindBuildersByPrefix
func (m *mStorageMockFindBuildersByPrefix) Return(sa1 []string, err error) *StorageMock {
	if m.mock.funcFindBuildersByPrefix != nil {
		m.mock.t.Fatalf("StorageMock.FindBuildersByPrefix mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockFindBuildersByPrefixExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StorageMockFindBuildersByPrefixResults{sa1, err}
	return m.mock
}

//Set uses given function f to mock the Storage.FindBuildersByPrefix method
func (m *mStorageMockFindBuildersByPrefix) Set(f func(ctx context.Context, prefix string, limit int) (sa1 []string, err error)) *StorageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Storage.FindBuildersByPrefix method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Storage.FindBuildersByPrefix method")
	}

	m.mock.funcFindBuildersByPrefix = f
	return m.mock
}

// When sets expectation for the Storage.FindBuildersByPrefix which will trigger the result defined by the following
// Then helper
func (m *mStorageMockFindBuildersByPrefix) When(ctx context.Context, prefix string, limit int) *StorageMockFindBuildersByPrefixExpectation {
	if m.mock.funcFindBuildersByPrefix != nil {
		m.mock.t.Fatalf("StorageMock.FindBuildersByPrefix mock is already set by Set")
	}

	expectation := &StorageMockFindBuildersByPrefixExpectation{
		mock:   m.mock,
		params: &StorageMockFindBuildersByPrefixParams{ctx, prefix, limit},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Storage.FindBuildersByPrefix return parameters for the expectation previously defined by the When method
func (e *StorageMockFindBuildersByPrefixExpectation) Then(sa1 []string, err error) *StorageMock {
	e.results = &StorageMockFindBuildersByPrefixResults{sa1, err}
	return e.mock
}

// FindBuildersByPrefix implements storage.Storage
func (m *StorageMock) FindBuildersByPrefix(ctx context.Context, prefix string, limit int) (sa1 []string, err error) {
	atomic.AddUint64(&m.beforeFindBuildersByPrefixCounter, 1)
	defer atomic.AddUint64(&m.afterFindBuildersByPrefixCounter, 1)

	for _, e := range m.FindBuildersByPrefixMock.expectations {
		if minimock.Equal(*e.params, StorageMockFindBuildersByPrefixParams{ctx, prefix, limit}) {
			atomic.AddUint64(&e.Counter, 1)
			return e.results.sa1, e.results.err
		}
	}

	if m.FindBuildersByPrefixMock.defaultExpectation != nil {
		atomic.AddUint64(&m.FindBuildersByPrefixMock.defaultExpectation.Counter, 1)
		want := m.FindBuildersByPrefixMock.defaultExpectation.params
		got := StorageMockFindBuildersByPrefixParams{ctx, prefix, limit}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StorageMock.FindBuildersByPrefix got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.FindBuildersByPrefixMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StorageMock.FindBuildersByPrefix")
		}
		return (*results).sa1, (*results).err
	}
	if m.funcFindBuildersByPrefix != nil {
		return m.funcFindBuildersByPrefix(ctx, prefix, limit)
	}
	m.t.Fatalf("Unexpected call to StorageMock.FindBuildersByPrefix. %v %v %v", ctx, prefix, limit)
	return
}

// FindBuildersByPrefixAfterCounter returns a count of finished StorageMock.FindBuildersByPrefix invocations
func (m *StorageMock) FindBuildersByPrefixAfterCounter() uint64 {
	return atomic.LoadUint64(&m.afterFindBuildersByPrefixCounter)
}

// FindBuildersByPrefixBeforeCounter returns a count of StorageMock.FindBuildersByPrefix invocations
func (m *StorageMock) FindBuildersByPrefixBeforeCounter() uint64 {
	return atomic.LoadUint64(&m.beforeFindBuildersByPrefixCounter)
}

// MinimockFindBuildersByPrefixDone returns true if the count of the FindBuildersByPrefix invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockFindBuildersByPrefixDone() bool {
	for _, e := range m.FindBuildersByPrefixMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindBuildersByPrefixMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterFindBuildersByPrefixCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindBuildersByPrefix != nil && atomic.LoadUint64(&m.afterFindBuildersByPrefixCounter) < 1 {
		return false
	}
	return true
}

// MinimockFindBuildersByPrefixInspect logs each unmet expectation
func (m *StorageMock) MinimockFindBuildersByPrefixInspect() {
	for _, e := range m.FindBuildersByPrefixMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.FindBuildersByPrefix with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindBuildersByPrefixMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterFindBuildersByPrefixCounter) < 1 {
		m.t.Errorf("Expected call to StorageMock.FindBuildersByPrefix with params: %#v", *m.FindBuildersByPrefixMock.defaultExpectation.params)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindBuildersByPrefix != nil && atomic.LoadUint64(&m.afterFindBuildersByPrefixCounter) < 1 {
		m.t.Error("Expected call to StorageMock.FindBuildersByPrefix")
	}
}

type mStorageMockFindModelsByPrefix struct {
	mock               *StorageMock
	defaultExpectation *StorageMockFindModelsByPrefixExpectation
	expectations       []*StorageMockFindModelsByPrefixExpectation
}

// StorageMockFindModelsByPrefixExpectation specifies expectation struct of the Storage.FindModelsByPrefix
type StorageMockFindModelsByPrefixExpectation struct {
	mock    *StorageMock
	params  *StorageMockFindModelsByPrefixParams
	results *StorageMockFindModelsByPrefixResults
	Counter uint64
}

// StorageMockFindModelsByPrefixParams contains parameters of the Storage.FindModelsByPrefix
type StorageMockFindModelsByPrefixParams struct {
	ctx    context.Context
	prefix string
	limit  int
}

// StorageMockFindModelsByPrefixResults contains results of the Storage.FindModelsByPrefix
type StorageMockFindModelsByPrefixResults struct {
	sa1 []string
	err error
}

// Expect sets up expected params for Storage.FindModelsByPrefix
func (m *mStorageMockFindModelsByPrefix) Expect(ctx context.Context, prefix string, limit int) *mStorageMockFindModelsByPrefix {
	if m.mock.funcFindModelsByPrefix != nil {
		m.mock.t.Fatalf("StorageMock.FindModelsByPrefix mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockFindModelsByPrefixExpectation{}
	}

	m.defaultExpectation.params = &StorageMockFindModelsByPrefixParams{ctx, prefix, limit}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Storage.FindModelsByPrefix
func (m *mStorageMockFindModelsByPrefix) Return(sa1 []string, err error) *StorageMock {
	if m.mock.funcFindModelsByPrefix != nil {
		m.mock.t.Fatalf("StorageMock.FindModelsByPrefix mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockFindModelsByPrefixExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StorageMockFindModelsByPrefixResults{sa1, err}
	return m.mock
}

//Set uses given function f to mock the Storage.FindModelsByPrefix method
func (m *mStorageMockFindModelsByPrefix) Set(f func(ctx context.Context, prefix string, limit int) (sa1 []string, err error)) *StorageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Storage.FindModelsByPrefix method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Storage.FindModelsByPrefix method")
	}

	m.mock.funcFindModelsByPrefix = f
	return m.mock
}

// When sets expectation for the Storage.FindModelsByPrefix which will trigger the result defined by the following
// Then helper
func (m *mStorageMockFindModelsByPrefix) When(ctx context.Context, prefix string, limit int) *StorageMockFindModelsByPrefixExpectation {
	if m.mock.funcFindModelsByPrefix != nil {
		m.mock.t.Fatalf("StorageMock.FindModelsByPrefix mock is already set by Set")
	}

	expectation := &StorageMockFindModelsByPrefixExpectation{
		mock:   m.mock,
		params: &StorageMockFindModelsByPrefixParams{ctx, prefix, limit},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Storage.FindModelsByPrefix return parameters for the expectation previously defined by the When method
func (e *StorageMockFindModelsByPrefixExpectation) Then(sa1 []string, err error) *StorageMock {
	e.results = &StorageMockFindModelsByPrefixResults{sa1, err}
	return e.mock
}

// FindModelsByPrefix implements storage.Storage
func (m *StorageMock) FindModelsByPrefix(ctx context.Context, prefix string, limit int) (sa1 []string, err error) {
	atomic.AddUint64(&m.beforeFindModelsByPrefixCounter, 1)
	defer atomic.AddUint64(&m.afterFindModelsByPrefixCounter, 1)

	for _, e := range m.FindModelsByPrefixMock.expectations {
		if minimock.Equal(*e.params, StorageMockFindModelsByPrefixParams{ctx, prefix, limit}) {
			atomic.AddUint64(&e.Counter, 1)
			return e.results.sa1, e.results.err
		}
	}

	if m.FindModelsByPrefixMock.defaultExpectation != nil {
		atomic.AddUint64(&m.FindModelsByPrefixMock.defaultExpectation.Counter, 1)
		want := m.FindModelsByPrefixMock.defaultExpectation.params
		got := StorageMockFindModelsByPrefixParams{ctx, prefix, limit}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StorageMock.FindModelsByPrefix got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.FindModelsByPrefixMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StorageMock.FindModelsByPrefix")
		}
		return (*results).sa1, (*results).err
	}
	if m.funcFindModelsByPrefix != nil {
		return m.funcFindModelsByPrefix(ctx, prefix, limit)
	}
	m.t.Fatalf("Unexpected call to StorageMock.FindModelsByPrefix. %v %v %v", ctx, prefix, limit)
	return
}

// FindModelsByPrefixAfterCounter returns a count of finished StorageMock.FindModelsByPrefix invocations
func (m *StorageMock) FindModelsByPrefixAfterCounter() uint64 {
	return atomic.LoadUint64(&m.afterFindModelsByPrefixCounter)
}

// FindModelsByPrefixBeforeCounter returns a count of StorageMock.FindModelsByPrefix invocations
func (m *StorageMock) FindModelsByPrefixBeforeCounter() uint64 {
	return atomic.LoadUint64(&m.beforeFindModelsByPrefixCounter)
}

// MinimockFindModelsByPrefixDone returns true if the count of the FindModelsByPrefix invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockFindModelsByPrefixDone() bool {
	for _, e := range m.FindModelsByPrefixMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindModelsByPrefixMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterFindModelsByPrefixCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindModelsByPrefix != nil && atomic.LoadUint64(&m.afterFindModelsByPrefixCounter) < 1 {
		return false
	}
	return true
}

// MinimockFindModelsByPrefixInspect logs each unmet expectation
func (m *StorageMock) MinimockFindModelsByPrefixInspect() {
	for _, e := range m.FindModelsByPrefixMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.FindModelsByPrefix with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindModelsByPrefixMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterFindModelsByPrefixCounter) < 1 {
		m.t.Errorf("Expected call to StorageMock.FindModelsByPrefix with params: %#v", *m.FindModelsByPrefixMock.defaultExpectation.params)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindModelsByPrefix != nil && atomic.LoadUint64(&m.afterFindModelsByPrefixCounter) < 1 {
		m.t.Error("Expected call to StorageMock.FindModelsByPrefix")
	}
}

type mStorageMockFindYachts struct {
	mock               *StorageMock
	defaultExpectation *StorageMockFindYachtsExpectation
	expectations       []*StorageMockFindYachtsExpectation
}

// StorageMockFindYachtsExpectation specifies expectation struct of the Storage.FindYachts
type StorageMockFindYachtsExpectation struct {
	mock    *StorageMock
	params  *StorageMockFindYachtsParams
	results *StorageMockFindYachtsResults
	Counter uint64
}

// StorageMockFindYachtsParams contains parameters of the Storage.FindYachts
type StorageMockFindYachtsParams struct {
	ctx               context.Context
	builderNamePrefix string
	modelNamePrefix   string
	limit             int
	offset            int
}

// StorageMockFindYachtsResults contains results of the Storage.FindYachts
type StorageMockFindYachtsResults struct {
	yachts []storage.YachtInfo
	total  int64
	err    error
}

// Expect sets up expected params for Storage.FindYachts
func (m *mStorageMockFindYachts) Expect(ctx context.Context, builderNamePrefix string, modelNamePrefix string, limit int, offset int) *mStorageMockFindYachts {
	if m.mock.funcFindYachts != nil {
		m.mock.t.Fatalf("StorageMock.FindYachts mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockFindYachtsExpectation{}
	}

	m.defaultExpectation.params = &StorageMockFindYachtsParams{ctx, builderNamePrefix, modelNamePrefix, limit, offset}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Storage.FindYachts
func (m *mStorageMockFindYachts) Return(yachts []storage.YachtInfo, total int64, err error) *StorageMock {
	if m.mock.funcFindYachts != nil {
		m.mock.t.Fatalf("StorageMock.FindYachts mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockFindYachtsExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StorageMockFindYachtsResults{yachts, total, err}
	return m.mock
}

//Set uses given function f to mock the Storage.FindYachts method
func (m *mStorageMockFindYachts) Set(f func(ctx context.Context, builderNamePrefix string, modelNamePrefix string, limit int, offset int) (yachts []storage.YachtInfo, total int64, err error)) *StorageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Storage.FindYachts method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Storage.FindYachts method")
	}

	m.mock.funcFindYachts = f
	return m.mock
}

// When sets expectation for the Storage.FindYachts which will trigger the result defined by the following
// Then helper
func (m *mStorageMockFindYachts) When(ctx context.Context, builderNamePrefix string, modelNamePrefix string, limit int, offset int) *StorageMockFindYachtsExpectation {
	if m.mock.funcFindYachts != nil {
		m.mock.t.Fatalf("StorageMock.FindYachts mock is already set by Set")
	}

	expectation := &StorageMockFindYachtsExpectation{
		mock:   m.mock,
		params: &StorageMockFindYachtsParams{ctx, builderNamePrefix, modelNamePrefix, limit, offset},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Storage.FindYachts return parameters for the expectation previously defined by the When method
func (e *StorageMockFindYachtsExpectation) Then(yachts []storage.YachtInfo, total int64, err error) *StorageMock {
	e.results = &StorageMockFindYachtsResults{yachts, total, err}
	return e.mock
}

// FindYachts implements storage.Storage
func (m *StorageMock) FindYachts(ctx context.Context, builderNamePrefix string, modelNamePrefix string, limit int, offset int) (yachts []storage.YachtInfo, total int64, err error) {
	atomic.AddUint64(&m.beforeFindYachtsCounter, 1)
	defer atomic.AddUint64(&m.afterFindYachtsCounter, 1)

	for _, e := range m.FindYachtsMock.expectations {
		if minimock.Equal(*e.params, StorageMockFindYachtsParams{ctx, builderNamePrefix, modelNamePrefix, limit, offset}) {
			atomic.AddUint64(&e.Counter, 1)
			return e.results.yachts, e.results.total, e.results.err
		}
	}

	if m.FindYachtsMock.defaultExpectation != nil {
		atomic.AddUint64(&m.FindYachtsMock.defaultExpectation.Counter, 1)
		want := m.FindYachtsMock.defaultExpectation.params
		got := StorageMockFindYachtsParams{ctx, builderNamePrefix, modelNamePrefix, limit, offset}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StorageMock.FindYachts got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.FindYachtsMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StorageMock.FindYachts")
		}
		return (*results).yachts, (*results).total, (*results).err
	}
	if m.funcFindYachts != nil {
		return m.funcFindYachts(ctx, builderNamePrefix, modelNamePrefix, limit, offset)
	}
	m.t.Fatalf("Unexpected call to StorageMock.FindYachts. %v %v %v %v %v", ctx, builderNamePrefix, modelNamePrefix, limit, offset)
	return
}

// FindYachtsAfterCounter returns a count of finished StorageMock.FindYachts invocations
func (m *StorageMock) FindYachtsAfterCounter() uint64 {
	return atomic.LoadUint64(&m.afterFindYachtsCounter)
}

// FindYachtsBeforeCounter returns a count of StorageMock.FindYachts invocations
func (m *StorageMock) FindYachtsBeforeCounter() uint64 {
	return atomic.LoadUint64(&m.beforeFindYachtsCounter)
}

// MinimockFindYachtsDone returns true if the count of the FindYachts invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockFindYachtsDone() bool {
	for _, e := range m.FindYachtsMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindYachtsMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterFindYachtsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindYachts != nil && atomic.LoadUint64(&m.afterFindYachtsCounter) < 1 {
		return false
	}
	return true
}

// MinimockFindYachtsInspect logs each unmet expectation
func (m *StorageMock) MinimockFindYachtsInspect() {
	for _, e := range m.FindYachtsMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.FindYachts with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindYachtsMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterFindYachtsCounter) < 1 {
		m.t.Errorf("Expected call to StorageMock.FindYachts with params: %#v", *m.FindYachtsMock.defaultExpectation.params)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindYachts != nil && atomic.LoadUint64(&m.afterFindYachtsCounter) < 1 {
		m.t.Error("Expected call to StorageMock.FindYachts")
	}
}

type mStorageMockGetLastUpdateInfo struct {
	mock               *StorageMock
	defaultExpectation *StorageMockGetLastUpdateInfoExpectation
	expectations       []*StorageMockGetLastUpdateInfoExpectation
}

// StorageMockGetLastUpdateInfoExpectation specifies expectation struct of the Storage.GetLastUpdateInfo
type StorageMockGetLastUpdateInfoExpectation struct {
	mock    *StorageMock
	params  *StorageMockGetLastUpdateInfoParams
	results *StorageMockGetLastUpdateInfoResults
	Counter uint64
}

// StorageMockGetLastUpdateInfoParams contains parameters of the Storage.GetLastUpdateInfo
type StorageMockGetLastUpdateInfoParams struct {
	ctx context.Context
}

// StorageMockGetLastUpdateInfoResults contains results of the Storage.GetLastUpdateInfo
type StorageMockGetLastUpdateInfoResults struct {
	t1  time.Time
	err error
}

// Expect sets up expected params for Storage.GetLastUpdateInfo
func (m *mStorageMockGetLastUpdateInfo) Expect(ctx context.Context) *mStorageMockGetLastUpdateInfo {
	if m.mock.funcGetLastUpdateInfo != nil {
		m.mock.t.Fatalf("StorageMock.GetLastUpdateInfo mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockGetLastUpdateInfoExpectation{}
	}

	m.defaultExpectation.params = &StorageMockGetLastUpdateInfoParams{ctx}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Storage.GetLastUpdateInfo
func (m *mStorageMockGetLastUpdateInfo) Return(t1 time.Time, err error) *StorageMock {
	if m.mock.funcGetLastUpdateInfo != nil {
		m.mock.t.Fatalf("StorageMock.GetLastUpdateInfo mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockGetLastUpdateInfoExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StorageMockGetLastUpdateInfoResults{t1, err}
	return m.mock
}

//Set uses given function f to mock the Storage.GetLastUpdateInfo method
func (m *mStorageMockGetLastUpdateInfo) Set(f func(ctx context.Context) (t1 time.Time, err error)) *StorageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Storage.GetLastUpdateInfo method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Storage.GetLastUpdateInfo method")
	}

	m.mock.funcGetLastUpdateInfo = f
	return m.mock
}

// When sets expectation for the Storage.GetLastUpdateInfo which will trigger the result defined by the following
// Then helper
func (m *mStorageMockGetLastUpdateInfo) When(ctx context.Context) *StorageMockGetLastUpdateInfoExpectation {
	if m.mock.funcGetLastUpdateInfo != nil {
		m.mock.t.Fatalf("StorageMock.GetLastUpdateInfo mock is already set by Set")
	}

	expectation := &StorageMockGetLastUpdateInfoExpectation{
		mock:   m.mock,
		params: &StorageMockGetLastUpdateInfoParams{ctx},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Storage.GetLastUpdateInfo return parameters for the expectation previously defined by the When method
func (e *StorageMockGetLastUpdateInfoExpectation) Then(t1 time.Time, err error) *StorageMock {
	e.results = &StorageMockGetLastUpdateInfoResults{t1, err}
	return e.mock
}

// GetLastUpdateInfo implements storage.Storage
func (m *StorageMock) GetLastUpdateInfo(ctx context.Context) (t1 time.Time, err error) {
	atomic.AddUint64(&m.beforeGetLastUpdateInfoCounter, 1)
	defer atomic.AddUint64(&m.afterGetLastUpdateInfoCounter, 1)

	for _, e := range m.GetLastUpdateInfoMock.expectations {
		if minimock.Equal(*e.params, StorageMockGetLastUpdateInfoParams{ctx}) {
			atomic.AddUint64(&e.Counter, 1)
			return e.results.t1, e.results.err
		}
	}

	if m.GetLastUpdateInfoMock.defaultExpectation != nil {
		atomic.AddUint64(&m.GetLastUpdateInfoMock.defaultExpectation.Counter, 1)
		want := m.GetLastUpdateInfoMock.defaultExpectation.params
		got := StorageMockGetLastUpdateInfoParams{ctx}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StorageMock.GetLastUpdateInfo got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.GetLastUpdateInfoMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StorageMock.GetLastUpdateInfo")
		}
		return (*results).t1, (*results).err
	}
	if m.funcGetLastUpdateInfo != nil {
		return m.funcGetLastUpdateInfo(ctx)
	}
	m.t.Fatalf("Unexpected call to StorageMock.GetLastUpdateInfo. %v", ctx)
	return
}

// GetLastUpdateInfoAfterCounter returns a count of finished StorageMock.GetLastUpdateInfo invocations
func (m *StorageMock) GetLastUpdateInfoAfterCounter() uint64 {
	return atomic.LoadUint64(&m.afterGetLastUpdateInfoCounter)
}

// GetLastUpdateInfoBeforeCounter returns a count of StorageMock.GetLastUpdateInfo invocations
func (m *StorageMock) GetLastUpdateInfoBeforeCounter() uint64 {
	return atomic.LoadUint64(&m.beforeGetLastUpdateInfoCounter)
}

// MinimockGetLastUpdateInfoDone returns true if the count of the GetLastUpdateInfo invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockGetLastUpdateInfoDone() bool {
	for _, e := range m.GetLastUpdateInfoMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetLastUpdateInfoMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterGetLastUpdateInfoCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetLastUpdateInfo != nil && atomic.LoadUint64(&m.afterGetLastUpdateInfoCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetLastUpdateInfoInspect logs each unmet expectation
func (m *StorageMock) MinimockGetLastUpdateInfoInspect() {
	for _, e := range m.GetLastUpdateInfoMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.GetLastUpdateInfo with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetLastUpdateInfoMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterGetLastUpdateInfoCounter) < 1 {
		m.t.Errorf("Expected call to StorageMock.GetLastUpdateInfo with params: %#v", *m.GetLastUpdateInfoMock.defaultExpectation.params)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetLastUpdateInfo != nil && atomic.LoadUint64(&m.afterGetLastUpdateInfoCounter) < 1 {
		m.t.Error("Expected call to StorageMock.GetLastUpdateInfo")
	}
}

type mStorageMockInsertBuilders struct {
	mock               *StorageMock
	defaultExpectation *StorageMockInsertBuildersExpectation
	expectations       []*StorageMockInsertBuildersExpectation
}

// StorageMockInsertBuildersExpectation specifies expectation struct of the Storage.InsertBuilders
type StorageMockInsertBuildersExpectation struct {
	mock    *StorageMock
	params  *StorageMockInsertBuildersParams
	results *StorageMockInsertBuildersResults
	Counter uint64
}

// StorageMockInsertBuildersParams contains parameters of the Storage.InsertBuilders
type StorageMockInsertBuildersParams struct {
	ctx      context.Context
	querier  sqlx.ExecerContext
	builders []nausys.Builder
}

// StorageMockInsertBuildersResults contains results of the Storage.InsertBuilders
type StorageMockInsertBuildersResults struct {
	err error
}

// Expect sets up expected params for Storage.InsertBuilders
func (m *mStorageMockInsertBuilders) Expect(ctx context.Context, querier sqlx.ExecerContext, builders []nausys.Builder) *mStorageMockInsertBuilders {
	if m.mock.funcInsertBuilders != nil {
		m.mock.t.Fatalf("StorageMock.InsertBuilders mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockInsertBuildersExpectation{}
	}

	m.defaultExpectation.params = &StorageMockInsertBuildersParams{ctx, querier, builders}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Storage.InsertBuilders
func (m *mStorageMockInsertBuilders) Return(err error) *StorageMock {
	if m.mock.funcInsertBuilders != nil {
		m.mock.t.Fatalf("StorageMock.InsertBuilders mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockInsertBuildersExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StorageMockInsertBuildersResults{err}
	return m.mock
}

//Set uses given function f to mock the Storage.InsertBuilders method
func (m *mStorageMockInsertBuilders) Set(f func(ctx context.Context, querier sqlx.ExecerContext, builders []nausys.Builder) (err error)) *StorageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Storage.InsertBuilders method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Storage.InsertBuilders method")
	}

	m.mock.funcInsertBuilders = f
	return m.mock
}

// When sets expectation for the Storage.InsertBuilders which will trigger the result defined by the following
// Then helper
func (m *mStorageMockInsertBuilders) When(ctx context.Context, querier sqlx.ExecerContext, builders []nausys.Builder) *StorageMockInsertBuildersExpectation {
	if m.mock.funcInsertBuilders != nil {
		m.mock.t.Fatalf("StorageMock.InsertBuilders mock is already set by Set")
	}

	expectation := &StorageMockInsertBuildersExpectation{
		mock:   m.mock,
		params: &StorageMockInsertBuildersParams{ctx, querier, builders},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Storage.InsertBuilders return parameters for the expectation previously defined by the When method
func (e *StorageMockInsertBuildersExpectation) Then(err error) *StorageMock {
	e.results = &StorageMockInsertBuildersResults{err}
	return e.mock
}

// InsertBuilders implements storage.Storage
func (m *StorageMock) InsertBuilders(ctx context.Context, querier sqlx.ExecerContext, builders []nausys.Builder) (err error) {
	atomic.AddUint64(&m.beforeInsertBuildersCounter, 1)
	defer atomic.AddUint64(&m.afterInsertBuildersCounter, 1)

	for _, e := range m.InsertBuildersMock.expectations {
		if minimock.Equal(*e.params, StorageMockInsertBuildersParams{ctx, querier, builders}) {
			atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.InsertBuildersMock.defaultExpectation != nil {
		atomic.AddUint64(&m.InsertBuildersMock.defaultExpectation.Counter, 1)
		want := m.InsertBuildersMock.defaultExpectation.params
		got := StorageMockInsertBuildersParams{ctx, querier, builders}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StorageMock.InsertBuilders got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.InsertBuildersMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StorageMock.InsertBuilders")
		}
		return (*results).err
	}
	if m.funcInsertBuilders != nil {
		return m.funcInsertBuilders(ctx, querier, builders)
	}
	m.t.Fatalf("Unexpected call to StorageMock.InsertBuilders. %v %v %v", ctx, querier, builders)
	return
}

// InsertBuildersAfterCounter returns a count of finished StorageMock.InsertBuilders invocations
func (m *StorageMock) InsertBuildersAfterCounter() uint64 {
	return atomic.LoadUint64(&m.afterInsertBuildersCounter)
}

// InsertBuildersBeforeCounter returns a count of StorageMock.InsertBuilders invocations
func (m *StorageMock) InsertBuildersBeforeCounter() uint64 {
	return atomic.LoadUint64(&m.beforeInsertBuildersCounter)
}

// MinimockInsertBuildersDone returns true if the count of the InsertBuilders invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockInsertBuildersDone() bool {
	for _, e := range m.InsertBuildersMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertBuildersMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterInsertBuildersCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertBuilders != nil && atomic.LoadUint64(&m.afterInsertBuildersCounter) < 1 {
		return false
	}
	return true
}

// MinimockInsertBuildersInspect logs each unmet expectation
func (m *StorageMock) MinimockInsertBuildersInspect() {
	for _, e := range m.InsertBuildersMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.InsertBuilders with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertBuildersMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterInsertBuildersCounter) < 1 {
		m.t.Errorf("Expected call to StorageMock.InsertBuilders with params: %#v", *m.InsertBuildersMock.defaultExpectation.params)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertBuilders != nil && atomic.LoadUint64(&m.afterInsertBuildersCounter) < 1 {
		m.t.Error("Expected call to StorageMock.InsertBuilders")
	}
}

type mStorageMockInsertCharters struct {
	mock               *StorageMock
	defaultExpectation *StorageMockInsertChartersExpectation
	expectations       []*StorageMockInsertChartersExpectation
}

// StorageMockInsertChartersExpectation specifies expectation struct of the Storage.InsertCharters
type StorageMockInsertChartersExpectation struct {
	mock    *StorageMock
	params  *StorageMockInsertChartersParams
	results *StorageMockInsertChartersResults
	Counter uint64
}

// StorageMockInsertChartersParams contains parameters of the Storage.InsertCharters
type StorageMockInsertChartersParams struct {
	ctx      context.Context
	querier  sqlx.ExecerContext
	charters []nausys.Charter
}

// StorageMockInsertChartersResults contains results of the Storage.InsertCharters
type StorageMockInsertChartersResults struct {
	err error
}

// Expect sets up expected params for Storage.InsertCharters
func (m *mStorageMockInsertCharters) Expect(ctx context.Context, querier sqlx.ExecerContext, charters []nausys.Charter) *mStorageMockInsertCharters {
	if m.mock.funcInsertCharters != nil {
		m.mock.t.Fatalf("StorageMock.InsertCharters mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockInsertChartersExpectation{}
	}

	m.defaultExpectation.params = &StorageMockInsertChartersParams{ctx, querier, charters}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Storage.InsertCharters
func (m *mStorageMockInsertCharters) Return(err error) *StorageMock {
	if m.mock.funcInsertCharters != nil {
		m.mock.t.Fatalf("StorageMock.InsertCharters mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockInsertChartersExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StorageMockInsertChartersResults{err}
	return m.mock
}

//Set uses given function f to mock the Storage.InsertCharters method
func (m *mStorageMockInsertCharters) Set(f func(ctx context.Context, querier sqlx.ExecerContext, charters []nausys.Charter) (err error)) *StorageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Storage.InsertCharters method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Storage.InsertCharters method")
	}

	m.mock.funcInsertCharters = f
	return m.mock
}

// When sets expectation for the Storage.InsertCharters which will trigger the result defined by the following
// Then helper
func (m *mStorageMockInsertCharters) When(ctx context.Context, querier sqlx.ExecerContext, charters []nausys.Charter) *StorageMockInsertChartersExpectation {
	if m.mock.funcInsertCharters != nil {
		m.mock.t.Fatalf("StorageMock.InsertCharters mock is already set by Set")
	}

	expectation := &StorageMockInsertChartersExpectation{
		mock:   m.mock,
		params: &StorageMockInsertChartersParams{ctx, querier, charters},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Storage.InsertCharters return parameters for the expectation previously defined by the When method
func (e *StorageMockInsertChartersExpectation) Then(err error) *StorageMock {
	e.results = &StorageMockInsertChartersResults{err}
	return e.mock
}

// InsertCharters implements storage.Storage
func (m *StorageMock) InsertCharters(ctx context.Context, querier sqlx.ExecerContext, charters []nausys.Charter) (err error) {
	atomic.AddUint64(&m.beforeInsertChartersCounter, 1)
	defer atomic.AddUint64(&m.afterInsertChartersCounter, 1)

	for _, e := range m.InsertChartersMock.expectations {
		if minimock.Equal(*e.params, StorageMockInsertChartersParams{ctx, querier, charters}) {
			atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.InsertChartersMock.defaultExpectation != nil {
		atomic.AddUint64(&m.InsertChartersMock.defaultExpectation.Counter, 1)
		want := m.InsertChartersMock.defaultExpectation.params
		got := StorageMockInsertChartersParams{ctx, querier, charters}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StorageMock.InsertCharters got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.InsertChartersMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StorageMock.InsertCharters")
		}
		return (*results).err
	}
	if m.funcInsertCharters != nil {
		return m.funcInsertCharters(ctx, querier, charters)
	}
	m.t.Fatalf("Unexpected call to StorageMock.InsertCharters. %v %v %v", ctx, querier, charters)
	return
}

// InsertChartersAfterCounter returns a count of finished StorageMock.InsertCharters invocations
func (m *StorageMock) InsertChartersAfterCounter() uint64 {
	return atomic.LoadUint64(&m.afterInsertChartersCounter)
}

// InsertChartersBeforeCounter returns a count of StorageMock.InsertCharters invocations
func (m *StorageMock) InsertChartersBeforeCounter() uint64 {
	return atomic.LoadUint64(&m.beforeInsertChartersCounter)
}

// MinimockInsertChartersDone returns true if the count of the InsertCharters invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockInsertChartersDone() bool {
	for _, e := range m.InsertChartersMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertChartersMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterInsertChartersCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertCharters != nil && atomic.LoadUint64(&m.afterInsertChartersCounter) < 1 {
		return false
	}
	return true
}

// MinimockInsertChartersInspect logs each unmet expectation
func (m *StorageMock) MinimockInsertChartersInspect() {
	for _, e := range m.InsertChartersMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.InsertCharters with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertChartersMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterInsertChartersCounter) < 1 {
		m.t.Errorf("Expected call to StorageMock.InsertCharters with params: %#v", *m.InsertChartersMock.defaultExpectation.params)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertCharters != nil && atomic.LoadUint64(&m.afterInsertChartersCounter) < 1 {
		m.t.Error("Expected call to StorageMock.InsertCharters")
	}
}

type mStorageMockInsertModels struct {
	mock               *StorageMock
	defaultExpectation *StorageMockInsertModelsExpectation
	expectations       []*StorageMockInsertModelsExpectation
}

// StorageMockInsertModelsExpectation specifies expectation struct of the Storage.InsertModels
type StorageMockInsertModelsExpectation struct {
	mock    *StorageMock
	params  *StorageMockInsertModelsParams
	results *StorageMockInsertModelsResults
	Counter uint64
}

// StorageMockInsertModelsParams contains parameters of the Storage.InsertModels
type StorageMockInsertModelsParams struct {
	ctx     context.Context
	querier sqlx.ExecerContext
	models  []nausys.Model
}

// StorageMockInsertModelsResults contains results of the Storage.InsertModels
type StorageMockInsertModelsResults struct {
	err error
}

// Expect sets up expected params for Storage.InsertModels
func (m *mStorageMockInsertModels) Expect(ctx context.Context, querier sqlx.ExecerContext, models []nausys.Model) *mStorageMockInsertModels {
	if m.mock.funcInsertModels != nil {
		m.mock.t.Fatalf("StorageMock.InsertModels mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockInsertModelsExpectation{}
	}

	m.defaultExpectation.params = &StorageMockInsertModelsParams{ctx, querier, models}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Storage.InsertModels
func (m *mStorageMockInsertModels) Return(err error) *StorageMock {
	if m.mock.funcInsertModels != nil {
		m.mock.t.Fatalf("StorageMock.InsertModels mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockInsertModelsExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StorageMockInsertModelsResults{err}
	return m.mock
}

//Set uses given function f to mock the Storage.InsertModels method
func (m *mStorageMockInsertModels) Set(f func(ctx context.Context, querier sqlx.ExecerContext, models []nausys.Model) (err error)) *StorageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Storage.InsertModels method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Storage.InsertModels method")
	}

	m.mock.funcInsertModels = f
	return m.mock
}

// When sets expectation for the Storage.InsertModels which will trigger the result defined by the following
// Then helper
func (m *mStorageMockInsertModels) When(ctx context.Context, querier sqlx.ExecerContext, models []nausys.Model) *StorageMockInsertModelsExpectation {
	if m.mock.funcInsertModels != nil {
		m.mock.t.Fatalf("StorageMock.InsertModels mock is already set by Set")
	}

	expectation := &StorageMockInsertModelsExpectation{
		mock:   m.mock,
		params: &StorageMockInsertModelsParams{ctx, querier, models},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Storage.InsertModels return parameters for the expectation previously defined by the When method
func (e *StorageMockInsertModelsExpectation) Then(err error) *StorageMock {
	e.results = &StorageMockInsertModelsResults{err}
	return e.mock
}

// InsertModels implements storage.Storage
func (m *StorageMock) InsertModels(ctx context.Context, querier sqlx.ExecerContext, models []nausys.Model) (err error) {
	atomic.AddUint64(&m.beforeInsertModelsCounter, 1)
	defer atomic.AddUint64(&m.afterInsertModelsCounter, 1)

	for _, e := range m.InsertModelsMock.expectations {
		if minimock.Equal(*e.params, StorageMockInsertModelsParams{ctx, querier, models}) {
			atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.InsertModelsMock.defaultExpectation != nil {
		atomic.AddUint64(&m.InsertModelsMock.defaultExpectation.Counter, 1)
		want := m.InsertModelsMock.defaultExpectation.params
		got := StorageMockInsertModelsParams{ctx, querier, models}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StorageMock.InsertModels got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.InsertModelsMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StorageMock.InsertModels")
		}
		return (*results).err
	}
	if m.funcInsertModels != nil {
		return m.funcInsertModels(ctx, querier, models)
	}
	m.t.Fatalf("Unexpected call to StorageMock.InsertModels. %v %v %v", ctx, querier, models)
	return
}

// InsertModelsAfterCounter returns a count of finished StorageMock.InsertModels invocations
func (m *StorageMock) InsertModelsAfterCounter() uint64 {
	return atomic.LoadUint64(&m.afterInsertModelsCounter)
}

// InsertModelsBeforeCounter returns a count of StorageMock.InsertModels invocations
func (m *StorageMock) InsertModelsBeforeCounter() uint64 {
	return atomic.LoadUint64(&m.beforeInsertModelsCounter)
}

// MinimockInsertModelsDone returns true if the count of the InsertModels invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockInsertModelsDone() bool {
	for _, e := range m.InsertModelsMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertModelsMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterInsertModelsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertModels != nil && atomic.LoadUint64(&m.afterInsertModelsCounter) < 1 {
		return false
	}
	return true
}

// MinimockInsertModelsInspect logs each unmet expectation
func (m *StorageMock) MinimockInsertModelsInspect() {
	for _, e := range m.InsertModelsMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.InsertModels with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertModelsMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterInsertModelsCounter) < 1 {
		m.t.Errorf("Expected call to StorageMock.InsertModels with params: %#v", *m.InsertModelsMock.defaultExpectation.params)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertModels != nil && atomic.LoadUint64(&m.afterInsertModelsCounter) < 1 {
		m.t.Error("Expected call to StorageMock.InsertModels")
	}
}

type mStorageMockInsertUpdateInfo struct {
	mock               *StorageMock
	defaultExpectation *StorageMockInsertUpdateInfoExpectation
	expectations       []*StorageMockInsertUpdateInfoExpectation
}

// StorageMockInsertUpdateInfoExpectation specifies expectation struct of the Storage.InsertUpdateInfo
type StorageMockInsertUpdateInfoExpectation struct {
	mock    *StorageMock
	params  *StorageMockInsertUpdateInfoParams
	results *StorageMockInsertUpdateInfoResults
	Counter uint64
}

// StorageMockInsertUpdateInfoParams contains parameters of the Storage.InsertUpdateInfo
type StorageMockInsertUpdateInfoParams struct {
	ctx     context.Context
	querier sqlx.ExecerContext
}

// StorageMockInsertUpdateInfoResults contains results of the Storage.InsertUpdateInfo
type StorageMockInsertUpdateInfoResults struct {
	err error
}

// Expect sets up expected params for Storage.InsertUpdateInfo
func (m *mStorageMockInsertUpdateInfo) Expect(ctx context.Context, querier sqlx.ExecerContext) *mStorageMockInsertUpdateInfo {
	if m.mock.funcInsertUpdateInfo != nil {
		m.mock.t.Fatalf("StorageMock.InsertUpdateInfo mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockInsertUpdateInfoExpectation{}
	}

	m.defaultExpectation.params = &StorageMockInsertUpdateInfoParams{ctx, querier}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Storage.InsertUpdateInfo
func (m *mStorageMockInsertUpdateInfo) Return(err error) *StorageMock {
	if m.mock.funcInsertUpdateInfo != nil {
		m.mock.t.Fatalf("StorageMock.InsertUpdateInfo mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockInsertUpdateInfoExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StorageMockInsertUpdateInfoResults{err}
	return m.mock
}

//Set uses given function f to mock the Storage.InsertUpdateInfo method
func (m *mStorageMockInsertUpdateInfo) Set(f func(ctx context.Context, querier sqlx.ExecerContext) (err error)) *StorageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Storage.InsertUpdateInfo method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Storage.InsertUpdateInfo method")
	}

	m.mock.funcInsertUpdateInfo = f
	return m.mock
}

// When sets expectation for the Storage.InsertUpdateInfo which will trigger the result defined by the following
// Then helper
func (m *mStorageMockInsertUpdateInfo) When(ctx context.Context, querier sqlx.ExecerContext) *StorageMockInsertUpdateInfoExpectation {
	if m.mock.funcInsertUpdateInfo != nil {
		m.mock.t.Fatalf("StorageMock.InsertUpdateInfo mock is already set by Set")
	}

	expectation := &StorageMockInsertUpdateInfoExpectation{
		mock:   m.mock,
		params: &StorageMockInsertUpdateInfoParams{ctx, querier},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Storage.InsertUpdateInfo return parameters for the expectation previously defined by the When method
func (e *StorageMockInsertUpdateInfoExpectation) Then(err error) *StorageMock {
	e.results = &StorageMockInsertUpdateInfoResults{err}
	return e.mock
}

// InsertUpdateInfo implements storage.Storage
func (m *StorageMock) InsertUpdateInfo(ctx context.Context, querier sqlx.ExecerContext) (err error) {
	atomic.AddUint64(&m.beforeInsertUpdateInfoCounter, 1)
	defer atomic.AddUint64(&m.afterInsertUpdateInfoCounter, 1)

	for _, e := range m.InsertUpdateInfoMock.expectations {
		if minimock.Equal(*e.params, StorageMockInsertUpdateInfoParams{ctx, querier}) {
			atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.InsertUpdateInfoMock.defaultExpectation != nil {
		atomic.AddUint64(&m.InsertUpdateInfoMock.defaultExpectation.Counter, 1)
		want := m.InsertUpdateInfoMock.defaultExpectation.params
		got := StorageMockInsertUpdateInfoParams{ctx, querier}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StorageMock.InsertUpdateInfo got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.InsertUpdateInfoMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StorageMock.InsertUpdateInfo")
		}
		return (*results).err
	}
	if m.funcInsertUpdateInfo != nil {
		return m.funcInsertUpdateInfo(ctx, querier)
	}
	m.t.Fatalf("Unexpected call to StorageMock.InsertUpdateInfo. %v %v", ctx, querier)
	return
}

// InsertUpdateInfoAfterCounter returns a count of finished StorageMock.InsertUpdateInfo invocations
func (m *StorageMock) InsertUpdateInfoAfterCounter() uint64 {
	return atomic.LoadUint64(&m.afterInsertUpdateInfoCounter)
}

// InsertUpdateInfoBeforeCounter returns a count of StorageMock.InsertUpdateInfo invocations
func (m *StorageMock) InsertUpdateInfoBeforeCounter() uint64 {
	return atomic.LoadUint64(&m.beforeInsertUpdateInfoCounter)
}

// MinimockInsertUpdateInfoDone returns true if the count of the InsertUpdateInfo invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockInsertUpdateInfoDone() bool {
	for _, e := range m.InsertUpdateInfoMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertUpdateInfoMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterInsertUpdateInfoCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertUpdateInfo != nil && atomic.LoadUint64(&m.afterInsertUpdateInfoCounter) < 1 {
		return false
	}
	return true
}

// MinimockInsertUpdateInfoInspect logs each unmet expectation
func (m *StorageMock) MinimockInsertUpdateInfoInspect() {
	for _, e := range m.InsertUpdateInfoMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.InsertUpdateInfo with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertUpdateInfoMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterInsertUpdateInfoCounter) < 1 {
		m.t.Errorf("Expected call to StorageMock.InsertUpdateInfo with params: %#v", *m.InsertUpdateInfoMock.defaultExpectation.params)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertUpdateInfo != nil && atomic.LoadUint64(&m.afterInsertUpdateInfoCounter) < 1 {
		m.t.Error("Expected call to StorageMock.InsertUpdateInfo")
	}
}

type mStorageMockInsertYachts struct {
	mock               *StorageMock
	defaultExpectation *StorageMockInsertYachtsExpectation
	expectations       []*StorageMockInsertYachtsExpectation
}

// StorageMockInsertYachtsExpectation specifies expectation struct of the Storage.InsertYachts
type StorageMockInsertYachtsExpectation struct {
	mock    *StorageMock
	params  *StorageMockInsertYachtsParams
	results *StorageMockInsertYachtsResults
	Counter uint64
}

// StorageMockInsertYachtsParams contains parameters of the Storage.InsertYachts
type StorageMockInsertYachtsParams struct {
	ctx     context.Context
	querier sqlx.ExecerContext
	yachts  []*nausys.Yacht
}

// StorageMockInsertYachtsResults contains results of the Storage.InsertYachts
type StorageMockInsertYachtsResults struct {
	err error
}

// Expect sets up expected params for Storage.InsertYachts
func (m *mStorageMockInsertYachts) Expect(ctx context.Context, querier sqlx.ExecerContext, yachts []*nausys.Yacht) *mStorageMockInsertYachts {
	if m.mock.funcInsertYachts != nil {
		m.mock.t.Fatalf("StorageMock.InsertYachts mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockInsertYachtsExpectation{}
	}

	m.defaultExpectation.params = &StorageMockInsertYachtsParams{ctx, querier, yachts}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Storage.InsertYachts
func (m *mStorageMockInsertYachts) Return(err error) *StorageMock {
	if m.mock.funcInsertYachts != nil {
		m.mock.t.Fatalf("StorageMock.InsertYachts mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockInsertYachtsExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StorageMockInsertYachtsResults{err}
	return m.mock
}

//Set uses given function f to mock the Storage.InsertYachts method
func (m *mStorageMockInsertYachts) Set(f func(ctx context.Context, querier sqlx.ExecerContext, yachts []*nausys.Yacht) (err error)) *StorageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Storage.InsertYachts method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Storage.InsertYachts method")
	}

	m.mock.funcInsertYachts = f
	return m.mock
}

// When sets expectation for the Storage.InsertYachts which will trigger the result defined by the following
// Then helper
func (m *mStorageMockInsertYachts) When(ctx context.Context, querier sqlx.ExecerContext, yachts []*nausys.Yacht) *StorageMockInsertYachtsExpectation {
	if m.mock.funcInsertYachts != nil {
		m.mock.t.Fatalf("StorageMock.InsertYachts mock is already set by Set")
	}

	expectation := &StorageMockInsertYachtsExpectation{
		mock:   m.mock,
		params: &StorageMockInsertYachtsParams{ctx, querier, yachts},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Storage.InsertYachts return parameters for the expectation previously defined by the When method
func (e *StorageMockInsertYachtsExpectation) Then(err error) *StorageMock {
	e.results = &StorageMockInsertYachtsResults{err}
	return e.mock
}

// InsertYachts implements storage.Storage
func (m *StorageMock) InsertYachts(ctx context.Context, querier sqlx.ExecerContext, yachts []*nausys.Yacht) (err error) {
	atomic.AddUint64(&m.beforeInsertYachtsCounter, 1)
	defer atomic.AddUint64(&m.afterInsertYachtsCounter, 1)

	for _, e := range m.InsertYachtsMock.expectations {
		if minimock.Equal(*e.params, StorageMockInsertYachtsParams{ctx, querier, yachts}) {
			atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.InsertYachtsMock.defaultExpectation != nil {
		atomic.AddUint64(&m.InsertYachtsMock.defaultExpectation.Counter, 1)
		want := m.InsertYachtsMock.defaultExpectation.params
		got := StorageMockInsertYachtsParams{ctx, querier, yachts}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StorageMock.InsertYachts got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.InsertYachtsMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StorageMock.InsertYachts")
		}
		return (*results).err
	}
	if m.funcInsertYachts != nil {
		return m.funcInsertYachts(ctx, querier, yachts)
	}
	m.t.Fatalf("Unexpected call to StorageMock.InsertYachts. %v %v %v", ctx, querier, yachts)
	return
}

// InsertYachtsAfterCounter returns a count of finished StorageMock.InsertYachts invocations
func (m *StorageMock) InsertYachtsAfterCounter() uint64 {
	return atomic.LoadUint64(&m.afterInsertYachtsCounter)
}

// InsertYachtsBeforeCounter returns a count of StorageMock.InsertYachts invocations
func (m *StorageMock) InsertYachtsBeforeCounter() uint64 {
	return atomic.LoadUint64(&m.beforeInsertYachtsCounter)
}

// MinimockInsertYachtsDone returns true if the count of the InsertYachts invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockInsertYachtsDone() bool {
	for _, e := range m.InsertYachtsMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertYachtsMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterInsertYachtsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertYachts != nil && atomic.LoadUint64(&m.afterInsertYachtsCounter) < 1 {
		return false
	}
	return true
}

// MinimockInsertYachtsInspect logs each unmet expectation
func (m *StorageMock) MinimockInsertYachtsInspect() {
	for _, e := range m.InsertYachtsMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.InsertYachts with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertYachtsMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterInsertYachtsCounter) < 1 {
		m.t.Errorf("Expected call to StorageMock.InsertYachts with params: %#v", *m.InsertYachtsMock.defaultExpectation.params)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertYachts != nil && atomic.LoadUint64(&m.afterInsertYachtsCounter) < 1 {
		m.t.Error("Expected call to StorageMock.InsertYachts")
	}
}

type mStorageMockWithTransaction struct {
	mock               *StorageMock
	defaultExpectation *StorageMockWithTransactionExpectation
	expectations       []*StorageMockWithTransactionExpectation
}

// StorageMockWithTransactionExpectation specifies expectation struct of the Storage.WithTransaction
type StorageMockWithTransactionExpectation struct {
	mock    *StorageMock
	params  *StorageMockWithTransactionParams
	results *StorageMockWithTransactionResults
	Counter uint64
}

// StorageMockWithTransactionParams contains parameters of the Storage.WithTransaction
type StorageMockWithTransactionParams struct {
	ctx context.Context
	f   func(tx *sqlx.Tx) error
}

// StorageMockWithTransactionResults contains results of the Storage.WithTransaction
type StorageMockWithTransactionResults struct {
	err error
}

// Expect sets up expected params for Storage.WithTransaction
func (m *mStorageMockWithTransaction) Expect(ctx context.Context, f func(tx *sqlx.Tx) error) *mStorageMockWithTransaction {
	if m.mock.funcWithTransaction != nil {
		m.mock.t.Fatalf("StorageMock.WithTransaction mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockWithTransactionExpectation{}
	}

	m.defaultExpectation.params = &StorageMockWithTransactionParams{ctx, f}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Storage.WithTransaction
func (m *mStorageMockWithTransaction) Return(err error) *StorageMock {
	if m.mock.funcWithTransaction != nil {
		m.mock.t.Fatalf("StorageMock.WithTransaction mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StorageMockWithTransactionExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StorageMockWithTransactionResults{err}
	return m.mock
}

//Set uses given function f to mock the Storage.WithTransaction method
func (m *mStorageMockWithTransaction) Set(f func(ctx context.Context, f func(tx *sqlx.Tx) error) (err error)) *StorageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Storage.WithTransaction method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Storage.WithTransaction method")
	}

	m.mock.funcWithTransaction = f
	return m.mock
}

// When sets expectation for the Storage.WithTransaction which will trigger the result defined by the following
// Then helper
func (m *mStorageMockWithTransaction) When(ctx context.Context, f func(tx *sqlx.Tx) error) *StorageMockWithTransactionExpectation {
	if m.mock.funcWithTransaction != nil {
		m.mock.t.Fatalf("StorageMock.WithTransaction mock is already set by Set")
	}

	expectation := &StorageMockWithTransactionExpectation{
		mock:   m.mock,
		params: &StorageMockWithTransactionParams{ctx, f},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Storage.WithTransaction return parameters for the expectation previously defined by the When method
func (e *StorageMockWithTransactionExpectation) Then(err error) *StorageMock {
	e.results = &StorageMockWithTransactionResults{err}
	return e.mock
}

// WithTransaction implements storage.Storage
func (m *StorageMock) WithTransaction(ctx context.Context, f func(tx *sqlx.Tx) error) (err error) {
	atomic.AddUint64(&m.beforeWithTransactionCounter, 1)
	defer atomic.AddUint64(&m.afterWithTransactionCounter, 1)

	for _, e := range m.WithTransactionMock.expectations {
		if minimock.Equal(*e.params, StorageMockWithTransactionParams{ctx, f}) {
			atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.WithTransactionMock.defaultExpectation != nil {
		atomic.AddUint64(&m.WithTransactionMock.defaultExpectation.Counter, 1)
		want := m.WithTransactionMock.defaultExpectation.params
		got := StorageMockWithTransactionParams{ctx, f}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StorageMock.WithTransaction got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.WithTransactionMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StorageMock.WithTransaction")
		}
		return (*results).err
	}
	if m.funcWithTransaction != nil {
		return m.funcWithTransaction(ctx, f)
	}
	m.t.Fatalf("Unexpected call to StorageMock.WithTransaction. %v %v", ctx, f)
	return
}

// WithTransactionAfterCounter returns a count of finished StorageMock.WithTransaction invocations
func (m *StorageMock) WithTransactionAfterCounter() uint64 {
	return atomic.LoadUint64(&m.afterWithTransactionCounter)
}

// WithTransactionBeforeCounter returns a count of StorageMock.WithTransaction invocations
func (m *StorageMock) WithTransactionBeforeCounter() uint64 {
	return atomic.LoadUint64(&m.beforeWithTransactionCounter)
}

// MinimockWithTransactionDone returns true if the count of the WithTransaction invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockWithTransactionDone() bool {
	for _, e := range m.WithTransactionMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.WithTransactionMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterWithTransactionCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcWithTransaction != nil && atomic.LoadUint64(&m.afterWithTransactionCounter) < 1 {
		return false
	}
	return true
}

// MinimockWithTransactionInspect logs each unmet expectation
func (m *StorageMock) MinimockWithTransactionInspect() {
	for _, e := range m.WithTransactionMock.expectations {
		if atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.WithTransaction with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.WithTransactionMock.defaultExpectation != nil && atomic.LoadUint64(&m.afterWithTransactionCounter) < 1 {
		m.t.Errorf("Expected call to StorageMock.WithTransaction with params: %#v", *m.WithTransactionMock.defaultExpectation.params)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcWithTransaction != nil && atomic.LoadUint64(&m.afterWithTransactionCounter) < 1 {
		m.t.Error("Expected call to StorageMock.WithTransaction")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StorageMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockClearAllInspect()

		m.MinimockFindBuildersByPrefixInspect()

		m.MinimockFindModelsByPrefixInspect()

		m.MinimockFindYachtsInspect()

		m.MinimockGetLastUpdateInfoInspect()

		m.MinimockInsertBuildersInspect()

		m.MinimockInsertChartersInspect()

		m.MinimockInsertModelsInspect()

		m.MinimockInsertUpdateInfoInspect()

		m.MinimockInsertYachtsInspect()

		m.MinimockWithTransactionInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *StorageMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-time.After(10 * time.Millisecond):
		}
	}
}

func (m *StorageMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockClearAllDone() &&
		m.MinimockFindBuildersByPrefixDone() &&
		m.MinimockFindModelsByPrefixDone() &&
		m.MinimockFindYachtsDone() &&
		m.MinimockGetLastUpdateInfoDone() &&
		m.MinimockInsertBuildersDone() &&
		m.MinimockInsertChartersDone() &&
		m.MinimockInsertModelsDone() &&
		m.MinimockInsertUpdateInfoDone() &&
		m.MinimockInsertYachtsDone() &&
		m.MinimockWithTransactionDone()
}
