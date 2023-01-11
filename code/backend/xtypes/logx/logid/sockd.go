package logid

const (
	SockdNewConnection = "sockd_new_conn"

	SockdSendDirect      = "sockd_send_direct"
	SockdSendDirectBatch = "sockd_send_direct_batch"
	SockdSendBroadcast   = "sockd_send_broadcast"
	SockdSendTagged      = "sockd_send_tagged"
	SockdSendMarshelErr  = "sockd_marshel_err"

	SockdRoomUpdateTags = "sockd_update_tags"

	SockdMsgReceived        = "sockd_msg_received"
	SockdMsgReceivedDebug   = "sockd_msg_received_debug"
	SockdMsgInvalidId       = "sockd_msg_invalid_id"
	SockdMsgInvalidMType    = "sockd_msg_invalid_mtype"
	SockdMsgEmptyTargetIds  = "sockd_msg_empty_target_ids"
	SockdMsgEmptyTargetTags = "sockd_msg_empty_target_tags"

	SockdWriterStarting = "sockd_writer_starting"
	SockdWriterClosing  = "sockd_writer_closing"
	SockdWriteErr       = "sockd_write_err"
	SockdReaderStarting = "sockd_reader_starting"
	SockdReaderClosing  = "sockd_reader_closing"
	SockdReadErr        = "sockd_read_err"
	SockdConnClosed     = "sockd_conn_closed"
)
