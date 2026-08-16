package Java.InterviewPrep.Fundamentals.InteractiveEditor;

import java.util.Queue;
import java.util.concurrent.ConcurrentLinkedQueue;

// Separate class responsible for managing user messages safely across threads
public class DocumentBuffer {
    private final Queue<String> messageQueue = new ConcurrentLinkedQueue<>();

    // Main thread calls this when user types a message
    public void addMessage(String message) {
        messageQueue.add(message);
    }

    // Daemon thread calls this to drain and collect unsaved messages
    public String getNextUnsavedMessage() {
        return messageQueue.poll(); // Returns null if queue is empty
    }

    public boolean hasUnsavedMessages() {
        return !messageQueue.isEmpty();
    }
}