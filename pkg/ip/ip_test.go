package ip_test

import (
	"testing"

	"github.com/vndr/jv/pkg/ip"
)

func TestGetLocalIPWithInterfaceCheck(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("GetLocalIPWithInterfaceCheck() panicked: %v", r)
		}
	}()
	ip.GetLocalIPWithInterfaceCheck()
}

func TestGetPublicIP(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("GetPublicIP() panicked: %v", r)
		}
	}()
	ip.GetPublicIP()
}
