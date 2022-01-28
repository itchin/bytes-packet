#### 思路
加入一个固定长度的消息头，共4个字节，保存一个正整数表示消息体长度。接收一段新的数据后，先从消息头读取长度，如果消息体长度小于声明的长度，将消息体写入缓存变量，继续向后读取，直到读取完毕。
#### 使用

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
