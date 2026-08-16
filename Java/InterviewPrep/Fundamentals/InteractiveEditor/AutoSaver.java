package Java.InterviewPrep.Fundamentals.InteractiveEditor;

import java.io.FileWriter;
import java.io.IOException;
import java.io.PrintWriter;

public class AutoSaver implements Runnable {
    private final DocumentBuffer buffer;
    private final String filePath;

    public AutoSaver(DocumentBuffer buffer, String filePath) {
        this.buffer = buffer;
        this.filePath = filePath;
    }

    @Override
    public void run() {
        // Runs continuously in the background
        while (true) {
            try {
                // Sleep for 3 seconds between auto-saves
                Thread.sleep(3000);

                if (buffer.hasUnsavedMessages()) {
                    System.out.println("\n[AutoSave Daemon] 💾 Saving new messages to background draft...");

                    // Append unsaved messages to file
                    try (PrintWriter writer = new PrintWriter(new FileWriter(filePath, true))) {
                        while (buffer.hasUnsavedMessages()) {
                            String msg = buffer.getNextUnsavedMessage();
                            writer.println(msg);
                        }
                    } catch (IOException e) {
                        System.err.println("[AutoSave Daemon] Error saving file: " + e.getMessage());
                    }

                    System.out.println("[AutoSave Daemon] ✅ Draft updated successfully!\n> ");
                }
            } catch (InterruptedException e) {
                // Interrupted when shutting down
                break;
            }
        }
    }
}