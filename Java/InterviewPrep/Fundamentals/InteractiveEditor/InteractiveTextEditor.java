package Java.InterviewPrep.Fundamentals.InteractiveEditor;

import java.util.Scanner;



public class InteractiveTextEditor {
    public static void main(String[] args) {
        DocumentBuffer buffer = new DocumentBuffer();
        String draftFilePath = "autosave_draft.txt";

        // 1. Create and configure background daemon thread
        Thread daemonThread = new Thread(new AutoSaver(buffer, draftFilePath));
        daemonThread.setDaemon(true); // Ensures thread exits when main finishes
        daemonThread.start();

        System.out.println("==================================================");
        System.out.println(" Interactive Document Editor");
        System.out.println(" Type your text below and press Enter.");
        System.out.println(" Background Daemon will auto-save every 3 seconds.");
        System.out.println(" Type 'EXIT' to close the application.");
        System.out.println("==================================================\n");

        Scanner scanner = new Scanner(System.in);

        // 2. Main thread loop: interactive console input
        while (true) {
            System.out.print("> ");
            String input = scanner.nextLine();

            if ("EXIT".equalsIgnoreCase(input.trim())) {
                System.out.println("\n[Main Thread] Exiting application...");
                break;
            }

            // Store message in the thread-safe buffer
            buffer.addMessage(input);
        }

        scanner.close();
        System.out.println("[Main Thread] App closed. JVM terminating (Daemon thread stopped automatically).");
    }
}
