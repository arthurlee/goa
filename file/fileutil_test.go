package fileutil

import (
	"github.com/kardianos/osext"
	"os"
	"testing"
)

// Exist

func TestExist(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}

	ok := Exist(wd)
	if !ok {
		t.Errorf("%q should be exist", wd)
	}
}

func TestNotExist(t *testing.T) {
	name := "afab#R@$#3rwq@#@"
	ok := Exist(name)
	if ok {
		t.Errorf("%q should not be exist", name)
	}
}

func TestIsDir(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}

	ok := IsDir(wd)
	if !ok {
		t.Errorf("%q should be a directory", wd)
	}
}

func TestNotIsDir(t *testing.T) {
	name := "afab#R@$#3rwq@#@"
	ok := IsDir(name)
	if ok {
		t.Errorf("%q should not be a directory", name)
	}
}

func TestIsFile(t *testing.T) {
	filename, err := osext.Executable()
	if err != nil {
		t.Error(err)
		return
	}

	ok := IsFile(filename)
	if !ok {
		t.Errorf("%q should be a file", filename)
	}
}

func TestNotIsFile(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}

	ok := IsFile(wd)
	if ok {
		t.Errorf("%q should not be a file", wd)
	}
}
