package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Application 应用程序结构体
type Application struct {
	pool *redis.Pool
}

func main() {
	app := Application{
		pool: initPool("localhost:6379", 16, 0, 300*time.Second),
	}
	// 取出一根连接池
	conn := app.pool.Get()
	// 延时关闭
	defer conn.Close()

	// 测试redis驱动是否成功
	conn.Do("Set", "test", "测试")
	res, err := redis.String(conn.Do("Get", "test"))
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	fmt.Println("res:", res)
}
