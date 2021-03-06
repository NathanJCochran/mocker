package example

import (
	"fmt"
	"html/template"
	. "os"
	"sync/atomic"
	renamed "text/template"
)

// MyInterfaceMock is a mock implementation of the MyInterface
// interface.
type MyInterfaceMock struct {
	NoParamsOrReturnStub                     func()
	NoParamsOrReturnCalled                   int32
	UnnamedParamStub                         func(string)
	UnnamedParamCalled                       int32
	UnnamedVariadicParamStub                 func(...string)
	UnnamedVariadicParamCalled               int32
	BlankParamStub                           func(_ string)
	BlankParamCalled                         int32
	BlankVariadicParamStub                   func(_ ...string)
	BlankVariadicParamCalled                 int32
	NamedParamStub                           func(str string)
	NamedParamCalled                         int32
	NamedVariadicParamStub                   func(strs ...string)
	NamedVariadicParamCalled                 int32
	SameTypeNamedParamsStub                  func(str1 string, str2 string)
	SameTypeNamedParamsCalled                int32
	ImportedParamStub                        func(tmpl template.Template)
	ImportedParamCalled                      int32
	ImportedVariadicParamStub                func(tmpl ...template.Template)
	ImportedVariadicParamCalled              int32
	RenamedImportParamStub                   func(tmpl renamed.Template)
	RenamedImportParamCalled                 int32
	RenamedImportVariadicParamStub           func(tmpls ...renamed.Template)
	RenamedImportVariadicParamCalled         int32
	DotImportParamStub                       func(file File)
	DotImportParamCalled                     int32
	DotImportVariadicParamStub               func(files ...File)
	DotImportVariadicParamCalled             int32
	SelfReferentialParamStub                 func(intf MyInterface)
	SelfReferentialParamCalled               int32
	SelfReferentialVariadicParamStub         func(intf ...MyInterface)
	SelfReferentialVariadicParamCalled       int32
	StructParamStub                          func(obj struct{ num int })
	StructParamCalled                        int32
	StructVariadicParamStub                  func(objs ...struct{ num int })
	StructVariadicParamCalled                int32
	EmbeddedStructParamStub                  func(obj struct{ int })
	EmbeddedStructParamCalled                int32
	EmbeddedStructVariadicParamStub          func(objs ...struct{ int })
	EmbeddedStructVariadicParamCalled        int32
	EmptyInterfaceParamStub                  func(intf interface{})
	EmptyInterfaceParamCalled                int32
	EmptyInterfaceVariadicParamStub          func(intf ...interface{})
	EmptyInterfaceVariadicParamCalled        int32
	InterfaceParamStub                       func(intf interface{ MyFunc(num int) error })
	InterfaceParamCalled                     int32
	InterfaceVariadicParamStub               func(intf ...interface{ MyFunc(num int) error })
	InterfaceVariadicParamCalled             int32
	InterfaceVariadicFuncParamStub           func(intf interface{ MyFunc(nums ...int) error })
	InterfaceVariadicFuncParamCalled         int32
	InterfaceVariadicFuncVariadicParamStub   func(intf ...interface{ MyFunc(nums ...int) error })
	InterfaceVariadicFuncVariadicParamCalled int32
	EmbeddedInterfaceParamStub               func(intf interface{ fmt.Stringer })
	EmbeddedInterfaceParamCalled             int32
	UnnamedReturnStub                        func() error
	UnnamedReturnCalled                      int32
	MultipleUnnamedReturnStub                func() (int, error)
	MultipleUnnamedReturnCalled              int32
	BlankReturnStub                          func() (_ error)
	BlankReturnCalled                        int32
	NamedReturnStub                          func() (err error)
	NamedReturnCalled                        int32
	SameTypeNamedReturnStub                  func() (err1 error, err2 error)
	SameTypeNamedReturnCalled                int32
	RenamedImportReturnStub                  func() (tmpl renamed.Template)
	RenamedImportReturnCalled                int32
	DotImportReturnStub                      func() (file File)
	DotImportReturnCalled                    int32
	SelfReferentialReturnStub                func() (intf MyInterface)
	SelfReferentialReturnCalled              int32
	StructReturnStub                         func() (obj struct{ num int })
	StructReturnCalled                       int32
	EmbeddedStructReturnStub                 func() (obj struct{ int })
	EmbeddedStructReturnCalled               int32
	EmptyInterfaceReturnStub                 func() (intf interface{})
	EmptyInterfaceReturnCalled               int32
	InterfaceReturnStub                      func() (intf interface{ MyFunc(num int) error })
	InterfaceReturnCalled                    int32
	InterfaceVariadicFuncReturnStub          func() (intf interface{ MyFunc(nums ...int) error })
	InterfaceVariadicFuncReturnCalled        int32
	EmbeddedInterfaceReturnStub              func() (intf interface{ fmt.Stringer })
	EmbeddedInterfaceReturnCalled            int32
}

var _ MyInterface = &MyInterfaceMock{}

// NoParamsOrReturn is a stub for the MyInterface.NoParamsOrReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) NoParamsOrReturn() {
	atomic.AddInt32(&m.NoParamsOrReturnCalled, 1)
	m.NoParamsOrReturnStub()
}

// UnnamedParam is a stub for the MyInterface.UnnamedParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) UnnamedParam(param1 string) {
	atomic.AddInt32(&m.UnnamedParamCalled, 1)
	m.UnnamedParamStub(param1)
}

// UnnamedVariadicParam is a stub for the MyInterface.UnnamedVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) UnnamedVariadicParam(param1 ...string) {
	atomic.AddInt32(&m.UnnamedVariadicParamCalled, 1)
	m.UnnamedVariadicParamStub(param1...)
}

// BlankParam is a stub for the MyInterface.BlankParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) BlankParam(param1 string) {
	atomic.AddInt32(&m.BlankParamCalled, 1)
	m.BlankParamStub(param1)
}

// BlankVariadicParam is a stub for the MyInterface.BlankVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) BlankVariadicParam(param1 ...string) {
	atomic.AddInt32(&m.BlankVariadicParamCalled, 1)
	m.BlankVariadicParamStub(param1...)
}

// NamedParam is a stub for the MyInterface.NamedParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) NamedParam(str string) {
	atomic.AddInt32(&m.NamedParamCalled, 1)
	m.NamedParamStub(str)
}

// NamedVariadicParam is a stub for the MyInterface.NamedVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) NamedVariadicParam(strs ...string) {
	atomic.AddInt32(&m.NamedVariadicParamCalled, 1)
	m.NamedVariadicParamStub(strs...)
}

// SameTypeNamedParams is a stub for the MyInterface.SameTypeNamedParams
// method that records the number of times it has been called.
func (m *MyInterfaceMock) SameTypeNamedParams(str1 string, str2 string) {
	atomic.AddInt32(&m.SameTypeNamedParamsCalled, 1)
	m.SameTypeNamedParamsStub(str1, str2)
}

// ImportedParam is a stub for the MyInterface.ImportedParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) ImportedParam(tmpl template.Template) {
	atomic.AddInt32(&m.ImportedParamCalled, 1)
	m.ImportedParamStub(tmpl)
}

// ImportedVariadicParam is a stub for the MyInterface.ImportedVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) ImportedVariadicParam(tmpl ...template.Template) {
	atomic.AddInt32(&m.ImportedVariadicParamCalled, 1)
	m.ImportedVariadicParamStub(tmpl...)
}

// RenamedImportParam is a stub for the MyInterface.RenamedImportParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) RenamedImportParam(tmpl renamed.Template) {
	atomic.AddInt32(&m.RenamedImportParamCalled, 1)
	m.RenamedImportParamStub(tmpl)
}

// RenamedImportVariadicParam is a stub for the MyInterface.RenamedImportVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) RenamedImportVariadicParam(tmpls ...renamed.Template) {
	atomic.AddInt32(&m.RenamedImportVariadicParamCalled, 1)
	m.RenamedImportVariadicParamStub(tmpls...)
}

// DotImportParam is a stub for the MyInterface.DotImportParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) DotImportParam(file File) {
	atomic.AddInt32(&m.DotImportParamCalled, 1)
	m.DotImportParamStub(file)
}

// DotImportVariadicParam is a stub for the MyInterface.DotImportVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) DotImportVariadicParam(files ...File) {
	atomic.AddInt32(&m.DotImportVariadicParamCalled, 1)
	m.DotImportVariadicParamStub(files...)
}

// SelfReferentialParam is a stub for the MyInterface.SelfReferentialParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) SelfReferentialParam(intf MyInterface) {
	atomic.AddInt32(&m.SelfReferentialParamCalled, 1)
	m.SelfReferentialParamStub(intf)
}

// SelfReferentialVariadicParam is a stub for the MyInterface.SelfReferentialVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) SelfReferentialVariadicParam(intf ...MyInterface) {
	atomic.AddInt32(&m.SelfReferentialVariadicParamCalled, 1)
	m.SelfReferentialVariadicParamStub(intf...)
}

// StructParam is a stub for the MyInterface.StructParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) StructParam(obj struct{ num int }) {
	atomic.AddInt32(&m.StructParamCalled, 1)
	m.StructParamStub(obj)
}

// StructVariadicParam is a stub for the MyInterface.StructVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) StructVariadicParam(objs ...struct{ num int }) {
	atomic.AddInt32(&m.StructVariadicParamCalled, 1)
	m.StructVariadicParamStub(objs...)
}

// EmbeddedStructParam is a stub for the MyInterface.EmbeddedStructParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) EmbeddedStructParam(obj struct{ int }) {
	atomic.AddInt32(&m.EmbeddedStructParamCalled, 1)
	m.EmbeddedStructParamStub(obj)
}

// EmbeddedStructVariadicParam is a stub for the MyInterface.EmbeddedStructVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) EmbeddedStructVariadicParam(objs ...struct{ int }) {
	atomic.AddInt32(&m.EmbeddedStructVariadicParamCalled, 1)
	m.EmbeddedStructVariadicParamStub(objs...)
}

// EmptyInterfaceParam is a stub for the MyInterface.EmptyInterfaceParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) EmptyInterfaceParam(intf interface{}) {
	atomic.AddInt32(&m.EmptyInterfaceParamCalled, 1)
	m.EmptyInterfaceParamStub(intf)
}

// EmptyInterfaceVariadicParam is a stub for the MyInterface.EmptyInterfaceVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) EmptyInterfaceVariadicParam(intf ...interface{}) {
	atomic.AddInt32(&m.EmptyInterfaceVariadicParamCalled, 1)
	m.EmptyInterfaceVariadicParamStub(intf...)
}

// InterfaceParam is a stub for the MyInterface.InterfaceParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) InterfaceParam(intf interface{ MyFunc(num int) error }) {
	atomic.AddInt32(&m.InterfaceParamCalled, 1)
	m.InterfaceParamStub(intf)
}

// InterfaceVariadicParam is a stub for the MyInterface.InterfaceVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) InterfaceVariadicParam(intf ...interface{ MyFunc(num int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicParamCalled, 1)
	m.InterfaceVariadicParamStub(intf...)
}

// InterfaceVariadicFuncParam is a stub for the MyInterface.InterfaceVariadicFuncParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) InterfaceVariadicFuncParam(intf interface{ MyFunc(nums ...int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicFuncParamCalled, 1)
	m.InterfaceVariadicFuncParamStub(intf)
}

// InterfaceVariadicFuncVariadicParam is a stub for the MyInterface.InterfaceVariadicFuncVariadicParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) InterfaceVariadicFuncVariadicParam(intf ...interface{ MyFunc(nums ...int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicFuncVariadicParamCalled, 1)
	m.InterfaceVariadicFuncVariadicParamStub(intf...)
}

// EmbeddedInterfaceParam is a stub for the MyInterface.EmbeddedInterfaceParam
// method that records the number of times it has been called.
func (m *MyInterfaceMock) EmbeddedInterfaceParam(intf interface{ fmt.Stringer }) {
	atomic.AddInt32(&m.EmbeddedInterfaceParamCalled, 1)
	m.EmbeddedInterfaceParamStub(intf)
}

// UnnamedReturn is a stub for the MyInterface.UnnamedReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) UnnamedReturn() error {
	atomic.AddInt32(&m.UnnamedReturnCalled, 1)
	return m.UnnamedReturnStub()
}

// MultipleUnnamedReturn is a stub for the MyInterface.MultipleUnnamedReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) MultipleUnnamedReturn() (int, error) {
	atomic.AddInt32(&m.MultipleUnnamedReturnCalled, 1)
	return m.MultipleUnnamedReturnStub()
}

// BlankReturn is a stub for the MyInterface.BlankReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) BlankReturn() (_ error) {
	atomic.AddInt32(&m.BlankReturnCalled, 1)
	return m.BlankReturnStub()
}

// NamedReturn is a stub for the MyInterface.NamedReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) NamedReturn() (err error) {
	atomic.AddInt32(&m.NamedReturnCalled, 1)
	return m.NamedReturnStub()
}

// SameTypeNamedReturn is a stub for the MyInterface.SameTypeNamedReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) SameTypeNamedReturn() (err1 error, err2 error) {
	atomic.AddInt32(&m.SameTypeNamedReturnCalled, 1)
	return m.SameTypeNamedReturnStub()
}

// RenamedImportReturn is a stub for the MyInterface.RenamedImportReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) RenamedImportReturn() (tmpl renamed.Template) {
	atomic.AddInt32(&m.RenamedImportReturnCalled, 1)
	return m.RenamedImportReturnStub()
}

// DotImportReturn is a stub for the MyInterface.DotImportReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) DotImportReturn() (file File) {
	atomic.AddInt32(&m.DotImportReturnCalled, 1)
	return m.DotImportReturnStub()
}

// SelfReferentialReturn is a stub for the MyInterface.SelfReferentialReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) SelfReferentialReturn() (intf MyInterface) {
	atomic.AddInt32(&m.SelfReferentialReturnCalled, 1)
	return m.SelfReferentialReturnStub()
}

// StructReturn is a stub for the MyInterface.StructReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) StructReturn() (obj struct{ num int }) {
	atomic.AddInt32(&m.StructReturnCalled, 1)
	return m.StructReturnStub()
}

// EmbeddedStructReturn is a stub for the MyInterface.EmbeddedStructReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) EmbeddedStructReturn() (obj struct{ int }) {
	atomic.AddInt32(&m.EmbeddedStructReturnCalled, 1)
	return m.EmbeddedStructReturnStub()
}

// EmptyInterfaceReturn is a stub for the MyInterface.EmptyInterfaceReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) EmptyInterfaceReturn() (intf interface{}) {
	atomic.AddInt32(&m.EmptyInterfaceReturnCalled, 1)
	return m.EmptyInterfaceReturnStub()
}

// InterfaceReturn is a stub for the MyInterface.InterfaceReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) InterfaceReturn() (intf interface{ MyFunc(num int) error }) {
	atomic.AddInt32(&m.InterfaceReturnCalled, 1)
	return m.InterfaceReturnStub()
}

// InterfaceVariadicFuncReturn is a stub for the MyInterface.InterfaceVariadicFuncReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) InterfaceVariadicFuncReturn() (intf interface{ MyFunc(nums ...int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicFuncReturnCalled, 1)
	return m.InterfaceVariadicFuncReturnStub()
}

// EmbeddedInterfaceReturn is a stub for the MyInterface.EmbeddedInterfaceReturn
// method that records the number of times it has been called.
func (m *MyInterfaceMock) EmbeddedInterfaceReturn() (intf interface{ fmt.Stringer }) {
	atomic.AddInt32(&m.EmbeddedInterfaceReturnCalled, 1)
	return m.EmbeddedInterfaceReturnStub()
}
