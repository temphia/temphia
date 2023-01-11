package idservice

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

type IDService struct {
	nodes  map[string]*snowflake.Node
	nodeId int64
	mlock  sync.Mutex
}

func New(id int64) *IDService {

	return &IDService{
		nodes:  make(map[string]*snowflake.Node),
		nodeId: id,
	}
}

func (s *IDService) DeviceNode() *snowflake.Node     { return s.NewNode("temphia.device") }
func (s *IDService) SessionNode() *snowflake.Node    { return s.NewNode("temphia.session") }
func (s *IDService) EventNode() *snowflake.Node      { return s.NewNode("temphia.event") }
func (s *IDService) RequestNode() *snowflake.Node    { return s.NewNode("temphia.request") }
func (s *IDService) DisplayErrNode() *snowflake.Node { return s.NewNode("temphia.display_err") }

func (s *IDService) NewNode(key string) *snowflake.Node {
	s.mlock.Lock()
	defer s.mlock.Unlock()

	n := s.nodes[key]
	if n == nil {
		nn, err := snowflake.NewNode(s.nodeId)
		if err != nil {
			panic(err)
		}
		s.nodes[key] = nn
		n = nn
	}

	return n
}
