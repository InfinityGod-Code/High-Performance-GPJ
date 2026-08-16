## Synchronization
In Java, synchronization is a mechanism used to control access to shared resources when multiple threads execute concurrently.The main goal is to prevent race conditions and maintain data consistency.

Suppose we have Counter object and two Thread wants to increment the variable at the same time, so it might happens both enter into the same race conditions.

```
class Counter {
    int count = 0;

    void increment() {
        count++;
    }
}
```

so in order to avoid that we can use the synchronized blocks, methods or at the class level. As of now just take a note that this will prevent multiple threads to get into the race conditions.

```
public void increment() {
    synchronized (this) {
        count++;
    }
}
```

#### 1. How does synchronized work?
Suppose we have the Counter object. 

```
class Counter {
    private int count = 0;

    public synchronized void increment() {
        count++;
    }
}
```
Here Java associates the Monitor with the object. The synchronized will provide the following two mechanisms : 
- Mutual exclusion : Only one thread can own the same monitor at a time.
- Memory visibility : Synchronization also establishes a happens-before relationship. When Thread A releases the monitor  after and Thread B subsequently acquires the same monitor, Thread B is guaranteed to see the write.

<p align="center">
  <img src="../diagrams/sync.png" alt="Alt" width="50%" height="50%">
</p>

#### 2. What is an intrinsic monitor?

