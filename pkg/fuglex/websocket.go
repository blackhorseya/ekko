package fuglex

import (
	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/gorilla/websocket"
)

// Websocket is a wrapper around the websocket connection to the Fugle API.
type Websocket struct {
	*websocket.Conn
}

// NewWebsocket returns a new Websocket connection to the Fugle API.
func NewWebsocket() (*Websocket, error) {
	conn, resp, err := websocket.DefaultDialer.Dial(configx.C.Fugle.Websocket.Endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &Websocket{
		Conn: conn,
	}, nil
}

// Auth sends an authentication message to the Fugle API.
func (w *Websocket) Auth() error {
	return w.Conn.WriteJSON(map[string]any{
		"event": "auth",
		"data": map[string]string{
			"apikey": configx.C.Fugle.APIKey,
		},
	})
}

// Ping sends a ping message to the Fugle API.
func (w *Websocket) Ping() error {
	return w.Conn.WriteJSON(map[string]string{
		"event": "ping",
	})
}
