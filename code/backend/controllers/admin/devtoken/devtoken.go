package devtoken

import (
	"encoding/base64"

	"github.com/ugorji/go/codec"
)

type Plug struct {
	HostAddrs []string `codec:"host_addrs,omitempty" json:"host_addrs,omitempty"`
	TenantId  string   `codec:"tenant_id,omitempty" json:"tenant_id,omitempty"`
	DevTicket string   `codec:"dev_ticket,omitempty" json:"dev_ticket,omitempty"`
}

func (pd *Plug) ToString() (string, error) {

	b := make([]byte, 0, 100)

	encoder := codec.NewEncoderBytes(&b, &codec.JsonHandle{})

	err := encoder.Encode(pd)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func NewPlug(rawstr string) (*Plug, error) {

	raw, err := base64.StdEncoding.DecodeString(rawstr)
	if err != nil {
		return nil, err
	}

	pd := &Plug{}

	decoder := codec.NewDecoderBytes(raw, &codec.JsonHandle{})

	err = decoder.Decode(pd)
	if err != nil {
		return nil, err
	}

	return pd, nil
}
