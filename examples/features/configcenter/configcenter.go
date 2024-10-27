package configcenter

import "github.com/im-pingo/gomodule"

type Feature interface {
	gomodule.IModule
	HelloWorld()
}
