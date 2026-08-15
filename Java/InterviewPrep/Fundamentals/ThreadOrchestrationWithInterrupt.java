package Java.InterviewPrep.Fundamentals;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

// Task 1: Implements Runnable for downloading a report
class ReportFetcherTask implements Runnable {
    @Override
    public void run() {
        try {
            System.out.println("[ReportFetcher] Starting report download...");
            Thread.sleep(1500); // Simulate work
            System.out.println("[ReportFetcher] Report downloaded successfully.");
        } catch (InterruptedException e) {
            System.out.println("[ReportFetcher] INTERRUPTED! Cleaning up downloaded temp files...");
            // Restore interrupted status if needed
            Thread.currentThread().interrupt();
        }
    }
}

// Task 2: Long-running task that handles interruptions gracefully
class HeavyDataProcessorTask implements Runnable {
    @Override
    public void run() {
        System.out.println("[DataProcessor] Starting long calculation loop...");
        try {
            for (int i = 1; i <= 10; i++) {
                System.out.println("[DataProcessor] Processing batch " + i + "/10...");
                
                // Thread.sleep checks for interrupts and throws InterruptedException
                Thread.sleep(1000); 
            }
            System.out.println("[DataProcessor] All batches completed.");
        } catch (InterruptedException e) {
            System.out.println("[DataProcessor] INTERRUPTED during batch processing! Halting execution safely.");
            Thread.currentThread().interrupt();
        }
    }
}

// Task 3: Implements Runnable for sending emails
class EmailNotifierTask implements Runnable {
    @Override
    public void run() {
        try {
            System.out.println("[EmailSender] Dispatching notification email...");
            Thread.sleep(800); // Simulate network delay
            System.out.println("[EmailSender] Email dispatched successfully.");
        } catch (InterruptedException e) {
            System.out.println("[EmailSender] INTERRUPTED! Email delivery canceled.");
            Thread.currentThread().interrupt();
        }
    }
}

public class ThreadOrchestrationWithInterrupt {

    public static void main(String[] args) {
        System.out.println("[Main Thread] Starting task orchestration...\n");

        ExecutorService executor = Executors.newFixedThreadPool(3);

        // Instantiate the separate Runnable task classes
        Runnable fetchTask = new ReportFetcherTask();
        Runnable processTask = new HeavyDataProcessorTask();
        Runnable emailTask = new EmailNotifierTask();

        // Submit tasks and keep the Future handles to control/monitor them
        Future<?> fetchFuture = executor.submit(fetchTask);
        Future<?> processFuture = executor.submit(processTask);
        Future<?> emailFuture = executor.submit(emailTask);

        try {
            // Let tasks run for 2.5 seconds
            Thread.sleep(2500);

            // Main thread decides HeavyDataProcessorTask is taking too long and cancels it.
            // passing 'true' sends an interrupt (Thread.interrupt()) to the running thread.
            System.out.println("\n[Main Thread] DataProcessor is taking too long! Interrupting Task 2...");
            boolean cancelled = processFuture.cancel(true); 
            System.out.println("[Main Thread] Cancel request sent to DataProcessor. Success: " + cancelled + "\n");

            // Wait for the remaining tasks to complete normally
            fetchFuture.get(); 
            emailFuture.get();

        } catch (Exception e) {
            System.err.println("[Main Thread] Exception caught: " + e.getMessage());
        } finally {
            System.out.println("[Main Thread] Shutting down executor service.");
            executor.shutdown();
        }

        System.out.println("[Main Thread] Orchestration finished.");
    }
}