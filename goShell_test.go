package goShell

import (
	"testing"
)

func TestReverseShell(t *testing.T) {
	//go reShell("127.0.0.1", "8083")
	go OpenShell("0.0.0.0:8082")
	select {}
}
