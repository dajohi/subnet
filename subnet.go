package subnet

import (
	"encoding/binary"
	"errors"
	"net"
)

type Subnet struct {
	ipnet *net.IPNet
	loc   uint32

	nwValue uint32
	bcValue uint32
	err     error
}

func (s *Subnet) Begin() net.IP {
	s.loc = 0

	var b [4]byte
	binary.BigEndian.PutUint32(b[:], s.nwValue)
	return net.IPv4(b[0], b[1], b[2], b[3])
}

func (s *Subnet) Next() net.IP {
	if s.nwValue+s.loc+1 > s.bcValue {
		return nil
	}
	s.loc++

	var b [4]byte
	binary.BigEndian.PutUint32(b[:], s.nwValue+s.loc)
	return net.IPv4(b[0], b[1], b[2], b[3])
}

func (s *Subnet) Prev() net.IP {
	if s.nwValue-s.loc-1 < s.nwValue {
		return nil
	}
	s.loc--

	var b [4]byte
	binary.BigEndian.PutUint32(b[:], s.nwValue+s.loc)
	return net.IPv4(b[0], b[1], b[2], b[3])
}

func New(subnet string) (*Subnet, error) {
	_, ipnet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, err
	}

	if ipnet.IP.DefaultMask() == nil {
		return nil, errors.New("only ipv4 supported")
	}

	ones, _ := ipnet.Mask.Size()
	nwValue := binary.BigEndian.Uint32(ipnet.IP.To4())
	bcValue := nwValue | (0xffffffff >> uint32(ones))

	s := Subnet{
		ipnet:   ipnet,
		nwValue: nwValue,
		bcValue: bcValue,
	}

	return &s, nil
}
