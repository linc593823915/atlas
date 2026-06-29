package bootstrap

import "testing"

func TestNewBootstrap(t *testing.T) {
	bootstrap, err := New()
	if err != nil {
		t.Fatal(err)
	}
	err = bootstrap.Start()
	if err != nil {
		return
	}
	err = bootstrap.Stop()
	if err != nil {
		return
	}
	t.Log("test success")
}
