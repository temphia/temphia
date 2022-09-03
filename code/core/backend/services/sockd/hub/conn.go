package sockdhub

import (
	"fmt"

	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
)

type PlugConnOptions struct {
	TenantId string
	UserId   string
	GroupId  string
	DeviceId string
	Plug     string
	Conn     sockdx.Conn
}

type UserConnOptions struct {
	TenantId string
	UserId   string
	GroupId  string
	DeviceId string
	Conn     sockdx.Conn
}

type UpdateDynRoomTagsOptions struct {
	TenantId  string
	DynSource string
	DynGroup  string
	ConnId    int64
}

type DevConnOptions struct {
	TenantId string
	UserId   string
	PlugId   string
	AgentId  string
	Conn     sockdx.Conn
}

func (s *SockdHub) AddPlugConn(opts PlugConnOptions) error {
	return s.sockd.NewConnection(sockdx.ConnOptions{
		NameSpace: opts.TenantId,
		Conn:      opts.Conn,
		Expiry:    0,
		Room:      ROOM_PLUG_DEV,
		Tags:      []string{fmt.Sprintf("plug_%s", opts.Plug)},
	})
}

func (s *SockdHub) AddDevConn(opts DevConnOptions) error {
	return s.sockd.NewConnection(sockdx.ConnOptions{
		NameSpace: opts.TenantId,
		Conn:      opts.Conn,
		Expiry:    0,

		Room: ROOM_PLUG_DEV,
		Tags: []string{

			fmt.Sprintf("plug_%s", opts.PlugId),
			fmt.Sprint("sys.user_", opts.UserId),
			"dev_user",
		},
	})
}

func (s *SockdHub) AddUserConn(opts UserConnOptions) error {

	return s.sockd.NewConnection(sockdx.ConnOptions{
		NameSpace: opts.TenantId,
		Conn:      opts.Conn,
		Expiry:    0,
		Room:      ROOM_SYS_USERS,
		Tags: []string{
			fmt.Sprint("sys.user_", opts.UserId),
			fmt.Sprint("sys.ugroup_", opts.GroupId),
			fmt.Sprint("sys.device_", opts.DeviceId),
			TAG_REALUSER,
		},
	})

}

func (s *SockdHub) UpdateDynRoomTags(opts UpdateDynRoomTagsOptions) error {

	return s.sockd.RoomUpdateTags(
		opts.TenantId,
		ROOM_SYSTABLE,
		sockdx.UpdateTagOptions{
			AddTags:    []string{fmt.Sprintf("dgroup.%s.%s", opts.DynSource, opts.DynGroup)},
			ClearOld:   true,
			RemoveTags: []string{},
			Id:         int64(opts.ConnId),
		})
}

func (s *SockdHub) UpdateRoomTags(tenantId, room string, opts *sockdx.UpdateTagOptions) error {
	return s.sockd.RoomUpdateTags(
		tenantId,
		room,
		sockdx.UpdateTagOptions{
			Id:         opts.Id,
			AddTags:    opts.AddTags,
			RemoveTags: opts.RemoveTags,
			ClearOld:   opts.ClearOld,
		},
	)
}
