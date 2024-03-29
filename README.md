Jungle Challenge

Configuration

To start, change the files `config.yaml` to match your environment settings, such as database

Database initialization

To popuplate the database, go to the migrations folder. Run the following command:

`go run . init <DATA_FILE_PATH>`

There are other options available, such as:
`go run . migrate up` -> Runs all migrations
`go run . migrate down` -> Rollback all migrations

Running

To run the code locally, just type in the root path.

`go run .`

The following APIs are available:

- `/info/?clientId=<CLIENT_ID>` -> Returns the client that matches the ID
- `/info/?startBirthday=<YYYY-MM-DD>&endBirthday=<YYYY-MM-DD>` -> Returns the list of clients with birthdays between the provided dates
- `/search/?name=<START_NAME>` -> Returns the first user with name that starts with the START_NAME parameter
