```mermaid
graph TD
    A[Main] --> B[Commands]
    B --> C[Task]
    C --> D[Storage]
    B -->|add| C
    B -->|update| C
    B -->|list| C
    B -->|delete| C
    B -->|mark| C