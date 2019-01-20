# Mutex and Channel basics

### What is an atomic operation?
> An atom operation is a operation or a set of instructions that cannot be interrupted while beeing executed

### What is a semaphore?
> A semaphore is a essentially just a flag used to control access to a common resource by multiple processes in a concurrent system. The semaphore can be multi-bits which is useful if more than one process can access a resource.

### What is a mutex?
> A mutex is similar to a binary semaphore and used to synchronize concurrent processes.

### What is the difference between a mutex and a binary semaphore?
> A mutex is a locking mechanism while a semaphore is a signaling mechanism. A process that locks the mutex must be the one who unlocks it, this is not the case with semaphores. 

### What is a critical section?
> A critical section is a section in a process which must not be interrupted because it accesses shared resources.

### What is the difference between race conditions and data races?
 > A data race happens when two or more threads access the same memory location and at least one of the accesses if for writing. (A memory location is beeing written to and read from at the same time). A race condition on the other hand happens when more threads access shared data and try to change it at the same time.

### List some advantages of using message passing over lock-based synchronization primitives.
> message passing is more predictable since the order of operations between processes are controlled by the programmer, making the synchronization very transparent and making sure all processes get access to the shared resources. A lock-based synchronization on the other hand can result in dead-locks if the process holding the lock is dependent on another process to finnish, in which the other process can not finnish until the lock is opened. 

### List some advantages of using lock-based synchronization primitives over message passing.
> A lock-based synchronization can be faster, since it does not care about the synchronization of the processes, but only that the critical sections are not interrupted.
