package internal

import (
	"log"
	"time"

	"github.com/sch8ill/mclib"
	"github.com/sch8ill/mclib/slp"
)

var PlayerCounts = make(map[string]int)

func getServerInfo(ip string) (*slp.Response, bool) {
	client, _ := mclib.NewClient(ip, mclib.WithTimeout(time.Millisecond*900))
	res, _ := client.StatusPing()
	return res, res != nil
}

func updateServer(ip string) {
	res, ok := getServerInfo(ip)

	if ok {
		PlayerCounts[ip] = res.Players.Online
	}
}

func Update() {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for t := range ticker.C {
			for _, server := range Servers {
				go updateServer(server)
			}

			time.Sleep(time.Millisecond * 950)

			if isDebug() {
				log.Println("Player counts updated:", t.Format(time.TimeOnly))

				for s, i := range PlayerCounts {
					log.Printf(" - %s %d\n", s, i)
				}
			}
		}
	}()
}
