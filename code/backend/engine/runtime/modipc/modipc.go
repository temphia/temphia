package modipc

import (
	"fmt"
	"reflect"

	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes"
)

type ModIPC struct {
	innerMod reflect.Value
}

func NewModIPC(mod any) *ModIPC {
	return &ModIPC{
		innerMod: reflect.ValueOf(mod),
	}
}

func (m *ModIPC) Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error) {

	mFunc := m.innerMod.MethodByName(method)

	fargType := mFunc.Type().In(0).Elem()

	fargElem := reflect.New(fargType).Elem()

	err := args.AsObject(fargElem.Addr().Interface())
	if err != nil {
		return nil, err
	}

	result := mFunc.Call([]reflect.Value{fargElem.Addr()})

	switch len(result) {
	case 1:
		if !result[0].IsNil() {
			return nil, result[0].Interface().(error)
		}
		return lazydata.NewAnyData(nil), nil
	case 2:
		if !result[1].IsNil() {
			return nil, result[0].Interface().(error)
		}
		return lazydata.NewReflectData(result[0]), nil
	default:
		panic(fmt.Sprintf("wrong func signature for  method %s", method))
	}
}
