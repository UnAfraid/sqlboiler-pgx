# sqlboiler-pgx

The main difference between sqlboiler-psql and sqlboiler-pgx is the driver used underneath:
- sqlboiler-psql is using https://github.com/lib/pq
- sqlboiler-pgx is using https://github.com/jackc/pgx

## Installation
```
# Install sqlboiler pgx driver
go install github.com/UnAfraid/sqlboiler-pgx@latest

# Generate models
sqlboiler pgx
```

Example configuration:
```toml
[pgx]
  dbname = "dbname"
  host   = "localhost"
  port   = 5432
  user   = "dbusername"
  pass   = "dbpassword"
  schema = "myschema"
  blacklist = ["migrations", "other"]
```
