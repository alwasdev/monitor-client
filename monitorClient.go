package monitor_client

import (
	"time"
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"bytes"
)

func MonitorClient(serverUrl string, applicationName string) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				pushSnapshot(serverUrl + "/monitor/push", applicationName)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
func pushSnapshot(serverUrl string, applicationName string) {
	u := MonitorSnapshot{application:applicationName, usedMemory:10, freeMemory:20, commitedMemory:10, maxMemory:300, activeThreads:10, timestamp:time.Now().UTC()}
	log.Println("Sending MonitorSnapshot", u)
	b, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Sending MonitorSnapshot", b)
	res, _ := http.Post(serverUrl, "application/json; charset=utf-8", bytes.NewReader(b))
	fmt.Println("Pusher", res)
}

type MonitorSnapshot struct {
	application    string `json:"application"`
	usedMemory     int64 `json:"usedMemory"`
	freeMemory     int64 `json:"freeMemory"`
	commitedMemory int64 `json:"commitedMemory"`
	maxMemory      int64 `json:"maxMemory"`
	activeThreads  int `json:"activeThreads"`
	timestamp      time.Time `json:"timestamp"`
}