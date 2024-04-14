# Producer-Consumer Problem and Solution with Concurrency in Go

## Problem Statement

The producer-consumer problem is a classic concurrency problem in computer science. It involves two types of processes: producers and consumers. Producers generate data items and place them in a shared buffer, while consumers consume the data items from the buffer. The problem arises when coordinating the access to the shared buffer to ensure that producers don't overwrite data items before they are consumed, and consumers don't attempt to consume data items that haven't been produced yet.

In the context of this Go program, the problem is to simulate a pizza shop where orders are received (produced) and pizzas are made (consumed) concurrently. The program needs to ensure that orders are not lost or overwritten, and that pizzas are made in the correct order.

## Solution

The provided Go program solves the producer-consumer problem using concurrency and channels. Here's how it works:

1. **Defining Constants and Variables**
   - `NumberOfPizzas` is a constant that represents the maximum number of pizza orders to be processed.
   - `pizzasMade`, `pizzasFailed`, and `total` are variables to keep track of the number of successfully made pizzas, failed pizzas, and total orders received, respectively.

2. **Defining Structures**
   - `Producer` is a struct that contains two channels: `data` for sending pizza orders and `quit` for signaling the producer to stop.
   - `PizzaOrder` is a struct that represents a pizza order with fields for the pizza number, a message (e.g., "Pizza is burnt"), and a boolean indicating whether the pizza was successfully made or not.

3. **Defining Functions**
   - `makePizza` is a function that simulates the process of making a pizza. It generates a random delay and a random result (success, burnt, or undercooked) for each pizza order.
   - `pizzazone` is a goroutine that acts as the producer. It continuously generates pizza orders and sends them through the `data` channel of the `Producer` struct.
   - `Close` is a method on the `Producer` struct that allows the main goroutine to signal the producer to stop and wait for it to finish.

4. **Main Function**
   - The `main` function initializes the `Producer` struct and starts the `pizzazone` goroutine.
   - It then enters a loop that receives pizza orders from the `data` channel of the `Producer` struct.
   - For each received order, the program checks the result (success or failure) and prints a colored message accordingly.
   - If the maximum number of orders (`NumberOfPizzas`) is reached, the program signals the producer to stop using the `Close` method and exits the loop.
   - Finally, the program prints a summary of the day's performance based on the number of successful and failed pizza orders.

By using channels and goroutines, the program achieves concurrency and solves the producer-consumer problem. The `pizzazone` goroutine acts as the producer, generating pizza orders and sending them through the `data` channel. The main goroutine acts as the consumer, receiving the orders and processing them. The channels ensure safe communication between the producer and consumer, preventing data races and synchronization issues.

The program also showcases the use of colored output for better visualization and readability.