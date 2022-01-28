//go:build wireinject
// +build wireinject

package app

import (
	internal "github.com/ghzmhrm/go-blockchain/internal"
	"github.com/ghzmhrm/go-blockchain/internal/service"
	"github.com/google/wire"
)

var (
	servicesSet = wire.NewSet(
		service.NewBlockChainServices,
		wire.Bind(new(internal.BlockchainSvc), new(*service.BlockChainServices)),
	)

	allSet = wire.NewSet(
		servicesSet,
	)
)

// InitializeBlockchain init the blockchain application
func InitializeBlockchain() internal.BlockchainSvc {
	wire.Build(allSet)
	return new(service.BlockChainServices)
}
