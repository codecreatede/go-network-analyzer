# go-network-devops-systemd

- write the tcp,udp, 4tcp and 6tcp files 
- timevalues and automatic backup to the destination drive.
- It will automatically backup all your tcp, udp, 4tcp and 6tcp connections and also the port connections to the containers to the given drive or the user defined folder. 

```
➜  go-network-devops-systemd git:(main) ✗ go run main.go -h
Provide the tcp, udp, 4tcp, 6tcp or the port for the check for the running containers ip address

Usage:
  options [flags]

Flags:
  -h, --help          help for options
  -t, --tcp string    tcp as a flag (default "provide the tcp")
  -f, --tcp4 string   tcp4 as a flag (default "provide the tcp4")
  -s, --tcp6 string   tcp6 as a flag (default "provide the tcp")
  -u, --udp string    udp as a flag (default "provide the udp")
```
Gaurav Sablok
