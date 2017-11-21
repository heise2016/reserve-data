package binance

import (
	"os"
)

type Interface interface {
	PublicEndpoint() string
	AuthenticatedEndpoint() string
}

type RealInterface struct{}

func (self *RealInterface) PublicEndpoint() string {
	return "https://www.binance.com"
}

func (self *RealInterface) AuthenticatedEndpoint() string {
	return "https://www.binance.com"
}

func NewRealInterface() *RealInterface {
	return &RealInterface{}
}

type SimulatedInterface struct{}

func (self *SimulatedInterface) baseurl() string {
	baseurl := "127.0.0.1"
	if len(os.Args) > 1 {
		baseurl = os.Args[1]
	}
	return baseurl + ":5100"
}

func (self *SimulatedInterface) PublicEndpoint() string {
	return self.baseurl()
}

func (self *SimulatedInterface) AuthenticatedEndpoint() string {
	return self.baseurl()
}

func NewSimulatedInterface() *SimulatedInterface {
	return &SimulatedInterface{}
}
