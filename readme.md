# Go Concurrency and Parallelism

Go is a statically typed, compiled programming language designed with concurrency and parallelism in mind. Unlike many other languages, Go provides built-in support for concurrency through goroutines and channels, making it easier to write efficient and scalable concurrent programs.

## Goroutines and Channels

In Go, concurrency is achieved through the use of goroutines and channels. Goroutines are lightweight, independently executing functions that run concurrently with other goroutines in the same address space. They are similar to threads but are more lightweight, efficient, and easier to create and manage.

Channels, on the other hand, are the mechanism used for communication between goroutines. They provide a safe way for goroutines to send and receive data without the need for explicit locks or condition variables. Channels help prevent common concurrency issues like race conditions and deadlocks.

## Parallelism

While concurrency is about dealing with multiple tasks in an interleaved manner, parallelism is about executing multiple tasks simultaneously. Go supports parallelism through the use of goroutines that can be scheduled across multiple logical processors (logical CPUs or threads).

Go's runtime automatically multiplexes goroutines onto the available logical processors, allowing programs to take advantage of parallel execution on multi-core systems without the need for explicit thread management.

## Classic Concurrency Problems

This repository contains three classic concurrency problems solved in Go, demonstrating the usage of goroutines, channels, and other synchronization primitives. Each problem is contained in a separate folder with its own README file explaining the problem and the solution implemented.

1. **[Producer-Consumer Problem](producer-consumer/README.md)**
   - This problem involves coordinating the production and consumption of data items in a shared buffer.
   - The solution demonstrates the use of goroutines and channels for safe communication between producers and consumers.

2. **[Dining Philosophers Problem](dining-philosophers/README.md)**
   - This problem revolves around a group of philosophers who must coordinate their actions to avoid deadlock and starvation while dining.
   - The solution utilizes mutexes and goroutines to ensure proper synchronization and prevent deadlock and starvation.

3. **[Sleeping Barber Problem](sleeping-barber/README.md)**
   - This problem involves managing a barbershop with limited waiting room capacity and coordinating between the barber and customers.
   - The solution employs goroutines and channels to simulate the arrival of customers, manage the waiting room capacity, and ensure proper synchronization between the barber and customers.

