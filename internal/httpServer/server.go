package httpServer

import (
	"fmt"
	"github.com/call-me-snake/date_service/internal/model"
	"github.com/gorilla/mux"
	"net/http"
)

type Connector struct {
	router  *mux.Router
	address string
}

//New - Конструктор *Connector
func New(addr string) *Connector {
	c := &Connector{}
	c.router = mux.NewRouter()
	c.address = addr
	return c
}

func (c *Connector) executeHandlers(serverTime model.Itime) {
	c.router.HandleFunc("/time/now",timeNow(serverTime)).Methods("GET")
	c.router.HandleFunc("/time/string",timeToString()).Methods("GET")
	c.router.HandleFunc("/time/add",addTime()).Methods("GET")
	c.router.HandleFunc("/time/correct",correctTime(serverTime)).Methods("POST")
}

func (c *Connector) Start(serverTime model.Itime) error {
	c.executeHandlers(serverTime)
	err := http.ListenAndServe(c.address, c.router)
	return fmt.Errorf("httpServer.Start: %v", err)
}
