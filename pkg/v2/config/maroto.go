package config

import "github.com/johnfercher/maroto/pkg/v2/provider"

type Maroto struct {
	ProviderType provider.Type
	Dimensions   *Dimensions
	Margins      *Margins
	ThreadPool   int
}