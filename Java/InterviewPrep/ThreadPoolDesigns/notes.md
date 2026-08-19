## Thread Pool Design
In Java, a Thread Pool is a design where you create a fixed or controlled number of worker threads and reuse them to execute many tasks, instead of creating a new thread for every request.
A Java thread pool is a concurrency mechanism that maintains a reusable set of worker threads and executes submitted tasks through them, reducing thread-creation overhead and controlling concurrency.

#### 1.What is the Thread Pool Design ? 
Imagine your Spring Boot API receives : 1000 HTTP requests. Now we don't want : 
```
1000 requests
   ↓
1000 threads
```
Instead its more beneficial to have like this : 
```
1000 requests
      ↓
   Task Queue
      ↓
┌─────────────────┐
│ Thread Pool     │
│                 │
│ Thread 1        │
│ Thread 2        │
│ ...             │
│ Thread 20       │
└─────────────────┘
```
Only 20 tasks execute concurrently; the rest wait in the queue. Here is the simple java implementation for the Thread pool service : 
```
ExecutorService executor = Executors.newFixedThreadPool(5);

for (int i = 1; i <= 20; i++) {

    int taskId = i;

    executor.submit(() -> {
        System.out.println(
            "Executing task " + taskId +
            " on " + Thread.currentThread().getName()
        );
    });
}

executor.shutdown();
```

#### Question 2. Suppose we have 1000 HTTP requests and 20 DB connection pools. How we can manage this in effectively without using external services like Redis and perform on the application side.

With 1,000 requests competing for 20 database connections, 20 requests will execute queries while 980 wait in the HikariCP pool queue.Fast DB queries (e.g., 5ms Execution Time) mean 1 DB connection can serve 200 requests per second.Across 20 connections, your application can process 4,000 DB queries per second (20 \times 200).

#### 3.How does the Hikari Pool can help use in the above ? 
Every time an application talks to a database without a pool, it must open a new connection. This involves a slow multi-step process:
- Starting a TCP connection network handshake.
- Authenticating the user with database credentials.
- Setting up a user session on the database server.

A pool like HikariCP opens a set number of connections right when the application starts up. When a user needs data, they grab an existing connection instantly, use it, and return it back to the pool. This saves time and keeps the database server from crashing due to too many new connection requests.