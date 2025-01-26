# Order Management System Specification

You need to implement an order management system that processes incoming orders based on their priority and updates their status. This system functions as a **Worker**, and orders are continuously fed into it.

## Order Structure:
- **OrderID**: Order identifier.
- **Priority**: Order priority (possible values: `High` or `Normal`).
- **Status**: Order status (possible values: `Failed`, `Pending`, `Processed`).
- **ProcessingTime**: Processing time in seconds (this value should be defined within each order for testing purposes).

## System Requirements:
1. Orders with `High` priority must be processed before `Normal` priority orders, even if they arrive later.
2. After successful processing of an order, its status should be updated to `Processed`.
3. If the processing time of an order exceeds 5 seconds, the order should be canceled, and its status should be updated to `Failed`.
4. Every 2 seconds, a report should be printed, showing the count of orders in each status: `Pending`, `Processed`, and `Failed`.
5. The implementation should be optimized in terms of performance and follow clean coding principles.

## Additional Notes and Guidelines:
- Only the Go programming language is allowed for this task.
- Simulate the incoming order stream.
- If possible, Dockerize the project.
- Finally, provide a compressed file containing the complete project, including the code structure, related files, and instructions for running the program.
- If additional information is needed but not included in the question, you can make assumptions and proceed with the implementation.
