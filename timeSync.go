package main

// import (
// 	"bytes"
// 	"fmt"
// 	"math/rand"
// 	"sync/atomic"

// 	"os"
// 	"sync"
// 	"time"
// )

const (
	name    = "timeSync"
	version = "0.0.1"
	desc    = "timeSync is utility tool to check node date is sync with ntp server."
)

type NTPServer struct {
	server []string
}

type NTPDate struct {
	date map[string]map[string]string
	NTPServer
}

type ExecCommand struct {
	command *string
	status  int
	output  *string
}

func (ntp *NTPDate) syncNTPServer() {
	for _, server := range ntp.server {
		mm := make(map[string]string)
		mm["currentDate"] = "serveDate"
		ntp.date[server] = mm
	}

}

func main() {

}
