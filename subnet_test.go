package subnet_test

import (
	"fmt"
	"testing"

	"github.com/dajohi/subnet"
)

func TestSubnet(t *testing.T) {
	s, err := subnet.New("10.30.20.18/29")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	for ip := s.Begin(); ip != nil; ip = s.Next() {
		fmt.Printf("%v\n", ip)
	}
}
