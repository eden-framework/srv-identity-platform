package providers

import (
	"fmt"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers/common"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers/dingding"
)

var P Provider

func Initializer() error {
	if P.providers == nil {
		P.providers = make(map[string]common.Provider)
	}
	P.RegisterProvider(dingding.NewDingDingProvider(global.ProviderConfig.DingDing))
	return nil
}

type Provider struct {
	providers map[string]common.Provider
}

func (p *Provider) RegisterProvider(provider common.Provider) {
	if _, exist := p.providers[provider.ProviderID().String()]; exist {
		panic(fmt.Sprintf("already exist provider [%s]", provider.ProviderID()))
	}
	p.providers[provider.ProviderID().String()] = provider
}

func (p *Provider) GetProvider(typ enums.BindType) (common.Provider, bool) {
	provider, exist := p.providers[typ.String()]
	return provider, exist
}
