## Architecture Overview ##
The system is designed as a stateless API-driven service to handle larage CSV file upload with asynchronous background processing to ensure scalability, fault isolation, and predictable performance under load.

## Components ##
- An HTTP API Layer using Fiber
- A background worker pool
- Go In-memory data management

## Request Flow ##
### 1. Client Request ###
- Client interact via RESTful HTTP API
### 2. API Layer ###
- The API Layer is stateless
- Performs form validation
- For long-running or failure-prone tasks for example ```/statements``` endpoint, the API publishes messages to the broker and returns an acknowledgment.
### 3. Message Broker ###
- Acts as the decoupling mechanism between request handling and background processing.
- Enables controlled retry, backoff, and dead-letter handling without blocking user requests.
### 4. Worker Layer ###
- Workers consume messages independently of the API lifecycle.
- Processing is idempotent to tolerate retries and duplicate delivery.
- Concurrency is controlled at the worker level to avoid resource contention.
### 5. Data Consistency & Reliability ###
- Data consistency is handled by ```sync``` package in golang.
- Failures in background processing do not impact API availability.
### 6. Observability ###
- For observability, data can be accessed through ```/health``` endpoint
- Showing number of total worker, busy worker, and available worker
- Showing number of retry inflight process
### 7. Security Considerations ###
- Authentication should be implmented on all endpoint in the next iteration
### 8. Tradeoffs ###
- Currently, all data were saved to in-memory golang variable, will be cleaned on every restart

## How To Run ##
- run ```go build .```
- run ```./flip```

## How To Test ##
- ```curl --location 'localhost:8080/statements' \ --form 'file=@"/flip/transaction.csv"'```
- take ```upload_id``` from the response for next curl
- ```curl --location 'localhost:8080/balance?upload_id={upload_id}'```
- ```curl --location 'localhost:8080/transactions/issues?upload_id={upload_id}&page=1&size=10'```
- ```curl --location 'localhost:8080/health'```