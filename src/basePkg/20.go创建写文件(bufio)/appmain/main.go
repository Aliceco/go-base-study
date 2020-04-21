package main

import (
	"bufio"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	//var name bytes.Buffer
	//name.Write([]byte("an"))
	//name.WriteTo(os.Stdout) // an

	// 创建文件
	//file, _ := os.OpenFile("./test/test", os.O_RDONLY | os.O_CREATE | os.O_TRUNC, 666)
	////defer file.Close() // 最后执行 defer先进后出
	//file.Write([]byte("an"))
	//file.Close()

	// 创建文件并追加内容
	file, _ := os.OpenFile("./test/test", os.O_RDWR | os.O_CREATE | os.O_APPEND | os.O_TRUNC, 666)
	defer file.Close()
	fw:=bufio.NewWriterSize(file, 3)
	for i:=0;i<10;i++ {
		fw.Write([]byte("ans\n")) // 达不到缓存区大小
		//fw.Flush() // 写入内容
		// time.Sleep(time.Second*2)
	}

}
