# OSXTUN
---------------

A demo that create tun in osx

## usage
```{shell}
$ go get github.com/chzyer/osxtun
$ osxtun
$ sudo ifconfig utun4 10.2.0.2 10.2.0.1 mtu 1450 netmask 255.255.255.255 up
$ sudo route add -net 10.2.0.0/24 -interface utun4
```
