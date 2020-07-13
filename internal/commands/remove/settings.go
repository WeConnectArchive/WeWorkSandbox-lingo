package remove

import (
	"github.com/spf13/viper"
)

func getSettings() *settings {
	s := &settings{}
	s.rootDirectory = viper.GetString(flagDir)
	s.recursive = viper.GetBool(flagRecursive)
	s.verbose = viper.GetBool(flagVerbose)
	return s
}

type settings struct {
	rootDirectory string
	recursive     bool
	verbose       bool
}

func (s settings) RootDirectory() string {
	return s.rootDirectory
}

func (s settings) Recursive() bool {
	return s.recursive
}

func (s settings) Verbose() bool {
	return s.verbose
}
