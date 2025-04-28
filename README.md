# Go-PSQL
Personal package to make working with JSONs easiers in postgres


## Usage

### SQL Scanning

```go
var res gopsql.JSONB
err = db.QueryRow("SELECT data FROM test").Scan(&res)

```

### SQL Arguments

```go
sample := gopsql.JSONB{
    "foo": "bar",
}
_, err = db.Exec("INSERT INTO test (data) VALUES ($1)", sample)

```


## License

BSD 3-Clause License

## Author

karim-w

## Contributing

Pull requests are welcome. For major changes, please open an issue first to
discuss what you would like to change.
