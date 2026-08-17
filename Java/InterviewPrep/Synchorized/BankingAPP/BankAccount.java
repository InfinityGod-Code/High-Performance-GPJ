package Java.InterviewPrep.Synchorized.BankingAPP;

// This is basically we trying to simulate the real transactions that happen in a bank account.
// Where we might have multiple concurrent users trying to perform transactions
//  on the same bank account at the same time. Demonstrate the use of synchronized keyword in Java to
// avoid race conditions and ensure thread safety while performing transactions.
public class BankAccount {

    // for the demo purpose, we will keep the balance as an integer.
    // In real-world scenarios, it would be better to use BigDecimal for monetary
    // values.
    private int balance;

    public BankAccount(int balance) {
        this.balance = balance;
    }

    public synchronized void deposit(int amount) {
        balance += amount;
        System.out.println("Deposited: " + amount + ", New Balance: " + balance);
    }

    public synchronized void withdraw(int amount) {
        if (balance >= amount) {
            System.out.println(
                    Thread.currentThread().getName() +
                            " checked balance: " + balance);

            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
            balance -= amount;

            System.out.println(
                    Thread.currentThread().getName() +
                            " withdrew: " + amount +
                            ", balance: " + balance);

            System.out.println("Withdrew: " + amount + ", New Balance: " + balance);
        } else {
            System.out.println("Insufficient funds for withdrawal of: " + amount);
        }
    }

    public int getBalance() {
        return balance;
    }
}
