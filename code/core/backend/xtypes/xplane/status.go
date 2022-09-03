package xplane

type NodeStat struct {
	Epoch     int `json:"epoch,omitempty"`
	TotalMem  int `json:"total_mem,omitempty"`
	UsedMem   int `json:"used_mem,omitempty"`
	TotalSwap int `json:"total_swap,omitempty"`
	UsedSwap  int `json:"used_swap,omitempty"`
	CPU       int `json:"cpu,omitempty"`
	AvgLoad   int `json:"avg_load,omitempty"`
}

type NodeStatus struct {
	Id         string     `json:"id,omitempty"`
	Stats      []NodeStat `json:"stats,omitempty"`
	LastUpdate string     `json:"last_update,omitempty"`
	Tags       []string   `json:"tags,omitempty"`
}

type AppStatus struct {
	ClusterName string                `json:"cluster_name,omitempty"`
	Nodes       map[string][]NodeStat `json:"nodes,omitempty"`
}
