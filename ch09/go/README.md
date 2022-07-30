## 第 9 章 套接字的多种可选项

TODO

- Golang 网络编程 API
- 套接字可选项
- Nagle 算法

```go
conn.SetNoDelay(false) // 如果打开这行代码，则禁用TCP_NODELAY，打开Nagle算法
```