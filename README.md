# Exploring REST API and CLI Interaction in Go

This is an experiment to explore how a Go CLI and REST API interact. It covers:

- Making HTTP requests from a CLI.
- Exploring authentication between the REST API and CLI
- Embedding SQLite into a Go application

## WIP Requirements

- Offline First: The CLI works without signup, using an embedded SQLite database.
- Simplified Schema: The local SQLite database assumes a single user (no user table).
- User Signup & Syncing:
    - A user can optionally sign up, creating an account in the remote database.
    - Local data is synced to the remote database upon signup.
    - An API key is stored in a config file for remote connections.
- Authentication:
    - The REST API uses JWT-based authentication.
- Real-time Updates:
    - Any changes made via the CLI or REST API will be broadcast over a WebSocket connection.
