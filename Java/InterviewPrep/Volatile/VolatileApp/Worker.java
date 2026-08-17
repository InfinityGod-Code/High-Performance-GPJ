package Java.InterviewPrep.Volatile.VolatileApp;

public class Worker implements Runnable{
    private static boolean status = true;

    public void shutDown(){
        System.out.println("Main thread calls and finally Shut this down!!");
        status = false;
    }

    @Override
    public void run() {
        while (status) {
            System.out.println("This is the long running tasks");
        }
    }
}
