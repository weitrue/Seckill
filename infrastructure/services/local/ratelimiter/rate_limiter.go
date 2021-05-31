/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午5:57
 * Description: 有两种经典模式实现限流器：Fan-In 和 Fan-Out，也叫扇入和扇出。
				Fan-In 模式
				Fan-In 模式是一种多对一的模式，即有多个生产者，只有一个消费者。相当于限流算法中的漏桶算法，处理速度取决于消费者的速度。

				这种模式最大的作用是将生产者的数据排队，在消费端以固定速度消费并处理，它也比较适合需要在消费端操作分布式锁的场景。
				在秒杀系统中控制队列的消费速度，并通过分布式锁和原子操作扣减库存。

				在Go语言中，Channel可以当队列来用，而生产者和消费者就是Goroutine。
				对于Fan-In模式来说，它可以看作是多个Goroutine生产数据到一个Channel，由一个Goroutine来消费Channel。


				Fan-Out 模式
				Fan-Out是一种一对多的模式，即一个生产者，有多个消费者。它类似于限流算法中的令牌桶算法，处理速度取决于生产者的速度。

				这种模式最大的作用是在源头通过限制生产者的速度，来控制下游系统的压力。
				由于下游系统偏重业务，涉及 Redis、mysql 等操作，通常每次请求会有一定的延迟。
				为了提升吞吐量，通常是采用多个Goroutine的来消费队列中的数据，配合连接池并发请求下游系统。

				在高并发系统中，通常会将两者搭配使用，这种方式也叫组合模式。
				秒杀系统中，用户的抢购请求在获得初步资格后，先通过Fan-In进入队列，由一个Goroutine消费队列，
				并通过Fan-Out将请求转发给多个Goroutine来并发操作多个商品的库存。
 **/

package ratelimiter

import (
	"github.com/weitrue/Seckill/infrastructure/pool/worker/taski"
)

type RateLimiter interface {
	Push(t taski.Task) bool  // 提供给生产者推送数据
	Pop() (taski.Task, bool) // 提供给消费者消费数据
	Close() error                // 关闭限流器
}
