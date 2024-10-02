# go-network-analyzer

- write the tcp,udp, 4tcp and 6tcp files stamped at time-date.
- create all your tcp, udp, 4tcp and 6tcp connections for any host/container.
- the files generated with the go-network-analyzer can be analyzer using my [tcp-analyzer](https://github.com/codecreatede/tcp-analyzer)

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
➜  go-network-devops-systemd git:(main) ✗ go run main.go -s tcp6
➜  go-network-devops-systemd git:(main) ✗ ls -la
total 11060
drwxr-xr-x. 1 gauravsablok gauravsablok     254 Oct  1 22:01  .
drwxr-xr-x. 1 gauravsablok gauravsablok    1616 Oct  1 14:04  ..
drwxr-xr-x. 1 gauravsablok gauravsablok     164 Oct  1 22:00  .git
-rw-r--r--. 1 gauravsablok gauravsablok     210 Oct  1 14:05  go.mod
-rw-r--r--. 1 gauravsablok gauravsablok     896 Oct  1 14:05  go.sum
-rw-r--r--. 1 gauravsablok gauravsablok    2225 Oct  1 21:59  main.go
-rwxr-xr-x. 1 gauravsablok gauravsablok 5648569 Oct  1 22:00  networkdrive
-rw-r--r--. 1 gauravsablok gauravsablok 5652480 Oct  1 22:00  networkdrive.tar
-rw-r--r--. 1 gauravsablok gauravsablok     843 Oct  1 21:49  README.md
-rw-r--r--. 1 gauravsablok gauravsablok    1245 Oct  1 22:01 'tcp6file.txt2024-10-01 22:01:16.000454544 +0200 CEST m=+0.156402181'

```
Gaurav Sablok
