package service

// SessMan is responsible for generating id (int64)
// for various propose using snowflake
type SessMan interface {
	DeviceId() int64
	SessionId() int64

	EventId() string
	RequestId() string
	DisplayErrId() string
}
