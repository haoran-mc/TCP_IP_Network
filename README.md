# 《TCP/IP网络编程》 学习笔记

## 仓库介绍

为了学习 Go 语言中的网络编程实现，将本书 C 语言实现的功能用 Go 语言重新实现一遍。

其中 C 代码由 [riba2534/TCP-IP-NetworkNote](https://github.com/riba2534/TCP-IP-NetworkNote) 仓库提供。

## 目录

| 目录                        | C                 | Go                 |
| --------------------------- | ----------------- | ------------------ |
| 理解网络编程和套接字        | [ch01](./ch01/c/) | [ch01](./ch01/go/) |
| 套接字类型与协议设置        | [ch02](./ch02/c/) | [ch02](./ch02/go/) |
| 地址族与数据序列            | [ch03](./ch03/c/) | [ch03](./ch03/go/) |
| 基于TCP的服务端/客户端（1） | [ch04](./ch04/c/) | [ch04](./ch04/go/) |
| 基于TCP的服务端/客户端（2） | [ch05](./ch05/c/) | [ch05](./ch05/go/) |
| 基于UDP的服务端/客户端      | [ch06](./ch06/c/) | [ch06](./ch06/go/) |
| 优雅地断开套接字的连接      | [ch07](./ch07/c/) | [ch07](./ch07/go/) |
| 域名及网络地址              | [ch08](./ch08/c/) | [ch08](./ch08/go/) |
| 套接字的多种可选项          | [ch09](./ch09/c/) | [ch09](./ch09/go/) |
| 多进程服务器端              | [ch10](./ch10/c/) | [ch10](./ch10/go/) |
| 进程间通信                  | [ch11](./ch11/c/) | [ch11](./ch11/go/) |
| I/O复用                     | [ch12](./ch12/c/) | [ch12](./ch12/go/) |
| 多种I/O函数                 | [ch13](./ch13/c/) | [ch13](./ch13/go/) |
| 多播与广播                  | [ch14](./ch14/c/) | [ch14](./ch14/go/) |
| 套接字和标准I/O             | [ch15](./ch15/c/) | [ch15](./ch15/go/) |
| 关于I/O流分离的其他内容     | [ch16](./ch16/c/) | [ch16](./ch16/go/) |
| 优于select的epoll           | [ch17](./ch17/c/) | [ch17](./ch17/go/) |
| 多线程服务器端的实现        | [ch18](./ch18/c/) | [ch18](./ch18/go/) |
| 制作HTTP服务器端            | [ch24](./ch24/c/) | [ch24](./ch24/go/) |

```shell
sudo apt install git-lfs
git lfs install
git lfs track "*.pdf"
git add .gitattributes
git add TCP\&IP网络编程.pdf
git commit -m ""
git push
```