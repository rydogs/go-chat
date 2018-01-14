package web

import (
	"net/http"
	"io/ioutil"
	"fmt"

 	"github.com/fanout/go-gripcontrol"
	"io"
)

type webSocket struct{}

func (ws *webSocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Sec-WebSocket-Extensions", "grip; message-prefix=\"\"")
	w.Header().Set("Content-Type", "application/websocket-events")
	body, _ := ioutil.ReadAll(r.Body)
	inEvents, err := gripcontrol.DecodeWebSocketEvents(string(body))
	if err != nil {
		panic("Failed to decode WebSocket events: " + err.Error())
	}
	for i, e := range inEvents {
		fmt.Printf("Event %d - type: %s, content: %s\n", i, e.Type, e.Content)
	}

	if inEvents[0].Type == "OPEN" {
		// Create the WebSocket control message:
		wsControlMessage, err := gripcontrol.WebSocketControlMessage("subscribe",
			map[string]interface{}{"channel": "test-channel"})
		if err != nil {
			panic("Unable to create control message: " + err.Error())
		}

		// Open the WebSocket and subscribe it to a channel:
		outEvents := []*gripcontrol.WebSocketEvent{
			&gripcontrol.WebSocketEvent{Type: "OPEN"},
			&gripcontrol.WebSocketEvent{Type: "TEXT",
				Content: "c:" + wsControlMessage}}
		io.WriteString(w, gripcontrol.EncodeWebSocketEvents(outEvents))
	}
}
