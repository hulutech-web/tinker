// Code generated by 'yaegi extract github.com/mitchellh/mapstructure'. DO NOT EDIT.

package symbols

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
)

func init() {
	Symbols["github.com/mitchellh/mapstructure/mapstructure"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ComposeDecodeHookFunc":        reflect.ValueOf(mapstructure.ComposeDecodeHookFunc),
		"Decode":                       reflect.ValueOf(mapstructure.Decode),
		"DecodeHookExec":               reflect.ValueOf(mapstructure.DecodeHookExec),
		"DecodeMetadata":               reflect.ValueOf(mapstructure.DecodeMetadata),
		"NewDecoder":                   reflect.ValueOf(mapstructure.NewDecoder),
		"OrComposeDecodeHookFunc":      reflect.ValueOf(mapstructure.OrComposeDecodeHookFunc),
		"RecursiveStructToMapHookFunc": reflect.ValueOf(mapstructure.RecursiveStructToMapHookFunc),
		"StringToIPHookFunc":           reflect.ValueOf(mapstructure.StringToIPHookFunc),
		"StringToIPNetHookFunc":        reflect.ValueOf(mapstructure.StringToIPNetHookFunc),
		"StringToSliceHookFunc":        reflect.ValueOf(mapstructure.StringToSliceHookFunc),
		"StringToTimeDurationHookFunc": reflect.ValueOf(mapstructure.StringToTimeDurationHookFunc),
		"StringToTimeHookFunc":         reflect.ValueOf(mapstructure.StringToTimeHookFunc),
		"TextUnmarshallerHookFunc":     reflect.ValueOf(mapstructure.TextUnmarshallerHookFunc),
		"WeakDecode":                   reflect.ValueOf(mapstructure.WeakDecode),
		"WeakDecodeMetadata":           reflect.ValueOf(mapstructure.WeakDecodeMetadata),
		"WeaklyTypedHook":              reflect.ValueOf(mapstructure.WeaklyTypedHook),

		// type definitions
		"DecodeHookFunc":      reflect.ValueOf((*mapstructure.DecodeHookFunc)(nil)),
		"DecodeHookFuncKind":  reflect.ValueOf((*mapstructure.DecodeHookFuncKind)(nil)),
		"DecodeHookFuncType":  reflect.ValueOf((*mapstructure.DecodeHookFuncType)(nil)),
		"DecodeHookFuncValue": reflect.ValueOf((*mapstructure.DecodeHookFuncValue)(nil)),
		"Decoder":             reflect.ValueOf((*mapstructure.Decoder)(nil)),
		"DecoderConfig":       reflect.ValueOf((*mapstructure.DecoderConfig)(nil)),
		"Error":               reflect.ValueOf((*mapstructure.Error)(nil)),
		"Metadata":            reflect.ValueOf((*mapstructure.Metadata)(nil)),

		// interface wrapper definitions
		"_DecodeHookFunc": reflect.ValueOf((*_github_com_mitchellh_mapstructure_DecodeHookFunc)(nil)),
	}
}

// _github_com_mitchellh_mapstructure_DecodeHookFunc is an interface wrapper for DecodeHookFunc type
type _github_com_mitchellh_mapstructure_DecodeHookFunc struct {
	IValue interface{}
}
