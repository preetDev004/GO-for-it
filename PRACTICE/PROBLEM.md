# Distributed Task Scheduler

## Problem Description

You're tasked with creating a simplified distributed task scheduler system for a company that processes large amounts of data. This system should allow multiple workers to pick up and complete tasks concurrently, ensuring efficient use of resources and preventing task duplication.

## Requirements

1. Task Management:
   - Create a `Task` struct with fields for ID, Description, Status (Pending, In Progress, Completed), and CreatedAt timestamp.
   - Implement a `TaskQueue` that stores tasks and provides thread-safe methods to add, retrieve, and update tasks.

2. Worker Management:
   - Create a `Worker` struct that represents a worker capable of processing tasks.
   - Implement a method for workers to pick up available tasks from the queue.
   - Ensure workers can update task status as they progress.

3. Scheduler:
   - Implement a `Scheduler` struct that manages the task queue and workers.
   - The scheduler should distribute tasks to available workers efficiently.
   - Implement logic to handle task timeouts and reassign tasks if a worker fails to complete them within a specified time.

4. Concurrency:
   - Use goroutines to simulate multiple workers processing tasks concurrently.
   - Implement proper synchronization to prevent race conditions when accessing shared resources.

5. Logging and Monitoring:
   - Implement a simple logging system to track task status changes and worker activities.
   - Create a method to display current system status, including pending tasks, active workers, and completed tasks.

6. Main Program:
   - In the `main` function, set up the scheduler with a predefined number of workers.
   - Generate a set of sample tasks and add them to the queue.
   - Start the scheduler and allow it to run for a set duration or until all tasks are completed.
   - Display final statistics about task completion and worker efficiency.

## Additional Challenges

- Implement task priorities and allow high-priority tasks to be processed first.
- Add support for task dependencies, where some tasks can only start after others are completed.
- Implement a simple REST API to allow external services to add tasks and query task status.
- Create a mechanism for workers to report partial progress on long-running tasks.
- Implement persistent storage for tasks to allow system recovery after a shutdown.

## Concepts to Utilize

Ensure your implementation makes use of the following Go concepts:

- Structs and interfaces
- Goroutines and channels for concurrency
- Mutexes or other synchronization primitives
- Error handling and logging
- Time-based operations (for timeouts and timestamps)
- Slices and maps for data management
- Packages for code organization
- Context for managing goroutine lifecycles (optional)

This problem provides a realistic scenario that incorporates many important Go concepts and real-world considerations like concurrency, resource management, and system design. It's a great opportunity to practice building a complex system with multiple interacting components.

Good luck with your implementation! Feel free to ask for clarification or hints if you need them as you work on this problem.