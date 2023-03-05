package sharedcli

import "github.com/alecthomas/kong"

type Context struct {
	BprintFile  string
	ConfigFile  string
	KongContext *kong.Context
}
