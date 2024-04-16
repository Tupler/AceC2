package Server

const (
	STATUS_RUN = iota
	STATUS_STOP
)

const (
	LISTEN_TCP = iota
	LISTEN_HTTP
	LISTEN_HTTPS
	LISTEN_WEBSOCKET
)
