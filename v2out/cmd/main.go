package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

const buf_size=8192
func main() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("client 建立连接")
		//新起独立连接到 远端
		server, _ := net.Dial("tcp", "127.0.0.1:8082")

		go handleClientRequest(client, server)

		go handleClientResp(client, server)
	}
}

func handleClientRequest(client net.Conn, desc_client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	defer desc_client.Close()

	buf := make([]byte, buf_size)

	for {
		n, err := client.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		data := buf[:n]

		encode_data := base64.StdEncoding.EncodeToString(data)

		log.Println(len(encode_data))

		//压缩成二进制包
		first := uint32(1)
		//第二个字节
		second := uint32(len(encode_data))
		buf_02 := make([]byte, 8)
		binary.BigEndian.PutUint32(buf_02[0:], first)
		binary.BigEndian.PutUint32(buf_02[4:], second)

		send_data := string(buf_02) + encode_data
		//建立隧道，加包

		desc_client.Write([]byte(send_data))

	}

}

//解析包内容， 写回浏览器
func handleClientResp(client net.Conn, desc_client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	defer desc_client.Close()

	for {
		var stream_buf bytes.Buffer

		for {
			buf := make([]byte, buf_size)

			n, err := desc_client.Read(buf)
			if err != nil {
				log.Println(err)
				goto RESULT
			}
			data := buf[:n]

			//打到流buf中
			stream_buf.Write(data)

			for {
				//获取全部数据
				all_data := stream_buf.Bytes()
				if len(all_data) > 8 {
					header := all_data[:8]

					data_len := binary.BigEndian.Uint32(header[4:])

					real_buf := all_data[8:]

					if len(real_buf) >= int(data_len) {
						encode_data := real_buf[:int(data_len)]
						data , _:= base64.StdEncoding.DecodeString(string(encode_data))
						client.Write(data)

						log.Println(data_len)

						stream_buf.Read(all_data[:(8+int(data_len))])
					} else {
						break
					}
				} else {
					break
				}

			}

		}

	}
RESULT:


}
