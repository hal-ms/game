package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/game/cnto"
	"github.com/hal-ms/game/log"
	"github.com/tarm/serial"
)

func main() {
	r := gin.Default()
	r.POST("/button", cnto.Button)
	r.POST("/is_wearing/:IsWearing", cnto.IsWearing)
	r.POST("/job/:job", cnto.Job)
	go r.Run()
	go hitScreen()

	for {
	}

}

func hitScreen() {
	c := &serial.Config{Name: "COM6", Baud: 9600}
	s, err := serial.OpenPort(c)
	buf := make([]byte, 128)
	if err != nil {
		log.SendSlack(err.Error())
		panic(err)
	}
	for {
		n, err := s.Read(buf)
		if err != nil {
			panic(err)
		}
		p, err := strconv.Atoi(string(buf[:n]))
		if err != nil {
			panic(err)
		}
		cnto.Hit(p)
	}
}
