package lux

import (
	"errors"
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"runtime"
	"testing"
	"time"
)

var testchan chan func()
var donechan chan struct{}

func TestMain(m *testing.M) {
	runtime.LockOSThread()
	testchan = make(chan func())
	donechan = make(chan struct{})
	InitGLFW()
	_ = StartHeadless()
	go m.Run()
	select {
	case test := <-testchan:
		test()
		donechan <- struct{}{}
	case <-time.After(5 * time.Second):
		break
	}
	MustNotGLError()
	TerminateGLFW()
}

func TestGLerrorString(t *testing.T) {
	if GLErrorToString(gl.NO_ERROR) != "GL_NO_ERROR" {
		t.Error(errors.New("error trying to get NO_ERROR to string"))
	}
	if GLErrorToString(gl.INVALID_ENUM) != "GL_INVALID_ENUM" {
		t.Error(errors.New("error trying to get INVALID_ENUM to string"))
	}
	if GLErrorToString(gl.INVALID_VALUE) != "GL_INVALID_VALUE" {
		t.Error(errors.New("error trying to get INVALID_VALUE to string"))
	}
	if GLErrorToString(gl.INVALID_OPERATION) != "GL_INVALID_OPERATION" {
		t.Error(errors.New("error trying to get INVALID_OPERATION to string"))
	}
	if GLErrorToString(gl.INVALID_FRAMEBUFFER_OPERATION) != "GL_INVALID_FRAMEBUFFER_OPERATION" {
		t.Error(errors.New("error trying to get INVALID_FRAMEBUFFER_OPERATION to string"))
	}
	if GLErrorToString(gl.OUT_OF_MEMORY) != "GL_OUT_OF_MEMORY" {
		t.Error(errors.New("error trying to get OUT_OF_MEMORY to string"))
	}
	if GLErrorToString(gl.STACK_UNDERFLOW) != "GL_STACK_UNDERFLOW" {
		t.Error(errors.New("error trying to get STACK_UNDERFLOW to string"))
	}
	if GLErrorToString(gl.STACK_OVERFLOW) != "GL_STACK_OVERFLOW" {
		t.Error(errors.New("error trying to get STACK_OVERFLOW to string"))
	}
	if GLErrorToString(1353) != "Unknown Error Code" {
		t.Error(errors.New("error trying to get unknown error to string"))
	}
}

func TestMustNotGLError(t *testing.T) {
	testchan <- func() {
		gl.Enable(gl.DEPTH_TEST)
		if gl.GetError() != gl.NO_ERROR {
			t.Fail()
		}
	}
	<-donechan
}
