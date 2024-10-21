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
pid = 24 3303 connection
pid = 25 3450 connection
pid = 26 3379 connection
pid = 27 3322 connection
```

2. hash $remote_addr;

```bash
pid = 24 5363 connection
pid = 25 5315 connection
pid = 26 2447 connection
pid = 27 5082 connection
```

## Jmter

一台压力机的 Jmeter 默认最大支持 1000 左右的并发用户数（线程数）
