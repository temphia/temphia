package plane

import (
	"github.com/bwmarrin/snowflake"
	"github.com/rs/xid"
)

type Sequencer struct {
	deviceNode *snowflake.Node
	sessNode   *snowflake.Node
	globalNode *snowflake.Node
	nodeId     int64
}

func NewSeq(id int64) Sequencer {

	deviceNode, err := snowflake.NewNode(id)
	if err != nil {
		panic(err)
	}

	sessNode, err := snowflake.NewNode(id)
	if err != nil {
		panic(err)
	}

	gNode, err := snowflake.NewNode(id)
	if err != nil {
		panic(err)
	}

	return Sequencer{
		deviceNode: deviceNode,
		sessNode:   sessNode,
		globalNode: gNode,
		nodeId:     id,
	}
}

func (s *Sequencer) DeviceId() int64      { return s.deviceNode.Generate().Int64() }
func (s *Sequencer) SessionId() int64     { return s.sessNode.Generate().Int64() }
func (s *Sequencer) EventId() string      { return xid.New().String() }
func (s *Sequencer) RequestId() string    { return xid.New().String() }
func (s *Sequencer) DisplayErrId() string { return xid.New().String() }
func (s *Sequencer) NewGlobalId() int64   { return s.globalNode.Generate().Int64() }
func (s *Sequencer) NewNode() *snowflake.Node {
	node, err := snowflake.NewNode(s.nodeId)
	if err != nil {
		panic(err)
	}
	return node
}
