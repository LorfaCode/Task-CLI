```mermaid
sequenceDiagram
    participant User
    participant CLI
    participant TaskService
    participant Database

    User->>CLI: add "New Task"
    CLI->>TaskService: createTask("New Task")
    TaskService->>Database: saveTask("New Task")
    Database-->>TaskService: Task saved
    TaskService-->>CLI: Task created
    CLI-->>User: Task added successfully