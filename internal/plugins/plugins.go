package plugins

import (
	"encoding/json"
	"github.com/bobbiegoede/gezondheid/internal/handlers"
	"log"
	"plugin"
)

type Plugin interface {
	Run(getContext func() []byte, next func())
	SetConfig(config []byte)
}

func LoadPlugin(path string) Plugin {
	plugin, err := plugin.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	sym, err := plugin.Lookup("Plugin")
	if err != nil {
		log.Fatalf("Error looking up symbol:\n%v", err)
	}

	middleware, ok := sym.(Plugin)
	if !ok {
		log.Fatalf("Invalid plugin type:\n%v", err)
	}

	return middleware
}

type PluginHandler struct {
	Plugin Plugin
	next   handlers.Handler
}

func (h *PluginHandler) HandleRequest(ctx *handlers.Ctx) {
	h.Plugin.Run(func() []byte {
		jsonData, err := json.Marshal(ctx)
		if err != nil {
			log.Fatalf("Error serializing the struct:\n%v", err)
		}
		return jsonData
	}, func() {
		if h.next != nil {
			h.next.HandleRequest(ctx)
		}
	})
}

func (h *PluginHandler) SetNext(handler handlers.Handler) {
	h.next = handler
}
