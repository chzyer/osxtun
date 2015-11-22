package main

import "gopkg.in/logex.v1"

// sudo ifconfig utun4 10.2.0.2 10.2.0.1 mtu 1450 netmask 255.255.255.255 up
// sudo route add -net 10.2.0.0/24 -interface utun4

func main() {
	f, err := OpenUTun()
	if err != nil {
		logex.Error(err)
		return
	}
	buf := make([]byte, 10)
	for {
		_, err := f.Read(buf)
		if err != nil {
			break
		}
		logex.Info(buf)
	}
}
