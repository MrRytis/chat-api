package response

type WebsocketMessage struct {
	Action string       `json:"action"`
	Data   GroupMessage `json:"data"`
}
