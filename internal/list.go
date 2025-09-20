package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"
)

const url = "https://gist.githubusercontent.com/TonimatasDEV/5ae290f13b45b05e2192ae2ceb8bc4b5/raw/minecraft-servers"

var Servers []string

func addServerList() ([]string, bool) {
	gisUrl := fmt.Sprintf("%s?t=%d", url, time.Now().UnixNano())
	resp, err := http.Get(gisUrl)

	if err != nil {
		log.Println("Error getting the server list:", err)
		return nil, false
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error al close body:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Response not OK:", resp.Status)
		return nil, false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false
	}

	var servers []string
	for _, s := range strings.Split(string(body), "\n") {
		servers = append(servers, s)
	}

	return servers, true
}

func UpdateServerList() {
	servers, ok := addServerList()
	if ok {
		log.Println("Servers list updated correctly.")
		Servers = servers
	} else {
		log.Fatalln("Error updating servers list.")
	}

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for t := range ticker.C {
			servers, ok := addServerList()
			if ok {
				var newServers []string
				for _, server := range servers {
					if !slices.Contains(Servers, server) {
						newServers = append(newServers, server)
					}
				}

				var deletedServers []string
				for _, server := range Servers {
					if !slices.Contains(servers, server) {
						deletedServers = append(deletedServers, server)
					}
				}

				if len(newServers) > 0 && len(deletedServers) > 0 {
					log.Println("Server list updated:", t.Format(time.TimeOnly))

					if len(newServers) > 0 {
						log.Println(" - Added: \n", strings.Join(newServers, "\n"))
					}

					if len(deletedServers) > 0 {
						log.Println(" - Removed: \n", strings.Join(deletedServers, "\n"))
					}
				}

				Servers = servers
			}
		}
	}()

}
