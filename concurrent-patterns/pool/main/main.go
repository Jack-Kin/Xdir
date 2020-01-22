package main

import (
	"Xdir/concurrent-patterns/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
	展示如何使用pool包来共享使用一组模拟的数据库连接
 */

const (
	maxGoroutines = 5
	pooledResources = 2
)

// 模拟要共享的资源
type dbConnection struct {
	ID int32
}

// Close()实现了io.Closer接口，以便dbConnection可以被Pool管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// 给每一个connection分配一个unique ID
var idCounter int32

// factory method， 当Pool需要创建新资源（连接）时， 这个函数被调用
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id},nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建Pool
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 使用Pool里的连接来完成查询
	for query:=0; query < maxGoroutines; query++  {
		// 匿名函数把query复制给q，然后执行
		// 这里需要复制，不然所有的查询会共享同一个query变量
		// TODO:Why?
		go func(q int) {
			performQueries(q,p)
			wg.Done()
		}(query)
	}

	// 等所有goroutine结束
	wg.Wait()

	// 关闭Pool
	log.Println("Shutdown Program.")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	// 从pool里请求一个连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// 把连接释放回pool里
	defer p.Release(conn)

	// 这里用等待时间模拟一个查询响应的过程
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("Query: QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)

}