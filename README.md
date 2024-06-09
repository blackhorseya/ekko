# Ekko: Todo List and Issue Tracking System

Ekko is a comprehensive Golang-based application that functions as both a todo list and a Jira-like issue tracking
system. It includes workflow management, collaboration tools, and comprehensive reporting features. Ekko is designed to
be highly scalable and customizable, with built-in user authentication and authorization to ensure data security.

## Background

The Ekko project was developed to streamline task management and issue tracking for teams of all sizes. It integrates
essential functionalities for project management, allowing users to manage their tasks and issues efficiently. By
combining todo list simplicity with the powerful capabilities of issue tracking systems like Jira, Ekko provides a
versatile solution for various project management needs.

## Features

- **Todo List Management**: Easily create, edit, delete, and mark tasks as completed to keep track of daily activities.
- **Issue Tracking**: Manage issues with attributes like priority, status, and assignments.
- **Workflow Management**: Define and implement custom workflows for managing processes.
- **Collaboration Tools**: Facilitate team communication with integrated discussion forums and commenting on tasks and
  issues.
- **Reporting and Analytics**: Generate detailed reports and analytics to gain insights into project progress.
- **Scalability and Customizability**: Flexible design to meet diverse team needs.
- **User Authentication and Authorization**: Secure access control for user data.

## Tech Stack

- **Backend**: Golang
- **Database**: MongoDB
- **Frontend**: React
- **API**: RESTful API
- **Containerization**: Docker
- **CI/CD**: GitHub Actions

## Architecture

Ekko is built following Clean Architecture principles, ensuring a clear separation of concerns and maintainable code
structure. The architecture is divided into several key layers:

1. **Entities**:

    - These are the core business objects of the application.
    - They contain the essential business rules and properties that are independent of any external dependencies or
      frameworks.
    - Examples include task, user, and issue objects.

2. **Use Cases**:

    - Use cases represent the application-specific business logic.
    - They orchestrate the flow of data to and from the entities, and direct how the data can be changed or interacted
      with.
    - They encapsulate all the use case-specific business rules, ensuring that the application logic remains decoupled
      from
      the outer layers.

3. **Interface Adapters**:

    - This layer adapts the data from the outer layers (e.g., web controllers, database gateways) to the inner layers (
      use
      cases and entities).
    - It contains the implementation of interfaces that convert data from the format most convenient for the use cases
      and
      entities into the format needed by the frameworks and drivers.
    - Examples include REST controllers, presenters, and database repositories.

4. **Frameworks & Drivers**:

    - This outermost layer includes frameworks and tools such as the database, web server, UI, and any external APIs.
    - It contains the code that interacts with external systems and frameworks, such as database connectors, HTTP
      clients,
      and third-party services.
    - This layer depends on both the interface adapters and the use cases, providing implementations that are used by
      the
      interface adapters.

### Clean Architecture Benefits

- **Separation of Concerns**: Each layer has a clear responsibility, making the system more understandable and
  maintainable.
- **Independent of Frameworks**: Business rules are not bound to specific frameworks, allowing easier migration to new
  technologies.
- **Testable**: Business rules and application logic can be tested independently of external dependencies, improving
  test coverage and reliability.
- **Flexible and Adaptable**: The system can easily be extended or modified without affecting unrelated parts of the
  codebase.

## Installation

To set up Ekko locally, follow these steps:

1. **Clone the repository**:
    ```sh
    git clone https://github.com/blackhorseya/ekko.git
    cd ekko
    ```

2. **Run the application**:
    ```sh
    docker compose up --build
    ```

3. **Access the application**:
   Open your browser and go to `http://localhost:8080/api/docs/index.html`.

## Contribution Guidelines

We welcome contributions! To contribute, please review our [Contribution Guidelines](CONTRIBUTING.md) for detailed
instructions on coding standards, branch naming conventions, and testing requirements.

## License

This project is licensed under the GPL-3.0 License. See the [LICENSE](LICENSE) file for more details.
