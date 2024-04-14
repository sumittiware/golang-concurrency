# Dining Philosophers Problem

## Problem Statement

The Dining Philosophers Problem is a classic synchronization problem in computer science. It involves a group of philosophers seated around a circular table, with one fork placed between each pair of philosophers. Each philosopher must alternately think and eat. However, a philosopher can only eat when they have acquired both the fork to their left and the fork to their right. The problem lies in designing a discipline of behavior (a concurrent algorithm) that prevents deadlock and starvation while allowing as many philosophers as possible to eat.

The key challenges in this problem are:

1. **Deadlock**: If all philosophers pick up their left fork simultaneously, they will all be waiting indefinitely for their right fork, resulting in a deadlock situation.
2. **Starvation**: A philosopher might be perpetually denied access to the forks they need to eat, leading to starvation.

## Solution

The provided Go code solves the Dining Philosophers Problem using mutexes (mutual exclusion locks) and goroutines (lightweight threads). Here's how the solution works:

1. **Defining Structures and Constants**
   - The `Philosopher` struct represents a philosopher with their name and the indices of their left and right forks.
   - An array of `Philosopher` structs is defined, representing the philosophers seated around the table.
   - Constants `hunger`, `eatTime`, `thinkTime`, and `sleepTime` are defined to control the behavior of the philosophers.

2. **Synchronization Primitives**
   - A `sync.WaitGroup` (`wg`) is used to wait for all philosophers to finish their meals.
   - Another `sync.WaitGroup` (`seated`) is used to ensure all philosophers are seated before starting to eat.
   - A map of `sync.Mutex` (`forks`) is used to represent the forks, where each fork has a mutex to control its access.
   - A `sync.Mutex` (`orderMutex`) is used to synchronize the order in which philosophers finish their meals.

3. **Main Function**
   - The `main` function initializes the problem and calls the `dine` function to start the dining process.
   - After the dining process completes, it prints the order in which the philosophers finished their meals.

4. **Dining Process**
   - The `dine` function creates a new `sync.WaitGroup` (`wg`) and adds a counter for each philosopher.
   - It also initializes the `forks` map with mutexes for each fork.
   - Then, it starts a new goroutine for each philosopher, calling the `dinningProblem` function.
   - The `dine` function waits for all philosophers to finish their meals using `wg.Wait()`.

5. **Dining Problem for Each Philosopher**
   - The `dinningProblem` function simulates the behavior of a philosopher during the dining process.
   - It first signals that the philosopher is seated at the table and waits for all philosophers to be seated using `seated.Wait()`.
   - Then, it enters a loop based on the `hunger` constant, representing the number of times a philosopher must eat.
   - Inside the loop, the philosopher picks up the forks using the `pickUpForks` function, eats for the specified `eatTime`, thinks for the specified `thinkTime`, and then puts down the forks using the `putDownForks` function.
   - After the loop finishes, the philosopher goes to sleep for the specified `sleepTime`.
   - Finally, the philosopher's name is added to the `orderFinished` slice using a mutex (`orderMutex`) to synchronize access.

6. **Picking Up and Putting Down Forks**
   - The `pickUpForks` function acquires the mutexes for the philosopher's left and right forks in a specific order to prevent deadlock.
   - The `putDownForks` function releases the mutexes for the philosopher's left and right forks.

By using mutexes and goroutines, the solution ensures that only one philosopher can hold a fork at a time, preventing deadlock and starvation situations. The synchronization primitives (`sync.WaitGroup` and `sync.Mutex`) help coordinate the philosophers' actions and maintain the correct order of acquiring and releasing forks.

The solution also demonstrates the use of goroutines for concurrent execution, allowing each philosopher to dine concurrently while sharing the forks safely.