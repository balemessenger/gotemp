package test

import (
	"testing"
)

func setup() {
}

func teardown() {

}

func TestMain(m *testing.M) {
	setup()
	r := m.Run()
	teardown()
}
