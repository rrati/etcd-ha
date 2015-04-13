package haclient

import (
	"fmt"
	"time"

        "github.com/coreos/go-etcd/etcd"
)

var expires uint64
func Go(name string) {
	key := "leader"
	expires = 10
	client := etcd.NewClient([]string {"http://localhost:4001"})
	for {
		if resp, err := client.Get(key, true, false); err != nil {
			fmt.Println("Did not find a leader.  Assuming leadership")
			if _, err := client.Create(key, name, expires); err != nil {
				fmt.Println("Failed to become leader")
				fmt.Println(err)
			}
		} else {
			fmt.Println("Found leader: ", resp.Node.Value)
			if resp.Node.Value == name {
				fmt.Println("Updating ttl")
				if _, err := client.Update(key, name, expires); err != nil {
					fmt.Println(err)
				}
			}
		}
		time.Sleep(time.Second*5)
	}
}
