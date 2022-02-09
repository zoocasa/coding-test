package main

import (
	"testing"
)

func TestServer(t *testing.T) {
	go StartServer()

	t.Run("Sample test", func(t *testing.T) {
		if true {
			t.Errorf("Test fails")
		}
	})
}
