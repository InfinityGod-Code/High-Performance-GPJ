package Java.Basics.Concurrency;
import java.util.concurrent.Executors;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Future;
import java.util.stream.IntStream;


public class Main {

    public static void calculate() {
        long result = 0;

        for (long i = 0; i < 1_000_000_000L; i++) {
            result += i;
        }
        System.out.println("Result: " + result);
    }

    public static void parallel(){
        // Example of parallel execution using ExecutorService
        long start = System.currentTimeMillis();

        /*
        The Java ForkJoinPool can distribute the CPU-bound work across multiple worker threads,
        which can execute simultaneously on different CPU cores. */

        IntStream.range(0, 10).parallel().forEach(i -> {
            System.out.println("Processing item: " + i + " in thread: " + Thread.currentThread().getName());
            calculate();
        });

        long end = System.currentTimeMillis();

        System.out.println("Time: " + (end - start) + " ms");

    }
    public static void main(String[] args) {
        ExecutorService executor = Executors.newFixedThreadPool(2);
        Future<?> task1 = executor.submit(() -> {
            System.out.println("Task 1 is running");
            try{
                Thread.sleep(7000);
            } catch (InterruptedException e) {
                e.printStackTrace();
                Thread.currentThread().interrupt();
            }
        });
        Future<?> task2 = executor.submit(() -> {
            System.out.println("Task 2 is running");
            try{
                Thread.sleep(7000);
            } catch (InterruptedException e) {
                e.printStackTrace();
                Thread.currentThread().interrupt();
            }
        });
        
        // Execute tasks and wait for completion
        try {
            task1.get();
            task2.get();
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            executor.shutdown();    
        }
    }
}