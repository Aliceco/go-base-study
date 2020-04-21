package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	// 读取文件
	file, _ := os.OpenFile("./test/test", os.O_RDONLY, 666)
	defer file.Close()
	fw:=bufio.NewReader(file)
	for i:=1;i<=2;i++ {
		go func(index int) {
			defer wg.Done()
			for {
				mutex.Lock() // 加锁
				str, err:=fw.ReadString('\n') // 读到换行就是一行
				if err!=nil {
					if err==io.EOF { //读取到结尾
						mutex.Unlock() // 解锁
						break
					}
					fmt.Println(err)
				}
				time.Sleep(time.Microsecond*200)
				fmt.Printf("[携程%d]：%s", index, str)
				mutex.Unlock() // 解锁
			}
		}(i)
	}
	wg.Add(2)
	wg.Wait()
	fmt.Println("读取完成")
}
