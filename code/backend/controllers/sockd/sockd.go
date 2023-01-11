package sockd

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
)

type Controller struct {
	sockd sockdx.SockdCore
}

func New(sockd sockdx.SockdCore) *Controller {
	return &Controller{
		sockd: sockd,
	}
}

func (s *Controller) AddPlugConn(opts PlugConnOptions) error {
	return s.sockd.NewConnection(sockdx.ConnOptions{
		NameSpace: opts.TenantId,
		Conn:      opts.Conn,
		Expiry:    0,
		Room:      sockdx.ROOM_PLUG_DEV,
		Tags:      []string{fmt.Sprintf("plug_%s", opts.Plug)},
	})
}

func (s *Controller) AddDevConn(opts DevConnOptions) error {
	return s.sockd.NewConnection(sockdx.ConnOptions{
		NameSpace: opts.TenantId,
		Conn:      opts.Conn,
		Expiry:    0,

		Room: sockdx.ROOM_PLUG_DEV,
		Tags: []string{

			fmt.Sprintf("plug_%s", opts.PlugId),
			fmt.Sprint("sys.user_", opts.UserId),
			"dev_user",
		},
	})
}

func (s *Controller) AddUserConn(opts UserConnOptions) error {

	return s.sockd.NewConnection(sockdx.ConnOptions{
		NameSpace: opts.TenantId,
		Conn:      opts.Conn,
		Expiry:    0,
		Room:      sockdx.ROOM_SYS_USERS,
		Tags: []string{
			fmt.Sprint("sys.user_", opts.UserId),
			fmt.Sprint("sys.ugroup_", opts.GroupId),
			fmt.Sprint("sys.device_", opts.DeviceId),
			sockdx.TAG_REALUSER,
		},
	})

}

func (s *Controller) AddData(opts DataConnOptions) error {

	return s.sockd.NewConnection(sockdx.ConnOptions{
		NameSpace: opts.TenantId,
		Conn:      opts.Conn,
		Expiry:    0,
		Room:      sockdx.ROOM_SYS_USERS,
		Tags: []string{
			fmt.Sprint("sys.user_", opts.UserId),
			fmt.Sprintf("dgroup.%s.%s", opts.DynSource, opts.DynGroup),
		},
	})

}

func (s *Controller) UpdateDynRoomTags(opts UpdateDynRoomTagsOptions) error {

	return s.sockd.RoomUpdateTags(
		opts.TenantId,
		sockdx.ROOM_SYSTABLE,
		sockdx.UpdateTagOptions{
			AddTags:    []string{fmt.Sprintf("dgroup.%s.%s", opts.DynSource, opts.DynGroup)},
			ClearOld:   true,
			RemoveTags: []string{},
			Id:         int64(opts.ConnId),
		})
}

func (s *Controller) UpdateRoomTags(tenantId, room string, opts *sockdx.UpdateTagOptions) error {
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
