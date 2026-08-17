package Java.InterviewPrep.Volatile.VolatileApp;


public class VolatileApp {

    // This is the main Thread orchestrating other Thread. So, from this main Thread we can showdown other that 
    // shares the variable via volatile.
    public static void main(String[] args) throws InterruptedException {
        Worker worker = new Worker();
        Thread workerThread  = new Thread(worker);

        // start this worker Thread
        workerThread.start();

        Thread.sleep(2000);

        // stopping the worker thread from the main Thread.
        worker.shutDown();
    }
}
