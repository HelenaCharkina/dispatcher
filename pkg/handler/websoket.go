package handler

import (
	"dispatcher/pkg/settings"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) wshandler(w http.ResponseWriter, r *http.Request) {
	var wsupgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == fmt.Sprintf("http://%s:%s", settings.Config.ClientHost, settings.Config.ClientPort)
		},
	}

	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorln("Failed to set websocket upgrade: ", err)
		return
	}

	for {
		select {
		case msg := <-h.wsChan:
			err = conn.WriteJSON(msg)
			if err != nil {
				logrus.Errorln("Failed to write to websocket: ", err)
				return
			}
		}
	}
}
