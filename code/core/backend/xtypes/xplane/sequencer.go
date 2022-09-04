package xplane

import "github.com/bwmarrin/snowflake"

// Sequencer is responsible for generating id (int64)
// for various propose using snowflake

type Sequencer interface {
	DeviceId() int64
	SessionId() int64

	EventId() string
	RequestId() string
	DisplayErrId() string

	NewNode() *snowflake.Node
}
