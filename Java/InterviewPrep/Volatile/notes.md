## Volatile
In Java, the volatile keyword guarantees memory visibility and instruction ordering across threads.

#### Question 1 : What does volatile guarantee?

Core MechanismsForces
 **Main Memory Access**: Normally, threads cache shared variables in local CPU registers or L1/L2 caches for speed. Marking a variable as volatile forces all reads and writes directly to and from main system memory, ensuring every thread reads the most recently updated value.  
 **Prevents Instruction Reordering**: The compiler and CPU often reorder instructions to optimize performance. For a volatile variable, Java establishes a Happens-Before relationship, preventing the JVM from reordering memory access around it. 

 Key Notes : In the Java Memory Model (JMM), the Happens-Before relationship is a formal consistency guarantee. It ensures that memory writes made by one thread are guaranteed to be visible to reads performed by another thread, preventing state corruption caused by CPU and compiler instruction reordering.

If operation A happens-before operation B:

**Visibility**: All memory modifications made up to operation A are guaranteed to be fully visible to operation B.

**Ordering**: The JVM guarantees that operation A will be completed before operation B runs, or will at least appear to have done so.

**Note** : So in order to understand this we can take keep in mind that Threads wheneven access variables in sharing we if we want to change that shared variable and wants that other Thread also sees that we can make that variable volatile.
Volatile doesn't guarantees atomicity and it is non-blocking as comapared to synchorized and atomic classes.

#### Question 2 : Does volatile guarantee mutual exclusion?

No, volatile does not guarantee mutual exclusion. It ensures memory visibility and instruction ordering, but it never blocks threads or prevents multiple threads from accessing a variable at the exact same time.

#### Question 4 : What is a happens-before relationship?

In the Java Memory Model (JMM), the Happens-Before relationship is a formal consistency guarantee. It ensures that memory writes made by one thread are guaranteed to be visible to reads performed by another thread, preventing state corruption caused by CPU and compiler instruction reordering.

