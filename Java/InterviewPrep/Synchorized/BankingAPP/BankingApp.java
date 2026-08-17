package Java.InterviewPrep.Synchorized.BankingAPP;


/*
In order to see the differene we can remove the synchronized keyword from the deposit
and withdraw methods in the BankAccount class and run the BankingApp.
You will notice that the output will be inconsistent and may show incorrect balances
 */
public class BankingApp {
    public static void main(String[] args) {
        // This section will basically orchestrate the whole banking application
        // Different users can perform transactions on the same bank account concurrently
        BankAccount account = new BankAccount(1000);
        Thread thread1 = new Thread(() -> {
            
            account.deposit(500);
            
        }, "User 1 started transaction of by demositing the amount of 500");

        // User 2 will try to withdraw 700 from the account
        Thread thread2 = new Thread(() -> {
            account.withdraw(700);
        }, "User 2 started transaction of by withdrawing the amount of 700");

        // User 3 will try to withdraw 300 from the account
        Thread thread3 = new Thread(() -> {
            account.withdraw(300);
        }, "User 3 started transaction of by withdrawing the amount of 300");   

        // Start all threads
        thread1.start();
        thread2.start();
        thread3.start();
    }
}


