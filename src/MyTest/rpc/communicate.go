package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
	"unsafe"
)

var zRecvCount = uint32(0) // 张大爷听到了多少句话
var lRecvCount = uint32(0) // 李大爷听到了多少句话
var total = uint32(100000) // 总共需要遇见多少次

var z0 = []byte("0吃了没，您吶?")
var z3 = []byte("3嗨！吃饱了溜溜弯儿。")
var z5 = []byte("5回头去给老太太请安！")
var l1 = []byte("1刚吃。")
var l2 = []byte("2您这，嘛去？")
var l4 = []byte("4有空家里坐坐啊。")

var liWriteLock sync.Mutex    // 李大爷的写锁
var zhangWriteLock sync.Mutex // 张大爷的写锁

var p = &sync.Pool{
	New: func() interface{} {
		return make([]byte, 64)
	},
}

type RequestResponse struct {
	Serial  uint32 // 序号
	Payload []byte // 内容
}

// 序列化RequestResponse，并发送
// 序列化后的结构如下：
//   长度  4字节
//   Serial 4字节
//   PayLoad 变长
func writeTo(serial []byte, payLoad []byte, tcpClient *TcpClient) {
	//byt := p.Get().([]byte)
	byt := make([]byte, 64)
	binary.BigEndian.PutUint32(byt, uint32(len(payLoad)+4))
	copy(byt[4:8], serial)
	copy(byt[8:], payLoad)
	tcpClient.conn.Write(byt[0:])
	//var pkg =  new(bytes.Buffer)
	//binary.Write(pkg, binary.BigEndian, int32(len(payLoad))+4)
	//binary.Write(pkg, binary.BigEndian, serial)
	//binary.Write(pkg, binary.BigEndian, payLoad)
	//fmt.Printf("发送消息大小:%d\n", len(payLoad)+8)
	//tcpClient.conn.Write(pkg.Bytes())
}

func writeIntTo(serial uint32, payLoad []byte, tcpClient *TcpClient) {
	//byt := p.Get().([]byte)
	byt := make([]byte, 64)
	binary.BigEndian.PutUint32(byt, uint32(len(payLoad)+4))
	binary.BigEndian.PutUint32(byt[4:8], serial)
	copy(byt[8:], payLoad)
	tcpClient.conn.Write(byt[0:])
	//var pkg = new(bytes.Buffer)
	//binary.Write(pkg, binary.BigEndian, int32(len(payLoad)) + 4)
	//binary.Write(pkg, binary.BigEndian, serial)
	//binary.Write(pkg, binary.BigEndian, payLoad)
	////fmt.Printf("发送消息大小:%d\n", len(payLoad)+8)
	//tcpClient.conn.Write(pkg.Bytes())
}

// 接收数据，反序列化成RequestResponse
func readFrom(tcpClient *TcpClient) ([]byte, error) {
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据，
	// 该操作不会将数据读出，只是引用，引用的数据在下一次读取操作之
	// 前是有效的。如果切片长度小于 n，则返回一个错误信息说明原因。
	// 如果 n 大于缓存的总大小，则返回 ErrBufferFull。
	lengthByte, err := tcpClient.r.Peek(4)
	length := binary.BigEndian.Uint32(lengthByte)
	//创建 Buffer缓冲器
	//lengthBuff := bytes.NewBuffer(lengthByte)
	//var length int32
	//// 通过Read接口可以将buf中得内容填充到data参数表示的数据结构中
	//err = binary.Read(lengthBuff, binary.BigEndian, &length)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Printf("读取消息大小:%d\n", length)
	// Buffered 返回缓存中未读取的数据的长度
	//if int32(c.r.Buffered()) < length+4 {
	//	fmt.Printf("读取消息大小:%d\n",length)
	//	return nil, err
	//}
	// 读取消息真正的内容
	//pack := p.Get().([]byte)
	pack := make([]byte, 64)
	// Read 从 b 中读出数据到 p 中，返回读出的字节数和遇到的错误。
	// 如果缓存不为空，则只能读出缓存中的数据，不会从底层 io.Reader
	// 中提取数据，如果缓存为空，则：
	// 1、len(p) >= 缓存大小，则跳过缓存，直接从底层 io.Reader 中读
	// 出到 p 中。
	// 2、len(p) < 缓存大小，则先将数据从底层 io.Reader 中读取到缓存
	// 中，再从缓存读取到 p 中。
	_, err = io.ReadFull(tcpClient.r, pack)
	if err != nil {
		log.Printf("fail read:%+v\n", pack)
		return nil, err
	}
	return pack[4 : length+4], nil
}

// 张大爷的耳朵
func zhangDaYeListen(tcpClient *TcpClient, wg *sync.WaitGroup) {
	defer wg.Done()
	for zRecvCount < total*3 {
		load, err := readFrom(tcpClient)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// fmt.Println("张大爷收到：" + r.Payload)
		if load[4] == '2' { // 如果收到：您这，嘛去？
			go writeTo(load[0:4], z3, tcpClient) // 回复：嗨！吃饱了溜溜弯儿。
		} else if load[4] == '4' { // 如果收到：有空家里坐坐啊。
			go writeTo(load[0:4], z5, tcpClient) // 回复：回头去给老太太请安！
		} else if load[4] == '1' { // 如果收到：刚吃。
			// 不用回复
		} else {
			fmt.Println("张大爷听不懂：" + Bytes2str(load[4:]))
			break
		}
		zRecvCount++
	}
}

// 张大爷的嘴
func zhangDaYeSay(tcpClient *TcpClient) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		writeIntTo(nextSerial, z0, tcpClient)
		nextSerial++
	}
}

// 李大爷的耳朵，实现是和张大爷类似的
func liDaYeListen(tcpClient *TcpClient, wg *sync.WaitGroup) {
	defer wg.Done()
	for lRecvCount < total*3 {
		load, err := readFrom(tcpClient)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// fmt.Println("李大爷收到：" + r.Payload)
		if load[4] == '0' { // 如果收到：吃了没，您吶?
			go writeTo(load[0:4], l1, tcpClient) // 回复：刚吃。
		} else if load[4] == '3' {
			// do nothing
		} else if load[4] == '5' {
			// do nothing
		} else {
			fmt.Println("李大爷听不懂：" + Bytes2str(load[4:]))
			break
		}
		lRecvCount++
	}
}

// 李大爷的嘴
func liDaYeSay(conn *TcpClient) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		writeIntTo(nextSerial, l2, conn)
		nextSerial++
		writeIntTo(nextSerial, l4, conn)
		nextSerial++
	}
}

func startServer(wg *sync.WaitGroup) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	fmt.Println("张大爷在胡同口等着 ...")
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		tcpClient := &TcpClient{conn: conn, r: bufio.NewReaderSize(conn, 1024*1024*32)}
		fmt.Println("碰见一个李大爷:" + conn.RemoteAddr().String())
		go zhangDaYeListen(tcpClient, wg)
		go zhangDaYeSay(tcpClient)
	}

}

func startClient(wg *sync.WaitGroup) *net.TCPConn {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	tcpClient := &TcpClient{conn: conn, r: bufio.NewReaderSize(conn, 1024*1024*32)}
	go liDaYeListen(tcpClient, wg)
	go liDaYeSay(tcpClient)
	return conn
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go startServer(&wg)
	time.Sleep(time.Second)
	conn := startClient(&wg)
	t1 := time.Now()
	wg.Wait()
	elapsed := time.Since(t1)
	conn.Close()
	fmt.Println("耗时: ", elapsed)
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type TcpClient struct {
	conn net.Conn
	r    *bufio.Reader
}

func roundup_power_of_2(v int32) int32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
