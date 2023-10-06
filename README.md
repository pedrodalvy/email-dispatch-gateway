# Email Dispatch Gateway

This repository was developed as part of the
"[Go (Golang) - Do zero ao avançado](https://www.udemy.com/course/golang-do-zero-ao-avancado/)" course
on [Udemy](https://www.udemy.com/). The project consists of an API that serves as a gateway for sending emails. It
implements authentication using JWT tokens with Keycloak. The chosen framework for developing the API is Chi.

## Important Notes:

This project primarily serves as a practical exercise to apply concepts learned in the course. While it features a
well-developed authentication layer, it's essential to emphasize that it's not intended to be a production-ready
solution. Its main purpose is educational, demonstrating best practices and the implementation of features in Go 
(Golang).

## Running this project:

1. Clone the repository to your local machine using the following command:
    ```sh
    git clone git@github.com:pedrodalvy/email-dispatch-gateway.git
    ```

2. Duplicate the `.env.EXAMPLE` file and name it as `.env`, then fill it with all the required information.

3. Start the infrastructure:
    ```sh
    make infra
    ```

4. Start the server:
    ```sh
    make server
    ```

## Consuming the API Using a JetBrains IDE:

1. Replicate the `example-http-client.env.json` file and rename it to `http-client.env.json`. Populate this new file
with all the necessary information.

2. To access all API endpoints, authenticate via the `Credtentials B2C` endpoint.

## Running Tests:

To execute the tests, use the following command:
```shell
go test ./... --cover
```
