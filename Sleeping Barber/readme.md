# The Sleeping Barber Problem

## Problem Statement

The Sleeping Barber Problem is a classic concurrency problem in computer science and operating systems theory. It involves a barbershop with a waiting room that has a limited number of seats. The problem arises when managing the coordination between the barber, who can only serve one customer at a time, and the customers who arrive at the barbershop.

The key challenges in this problem are:

1. **Synchronization**: Ensuring that the barber and customers are properly synchronized, so that the barber doesn't sleep when there are customers waiting, and customers don't leave the shop unnecessarily when the barber is available.
2. **Resource Management**: Managing the limited seating capacity in the waiting room, ensuring that customers don't enter the shop if the waiting room is full.
3. **Fairness**: Providing a fair scheduling mechanism for customers, so that no customer is indefinitely denied service (starvation).

## Solution

The provided Go code solves the Sleeping Barber Problem using goroutines (lightweight threads) and channels for communication and synchronization. Here's how the solution works:

1. **Defining Structures and Constants**
   - The `BarberShop` struct represents the barbershop with fields for seating capacity, haircut duration, number of barbers, client and barber channels, and a flag indicating whether the shop is open or closed.
   - Constants like `seatingCapacity`, `arrivalRate`, `cutDuration`, and `timeOpen` are defined to control the behavior of the barbershop.

2. **Main Function**
   - The `main` function initializes the problem by creating a `BarberShop` instance, adding a barber, and starting the barbershop as a goroutine.
   - It also starts a goroutine to simulate the arrival of clients at a random rate.
   - The `main` function blocks until the barbershop is closed.

3. **Adding Barbers and Clients**
   - The `addBarber` function starts a new goroutine for each barber, which continuously checks the client channel for new clients.
   - The `addClient` function simulates a new client arriving at the barbershop. If the shop is open and the waiting room has free seats, the client is added to the client channel. Otherwise, the client leaves the shop.

4. **Cutting Hair**
   - When a barber finds a client in the client channel, they remove the client from the channel and start cutting their hair using the `cutHair` function.
   - The `cutHair` function simulates the haircut process by sleeping for the specified `HairCutDuration`.

5. **Closing the Shop**
   - After the `timeOpen` duration has elapsed, the `closeShop` function is called to close the barbershop.
   - The `closeShop` function stops accepting new clients, waits for all barbers to finish their current haircuts, and then closes the shop.

6. **Synchronization and Communication**
   - The `ClientsChan` channel is used for clients to enter the waiting room, ensuring that they don't enter if the waiting room is full.
   - The `BarbersDoneChan` channel is used for barbers to signal when they are done with their work, allowing the shop to close properly.
   - The use of goroutines allows for concurrent execution of barbers and clients, while channels provide a safe mechanism for communication and synchronization.

By using goroutines and channels, the solution ensures that the barber doesn't sleep when there are customers waiting, and customers don't leave the shop unnecessarily when the barber is available. The seating capacity of the waiting room is managed by controlling the number of clients added to the client channel. The arrival of new clients is simulated at a random rate, ensuring fairness in customer scheduling.

The solution also demonstrates the use of colored output for better visualization and readability.