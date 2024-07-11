# ab test

`ab -n 50000 -c 20000 http://localhost:8080/ping`

```bash

cpu 4 cores;

# ulimit -a

max memory size         (kbytes, -m) unlimited
open files                      (-n) 65535
pipe size            (512 bytes, -p) 8
POSIX message queues     (bytes, -q) 819200
real-time priority              (-r) 0
stack size              (kbytes, -s) 8192
cpu time               (seconds, -t) unlimited
max user processes              (-u) 14988
```

nginx.conf

```conf
worker_processes  auto;
events {
    worker_connections  65535;
}
```

## 均衡算法

1. 默认 轮询

```bash
pid = 24 2718 connection
pid = 25 3029 connection
pid = 26 2551 connection
pid = 27 3336 connection
```

2. hash $remote_addr;

```bash
pid = 23 2497 connection
pid = 24 2114 connection
pid = 25 2242 connection
pid = 26 1967 connection
```
