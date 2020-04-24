// Code generated by Lingo for table information_schema.PLUGINS - DO NOT EDIT

package qplugins

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QPlugins {
	return instance
}

func PluginName() path.String {
	return instance.pluginName
}

func PluginVersion() path.String {
	return instance.pluginVersion
}

func PluginStatus() path.String {
	return instance.pluginStatus
}

func PluginType() path.String {
	return instance.pluginType
}

func PluginTypeVersion() path.String {
	return instance.pluginTypeVersion
}

func PluginLibrary() path.String {
	return instance.pluginLibrary
}

func PluginLibraryVersion() path.String {
	return instance.pluginLibraryVersion
}

func PluginAuthor() path.String {
	return instance.pluginAuthor
}

func PluginDescription() path.String {
	return instance.pluginDescription
}

func PluginLicense() path.String {
	return instance.pluginLicense
}

func LoadOption() path.String {
	return instance.loadOption
}
