# CRM Rules Engine Architecture

## Technical Learning Objectives

A rules processing system implementing microservices architecture with Go, focusing on modular design and message-driven communication.

## Architectural Patterns

### Core Design Approach
The system implements a message-driven microservices architecture, focusing on:
- Decoupled service communication
- Asynchronous processing
- Modular system design

### Service Responsibilities

1. **User Service**
   - Provides CRUD operations for user data
   - Manages basic user information storage

2. **Rules Service**
   - Manages rule definitions and configurations
   - Provides CRUD operations for rules
   - Separates storage and business logic concerns

3. **Rules Execution Worker**
   - Processes and executes rules
   - Handles background task processing
   - Implements message consumption patterns

4. **Log Service**
   - Manages log storage and retrieval
   - Provides endpoints for log access

5. **Log Worker**
   - Handles log ingestion
   - Processes and stores log entries

6. **Scheduler**
   - Manages rule scheduling
   - Triggers rule execution at specified intervals

7. **Gateway Service**
   - Provides API gateway functionality
   - Handles routing and request forwarding

8. **Migration Service**
   - Manages database schema migrations
   - Ensures database structure is up to date

   
## Technical Deep Dives

### Messaging Pattern
**Redis Pub/Sub Implementation**
- Uses Redis Pub/Sub for direct message passing between services
- Multiple channels for different message types (tasks, logs)

#### Go Language
- Chosen for its simplicity and standard library
- Provides straightforward approach to building microservices

### Database Interaction
**PostgreSQL Integration**
- Uses PostgreSQL for data storage
- Separates data access and service logic

### Scalability Considerations
- Modular service design supports independent scaling

### Messaging Workflow

1. **Client Interaction**:
   - All client requests are initially routed through gateways
   - Gateways act as the entry point and forward requests to the appropriate service

2. **Rule Creation Workflow**:
   - When a client wants to create a new rule:
     a. Gateway forwards the request to the `rules-service`
     b. `scheduler` picks up the new rule
     c. Publishes a message to the `rules-execution-worker`
     d. Rule is executed according to specified conditions
     e. Impacts are processed for the user retrieved from the `user-service` database

3. **Logging and Tracking**:
   - After rule execution, a message is published to the `log-worker`
   - Logs are stored in the database
   - Clients can access logs through the `log-service` via gateways

## Learning Insights

### Design Principles Applied

- **Loose Coupling**: Each service operates independently, minimizing dependencies. This design choice enhances flexibility and allows for individual scaling and deployment without affecting other services.

- **High Cohesion**: Services are designed with a single responsibility in mind, ensuring that each performs its specific function efficiently. This focused approach simplifies maintenance and enhances code clarity.

- **Extensibility**: The architecture is modular, supporting future enhancements and scalability. This design allows for the seamless addition of new features or services as the system evolves.

### Architectural Rationale

The CRM Rules Engine is structured as a microservices architecture, where each service is independent and modular. Core services include the User Service, Rules Service, Rules Execution Worker, Scheduler, Log Worker, and Log Service.

- **Separation of Scheduler and Rules Execution Worker**: From a scalability perspective, the scheduler's role is lightweight and doesn't require extensive resources, allowing it to remain stable without scaling. Conversely, the Rules Execution Worker handles resource-intensive tasks, necessitating scalability to manage heavy workloads effectively.

- **Task Execution and Logging Separation**: By decoupling task execution from logging, both can scale independently. As the Task Execution Worker scales to handle increased workloads, the logging system must also scale to capture and manage the resulting data efficiently.

- **Independent Log Worker and Log Service**: The separation ensures that log ingestion and log service endpoints operate independently. The Log Worker can handle high ingestion rates, while the Log Service manages access and retrieval, both capable of handling heavy traffic without bottlenecks.

This architecture ensures robust performance and scalability.

[Back to README](../README.md)
