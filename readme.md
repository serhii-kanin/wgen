### WGEN
#### Microservice-oriented workflow orchestrator

##### Features:
1. Provides a convenient way to specify your workflows via Girkin-like DSL which makes a task specification work fun and easy.
2. Interface independent. Workflow execution can be initiated via Kafka or gGRPC. There are many connectors for other protocols and message brokers.
3. Persistent and fault tolerant. Workflow persists its state into its database after each task execution. It provides retry mechanics and can be recovered after a failure. Workflow has many connectors for RDBMS and noSQL systems.
4. Comprehensive. WGEN implements sync/async combined execution, conditional tasks and many other nice features.

##### DSL Example:
```
@Retry each task two times                                                                       #workflow policies
@Retry wait period is fixed
@Retry wait period is one second
@Given (customer registration details)                                                           #this row tells your workflow about the input var CustomerRegistrationDetails
@Then validate (customer registration details) (with no retry)                                   #this row will validate customer's data with desabled retry policy
@Provide (customer validation errors)                                                            #stores validation errors
@Then create customer's (authentication record) using (customer registration details)            #authentication record object is available for other steps
@Then register (customer registration details) in first microserice                              #this and next 2 steps will be executed in parallel
@And register (customer registration details) in 2nd microserice
@And register (customer registration details) in 3rd microserice
@Then activate (authentication record)                                                           #both inputs and outputs are reusable
@Then notify customer about successful registration. Use (customer registration details - email) #You can specify tasks to receive specific fields.
```

Running 
`` go generate ./...
`` will generate a state machine for all flows in your project. 
No need to repeat a boring monkey job - you can concentrate on your business logic.

##### Integrations:
1. Kafka
2. gRPC client
3. Thrift client
4. Direct. You can integrate WGEN directly to your code.

##### Storage integrations:
1. MongoDb
2. go/database package (existent drivers)
3. Redis
You an easily integrate your storage. Check here and here.