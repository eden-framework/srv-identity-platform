package providers

import (
	"fmt"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers/common"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers/dingding"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers/ewechat"
)

var P Provider

func Initializer() error {
	if P.providers == nil {
		P.providers = make(map[enums.BindType]common.Provider)
	}
	if global.ProviderConfig.DingDing.Enabled {
		P.RegisterProvider(dingding.NewDingDingProvider(global.ProviderConfig.DingDing))
	}
	if global.ProviderConfig.EWechat.Enabled {
		P.RegisterProvider(ewechat.NewEWechatProvider(global.ProviderConfig.EWechat))
	}
	return nil
}

type Provider struct {
	providers map[enums.BindType]common.Provider
}

func (p *Provider) RegisterProvider(provider common.Provider) {
	if _, exist := p.providers[provider.ProviderID()]; exist {
		panic(fmt.Sprintf("already exist provider [%s]", provider.ProviderID()))
	}
	p.providers[provider.ProviderID()] = provider
}

func (p *Provider) GetProvider(typ enums.BindType) (common.Provider, bool) {
	provider, exist := p.providers[typ]
	return provider, exist
}

func (p *Provider) GetAvailableProviderInfo() []common.ProviderInfo {
	list := make([]common.ProviderInfo, 0)
	for _, provider := range p.providers {
		list = append(list, common.ProviderInfo{
			ID:    provider.ProviderID(),
			Name:  provider.ProviderName(),
			Icon:  provider.ProviderIcon(),
			Entry: provider.ProviderEntry(),
		})
	}
	return list
}
