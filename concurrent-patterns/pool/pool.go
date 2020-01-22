package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

/*
	展示如何使用一个有缓冲的通道实现资源池，
	这个资源池可以管理在任意多个goroutine之间共享以及独立使用的资源，
	比如网络连接、数据库连接等。
	这种模式在需要共享一组静态资源的情况下非常有用。

	每个goroutine可以从资源池里申请资源，
	使用完之后再放回资源池里，以便其他goroutine复用。
 */


type Pool struct {
	// 互斥锁
	m 				sync.Mutex
	// io.Closer接口的通道
	resources 		chan io.Closer
	// 当需要一个新的资源时，通过factory函数创建，具体实现由使用者提供
	factory 		func() (io.Closer, error)
	// 表示资源池是否关闭
	closed 			bool

	// io.Closer定义如下：
	//type Closer interface {
	//    Close() error
	//}
	// 用于关闭数据流，释放资源
}

// 定义一个资源池已经被关闭的错误
var ErrPoolClosed = errors.New("Pool has been closed.")

// New函数创建资源池
// 接收两个参数：1. 创建新资源的函数fn 2. 资源池的大小size
// 返回两个参数：1. 指向资源池的指针   2. error
func New(fn func() (io.Closer, error), size uint )(*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}
	return &Pool{
		//m:         sync.Mutex{},
		resources: make(chan io.Closer, size),
		factory:   fn,
		//closed:    false,
	}, nil
}

// Acquire方法从资源池获取资源， 如果没有资源就调用factory方法生成一个返回
// 函数名称前面的括号是Go定义这些函数将在其上运行的对象的方式。
func (p *Pool) Acquire()(io.Closer, error) {
	select {
	// 通道接收多参返回， 第一参数是接收的值， 第二参数表示通道是否关闭
	case r, ok := <- p.resources:
		log.Println("Acquire:", "共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "生成新资源")
		return p.factory()
	}
}

// 将一个使用后的资源放回资源池里
func (p *Pool) Release(r io.Closer) {
	// 和Close一样， 阻止这两个方法在不同goroutine里同时运行
	p.m.Lock()
	defer p.m.Unlock()
	
	if p.closed {
		_ = r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "资源释放到Pool里了")
	default:
		log.Println("Release:", "Pool满了， 释放这个资源吧")
		_ = r.Close()
	}
}


// 关闭资源池
func (p* Pool) Close() {
	// 必须保证同一时刻只有一个goroutine执行Close/Release方法
	// closed字段要同步
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)
	for r:= range p.resources {
		_ = r.Close()
	}
}
