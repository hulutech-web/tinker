// Code generated by 'yaegi extract github.com/spf13/viper'. DO NOT EDIT.

package symbol

import (
	"github.com/spf13/viper"
	"reflect"
)

func init() {
	Symbols["github.com/spf13/viper/viper"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AddConfigPath":            reflect.ValueOf(viper.AddConfigPath),
		"AddRemoteProvider":        reflect.ValueOf(viper.AddRemoteProvider),
		"AddSecureRemoteProvider":  reflect.ValueOf(viper.AddSecureRemoteProvider),
		"AllKeys":                  reflect.ValueOf(viper.AllKeys),
		"AllSettings":              reflect.ValueOf(viper.AllSettings),
		"AllowEmptyEnv":            reflect.ValueOf(viper.AllowEmptyEnv),
		"AutomaticEnv":             reflect.ValueOf(viper.AutomaticEnv),
		"BindEnv":                  reflect.ValueOf(viper.BindEnv),
		"BindFlagValue":            reflect.ValueOf(viper.BindFlagValue),
		"BindFlagValues":           reflect.ValueOf(viper.BindFlagValues),
		"BindPFlag":                reflect.ValueOf(viper.BindPFlag),
		"BindPFlags":               reflect.ValueOf(viper.BindPFlags),
		"ConfigFileUsed":           reflect.ValueOf(viper.ConfigFileUsed),
		"Debug":                    reflect.ValueOf(viper.Debug),
		"DebugTo":                  reflect.ValueOf(viper.DebugTo),
		"DecodeHook":               reflect.ValueOf(viper.DecodeHook),
		"EnvKeyReplacer":           reflect.ValueOf(viper.EnvKeyReplacer),
		"Get":                      reflect.ValueOf(viper.Get),
		"GetBool":                  reflect.ValueOf(viper.GetBool),
		"GetDuration":              reflect.ValueOf(viper.GetDuration),
		"GetEnvPrefix":             reflect.ValueOf(viper.GetEnvPrefix),
		"GetFloat64":               reflect.ValueOf(viper.GetFloat64),
		"GetInt":                   reflect.ValueOf(viper.GetInt),
		"GetInt32":                 reflect.ValueOf(viper.GetInt32),
		"GetInt64":                 reflect.ValueOf(viper.GetInt64),
		"GetIntSlice":              reflect.ValueOf(viper.GetIntSlice),
		"GetSizeInBytes":           reflect.ValueOf(viper.GetSizeInBytes),
		"GetString":                reflect.ValueOf(viper.GetString),
		"GetStringMap":             reflect.ValueOf(viper.GetStringMap),
		"GetStringMapString":       reflect.ValueOf(viper.GetStringMapString),
		"GetStringMapStringSlice":  reflect.ValueOf(viper.GetStringMapStringSlice),
		"GetStringSlice":           reflect.ValueOf(viper.GetStringSlice),
		"GetTime":                  reflect.ValueOf(viper.GetTime),
		"GetUint":                  reflect.ValueOf(viper.GetUint),
		"GetUint16":                reflect.ValueOf(viper.GetUint16),
		"GetUint32":                reflect.ValueOf(viper.GetUint32),
		"GetUint64":                reflect.ValueOf(viper.GetUint64),
		"GetViper":                 reflect.ValueOf(viper.GetViper),
		"InConfig":                 reflect.ValueOf(viper.InConfig),
		"IniLoadOptions":           reflect.ValueOf(viper.IniLoadOptions),
		"IsSet":                    reflect.ValueOf(viper.IsSet),
		"KeyDelimiter":             reflect.ValueOf(viper.KeyDelimiter),
		"MergeConfig":              reflect.ValueOf(viper.MergeConfig),
		"MergeConfigMap":           reflect.ValueOf(viper.MergeConfigMap),
		"MergeInConfig":            reflect.ValueOf(viper.MergeInConfig),
		"MustBindEnv":              reflect.ValueOf(viper.MustBindEnv),
		"New":                      reflect.ValueOf(viper.New),
		"NewWithOptions":           reflect.ValueOf(viper.NewWithOptions),
		"OnConfigChange":           reflect.ValueOf(viper.OnConfigChange),
		"ReadConfig":               reflect.ValueOf(viper.ReadConfig),
		"ReadInConfig":             reflect.ValueOf(viper.ReadInConfig),
		"ReadRemoteConfig":         reflect.ValueOf(viper.ReadRemoteConfig),
		"RegisterAlias":            reflect.ValueOf(viper.RegisterAlias),
		"RemoteConfig":             reflect.ValueOf(&viper.RemoteConfig).Elem(),
		"Reset":                    reflect.ValueOf(viper.Reset),
		"SafeWriteConfig":          reflect.ValueOf(viper.SafeWriteConfig),
		"SafeWriteConfigAs":        reflect.ValueOf(viper.SafeWriteConfigAs),
		"Set":                      reflect.ValueOf(viper.Set),
		"SetConfigFile":            reflect.ValueOf(viper.SetConfigFile),
		"SetConfigName":            reflect.ValueOf(viper.SetConfigName),
		"SetConfigPermissions":     reflect.ValueOf(viper.SetConfigPermissions),
		"SetConfigType":            reflect.ValueOf(viper.SetConfigType),
		"SetDefault":               reflect.ValueOf(viper.SetDefault),
		"SetEnvKeyReplacer":        reflect.ValueOf(viper.SetEnvKeyReplacer),
		"SetEnvPrefix":             reflect.ValueOf(viper.SetEnvPrefix),
		"SetFs":                    reflect.ValueOf(viper.SetFs),
		"SetTypeByDefaultValue":    reflect.ValueOf(viper.SetTypeByDefaultValue),
		"Sub":                      reflect.ValueOf(viper.Sub),
		"SupportedExts":            reflect.ValueOf(&viper.SupportedExts).Elem(),
		"SupportedRemoteProviders": reflect.ValueOf(&viper.SupportedRemoteProviders).Elem(),
		"Unmarshal":                reflect.ValueOf(viper.Unmarshal),
		"UnmarshalExact":           reflect.ValueOf(viper.UnmarshalExact),
		"UnmarshalKey":             reflect.ValueOf(viper.UnmarshalKey),
		"WatchConfig":              reflect.ValueOf(viper.WatchConfig),
		"WatchRemoteConfig":        reflect.ValueOf(viper.WatchRemoteConfig),
		"WithLogger":               reflect.ValueOf(viper.WithLogger),
		"WriteConfig":              reflect.ValueOf(viper.WriteConfig),
		"WriteConfigAs":            reflect.ValueOf(viper.WriteConfigAs),

		// type definitions
		"ConfigFileAlreadyExistsError":   reflect.ValueOf((*viper.ConfigFileAlreadyExistsError)(nil)),
		"ConfigFileNotFoundError":        reflect.ValueOf((*viper.ConfigFileNotFoundError)(nil)),
		"ConfigMarshalError":             reflect.ValueOf((*viper.ConfigMarshalError)(nil)),
		"ConfigParseError":               reflect.ValueOf((*viper.ConfigParseError)(nil)),
		"DecoderConfigOption":            reflect.ValueOf((*viper.DecoderConfigOption)(nil)),
		"FlagValue":                      reflect.ValueOf((*viper.FlagValue)(nil)),
		"FlagValueSet":                   reflect.ValueOf((*viper.FlagValueSet)(nil)),
		"Logger":                         reflect.ValueOf((*viper.Logger)(nil)),
		"Option":                         reflect.ValueOf((*viper.Option)(nil)),
		"RemoteConfigError":              reflect.ValueOf((*viper.RemoteConfigError)(nil)),
		"RemoteProvider":                 reflect.ValueOf((*viper.RemoteProvider)(nil)),
		"RemoteResponse":                 reflect.ValueOf((*viper.RemoteResponse)(nil)),
		"StringReplacer":                 reflect.ValueOf((*viper.StringReplacer)(nil)),
		"UnsupportedConfigError":         reflect.ValueOf((*viper.UnsupportedConfigError)(nil)),
		"UnsupportedRemoteProviderError": reflect.ValueOf((*viper.UnsupportedRemoteProviderError)(nil)),
		"Viper":                          reflect.ValueOf((*viper.Viper)(nil)),

		// interface wrapper definitions
		"_FlagValue":      reflect.ValueOf((*_github_com_spf13_viper_FlagValue)(nil)),
		"_FlagValueSet":   reflect.ValueOf((*_github_com_spf13_viper_FlagValueSet)(nil)),
		"_Logger":         reflect.ValueOf((*_github_com_spf13_viper_Logger)(nil)),
		"_Option":         reflect.ValueOf((*_github_com_spf13_viper_Option)(nil)),
		"_RemoteProvider": reflect.ValueOf((*_github_com_spf13_viper_RemoteProvider)(nil)),
		"_StringReplacer": reflect.ValueOf((*_github_com_spf13_viper_StringReplacer)(nil)),
	}
}

// _github_com_spf13_viper_FlagValue is an interface wrapper for FlagValue type
type _github_com_spf13_viper_FlagValue struct {
	IValue       interface{}
	WHasChanged  func() bool
	WName        func() string
	WValueString func() string
	WValueType   func() string
}

func (W _github_com_spf13_viper_FlagValue) HasChanged() bool {
	return W.WHasChanged()
}
func (W _github_com_spf13_viper_FlagValue) Name() string {
	return W.WName()
}
func (W _github_com_spf13_viper_FlagValue) ValueString() string {
	return W.WValueString()
}
func (W _github_com_spf13_viper_FlagValue) ValueType() string {
	return W.WValueType()
}

// _github_com_spf13_viper_FlagValueSet is an interface wrapper for FlagValueSet type
type _github_com_spf13_viper_FlagValueSet struct {
	IValue    interface{}
	WVisitAll func(fn func(viper.FlagValue))
}

func (W _github_com_spf13_viper_FlagValueSet) VisitAll(fn func(viper.FlagValue)) {
	W.WVisitAll(fn)
}

// _github_com_spf13_viper_Logger is an interface wrapper for Logger type
type _github_com_spf13_viper_Logger struct {
	IValue interface{}
	WDebug func(msg string, keyvals ...any)
	WError func(msg string, keyvals ...any)
	WInfo  func(msg string, keyvals ...any)
	WTrace func(msg string, keyvals ...any)
	WWarn  func(msg string, keyvals ...any)
}

func (W _github_com_spf13_viper_Logger) Debug(msg string, keyvals ...any) {
	W.WDebug(msg, keyvals...)
}
func (W _github_com_spf13_viper_Logger) Error(msg string, keyvals ...any) {
	W.WError(msg, keyvals...)
}
func (W _github_com_spf13_viper_Logger) Info(msg string, keyvals ...any) {
	W.WInfo(msg, keyvals...)
}
func (W _github_com_spf13_viper_Logger) Trace(msg string, keyvals ...any) {
	W.WTrace(msg, keyvals...)
}
func (W _github_com_spf13_viper_Logger) Warn(msg string, keyvals ...any) {
	W.WWarn(msg, keyvals...)
}

// _github_com_spf13_viper_Option is an interface wrapper for Option type
type _github_com_spf13_viper_Option struct {
	IValue interface{}
}

// _github_com_spf13_viper_RemoteProvider is an interface wrapper for RemoteProvider type
type _github_com_spf13_viper_RemoteProvider struct {
	IValue         interface{}
	WEndpoint      func() string
	WPath          func() string
	WProvider      func() string
	WSecretKeyring func() string
}

func (W _github_com_spf13_viper_RemoteProvider) Endpoint() string {
	return W.WEndpoint()
}
func (W _github_com_spf13_viper_RemoteProvider) Path() string {
	return W.WPath()
}
func (W _github_com_spf13_viper_RemoteProvider) Provider() string {
	return W.WProvider()
}
func (W _github_com_spf13_viper_RemoteProvider) SecretKeyring() string {
	return W.WSecretKeyring()
}

// _github_com_spf13_viper_StringReplacer is an interface wrapper for StringReplacer type
type _github_com_spf13_viper_StringReplacer struct {
	IValue   interface{}
	WReplace func(s string) string
}

func (W _github_com_spf13_viper_StringReplacer) Replace(s string) string {
	return W.WReplace(s)
}