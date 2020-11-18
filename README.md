#### 使用

可以使用go mod方式导入到项目中，代码中加入import "github.com/itchin/bytes-packet"，之后使用go mod tidy命令。

```
package main

import "github.com/itchin/bytes-packet"

func main() {
    // 打包一个消息
    buf := bpacket.Packet([]byte("this is message"))

    // 解包一个消息
    buf = bpacket.Unpack(buf)
}

```


#### 解决方案
将一个完整消息分为消息头和消息体，消息头共4个字节保存一个正整数为消息体长度。接收一段新的数据后，先从消息头读取长度，如果消息体长度小于声明的长度，将消息体写入缓存变量，继续向后读取，直到读取完毕。

代码非常简短，一目了然，详见example。
