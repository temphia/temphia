package coredb

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DB) AddUserMessage(msg *entities.UserMessage) (int64, error) {
	rid, err := d.userMessagesTable().Insert(msg)
	if err != nil {
		return 0, err
	}

	id := rid.ID().(int64)

	return id, nil
}

func (d *DB) UserMessageSetRead(tenantId, user string, id int64) error {
	return d.userMessagesTable().Find(
		db.Cond{
			"tenant_id": tenantId,
			"user_id":   user,
			"id":        id},
	).Update(db.Cond{
		"read": true,
	})
}

func (d *DB) RemoveUserMessage(tenantId string, userId string, id int64) error {
	return d.userMessagesTable().Find(db.Cond{"tenant_id": tenantId, "user_id": userId, "id": id}).Delete()
}

func (d *DB) ListUserMessages(tenantId string, data *entities.UserMessageReq) ([]*entities.UserMessage, error) {
	messages := make([]*entities.UserMessage, 0)
	cond := db.Cond{
		"tenant_id": tenantId,
		"user_id":   data.UserId,
	}

	// fixme => lookup paginate/cursor

	if data.Cursor != 0 {
		cond["id >"] = data.Cursor
	}

	result := d.userMessagesTable().Find(cond)

	var err error
	if data.Count != 0 {
		err = result.Limit(int(data.Count)).All(&messages)
	} else {
		err = result.All(&messages)
	}

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (d *DB) ReadUserMessages(tenantId, userId string, id []int64) error {
	return d.userMessagesTable().Find(db.Cond{
		"tenant_id": tenantId,
		"user_id":   userId,
		"id IN":     id,
	}).Update(db.Cond{
		"read": true,
	})
}

func (d *DB) DeleteUserMessages(tenantId, userId string, id []int64) error {
	return d.userMessagesTable().Find(db.Cond{
		"tenant_id": tenantId,
		"user_id":   userId,
		"id IN":     id,
	}).Delete()
}
