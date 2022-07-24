## 第二章 套接字类型与协议设置

### 创建套接字

Go 语言中已经将创建套接字的操作封装了起来，我们直接调用 `net.Listen` 监听指定的 ip:port 就可以了。如果觉得 Go 库使用起来有限制的话，还可以用 `syscall.Socket` 的方式调用，实际上 Go 库本身也是利用的 `syscall.Socket`，这里不再介绍。

```go
import "net"

func Listen(net, laddr string) (Listener, error)
/*
用于侦听tcp、unix（stream）等协议，返回一个Listener接口
net：表示传入的网络协议
laddr：（local address）表示本地的IP地址及端口
*/
```

### 支持的网络协议

- tcp：代表 TCP 协议，其基于的 IP 协议的版本根据参数 address 的值自适应；
- tcp4：代表基于 IP 协议第四版的 TCP 协议；
- tcp6：代表基于 IP 协议第六版的 TCP 协议；
- udp：代表 UDP 协议，其基于的 IP 协议的版本根据参数 address 的值自适应；
- udp4：代表基于 IP 协议第四版的 UDP 协议；
- udp6：代表基于 IP 协议第六版的 UDP 协议；
- unix：代表 Unix 通信域下的一种内部 socket 协议，以 SOCK_STREAM 为 socket 类型；
- unixgram：代表 Unix 通信域下的一种内部 socket 协议，以 SOCK_DGRAM 为 socket 类型；
- unixpacket：代表 Unix 通信域下的一种内部 socket 协议，以 SOCK_SEQPACKET 为 socket 类型；