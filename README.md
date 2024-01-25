<h2>Welcome to GoEvents</h2>

<p>The objective of this project is to create a Rest API where the user can mark in their calendar a date where a certain event will occur with specific data, a title and a description, after entering the data the API will notify the user how the date approaches.</p>

```json
{
    "title": "My graduation",
    "description": "My college education",
    "date": "2024-12-21",
    "interval": 12
}
```

- The `interval` field means the interval (in days) at which the user would like to be notified about the upcoming event

- Run project `go run main.go`
