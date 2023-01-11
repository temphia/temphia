package xplane

import "github.com/bwmarrin/snowflake"

// Sequencer is responsible for generating id (int64)
// for various propose using snowflake

type IDService interface {
	DeviceNode() *snowflake.Node
	SessionNode() *snowflake.Node
	EventNode() *snowflake.Node
	RequestNode() *snowflake.Node
	DisplayErrNode() *snowflake.Node

	NewNode(key string) *snowflake.Node
}
