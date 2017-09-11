package vueutil

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

func PropsData(vm *vue.ViewModel) *js.Object {
	obj := js.Global.Get("Object").New()
	keys := js.Keys(vm.Get("$options").Get("propsData"))
	for _, value := range keys {
		obj.Set(value, vm.Get(value))
	}
	return obj
}

func GetTemplateById(ElementId string) string {
	return js.Global.Get("document").Call("getElementById", ElementId).Get("innerText").String()
}
