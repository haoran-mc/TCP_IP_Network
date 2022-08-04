## 第 13 章 多种 I/O 函数

因为 Go 将套接字封装了起来，所以统一使用 `conn.Write()` 与 `conn.Read()`。

### MSG_OOB：发送紧急消息