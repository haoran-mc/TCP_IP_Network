# 《TCP/IP网络编程》 Go 语言实现

## 仓库介绍

为了学习 Go 网络编程，将本书 C 语言实现的功能用 Go 语言重新实现一遍。

其中 C 代码由 [riba2534/TCP-IP-NetworkNote](https://github.com/riba2534/TCP-IP-NetworkNote) 仓库提供。

<!-- ## 目录 -->

| 目录 | C                                        | Go                                 |
| ---- | ---------------------------------------- | ---------------------------------- |
| ch01 | [理解网络编程和套接字](./ch01/c/)        | [ch01](./ch01/go/)                 |
| ch02 | [套接字类型与协议设置](./ch02/c/)        | [ch02](./ch02/go/)                 |
| ch03 | [地址族与数据序列](./ch03/c/)            | [ch03](./ch03/go/)                 |
| ch04 | [基于TCP的服务端/客户端（1）](./ch04/c/) | [ch04](./ch04/go/)                 |
| ch05 | [基于TCP的服务端/客户端（2）](./ch05/c/) | [ch05](./ch05/go/)                 |
| ch06 | [基于UDP的服务端/客户端](./ch06/c/)      | [ch06](./ch06/go/)                 |
| ch07 | [优雅地断开套接字的连接](./ch07/c/)      | [ch07](./ch07/go/)                 |
| ch08 | [域名及网络地址](./ch08/c/)              | [ch08](./ch08/go/)                 |
| ch09 | [套接字的多种可选项](./ch09/c/)          | [ch09](./ch09/go/)                 |
| ch10 | [多进程服务器端](./ch10/c/)              | [ch10](./ch10/go/)                 |
| ch11 | [进程间通信](./ch11/c/)                  | [ch11](./ch11/go/)                 |
| ch12 | [I/O复用](./ch12/c/)                     | [ch12](./ch12/go/)                 |
| ch13 | [多种I/O函数](./ch13/c/)                 | [ch13](./ch13/go/)                 |
| ch14 | [多播与广播](./ch14/c/)                  | [ch14](./ch14/go/)                 |
| ch15 | [套接字和标准I/O](./ch15/c/)             | [ch15](./ch15/go/)                 |
| ch16 | [关于I/O流分离的其他内容](./ch16/c/)     | [ch16](./ch16/go/)                 |
| ch17 | [优于select的epoll](./ch17/c/)           | [ch17](./ch17/go/)                 |
| ch18 | [多线程服务器端的实现](./ch18/c/)        | [多协程服务器端的实现](./ch18/go/) |
| ch24 | [制作HTTP服务器端](./ch24/c/)            | [ch24](./ch24/go/)                 |

## 添加《TCP/IP网络编程》PDF

文件尺寸大于 50MB，超过 Github 上传限制，使用 git-lfs 上传。

```shell
sudo apt install git-lfs
git lfs install
git lfs track "*.pdf"
git add .gitattributes
git add TCP\&IP网络编程.pdf
git commit -m ""
git push
```