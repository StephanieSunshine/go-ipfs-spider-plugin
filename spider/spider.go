package spider

import (
	"fmt"

	plugin "github.com/ipfs/go-ipfs/plugin"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
)

type SpiderPlugin struct{}

var _ plugin.PluginDaemon = (*SpiderPlugin)(nil)

// Name returns the plugin's name, satisfying the plugin.Plugin interface.
func (*SpiderPlugin) Name() string {
	return "greeter"
}

// Version returns the plugin's version, satisfying the plugin.Plugin interface.
func (*SpiderPlugin) Version() string {
	return "0.1.0"
}

// Init initializes plugin, satisfying the plugin.Plugin interface. Put any
// initialization logic here.
func (*SpiderPlugin) Init(env *plugin.Environment) error {
	return nil
}

func (*SpiderPlugin) Start(_ coreiface.CoreAPI) error {
	fmt.Println("Spider distributed web caching server starting listener")
	return nil
}

func (*SpiderPlugin) Close() error {
	fmt.Println("Spider distributed web caching server stopping")
	return nil
}
