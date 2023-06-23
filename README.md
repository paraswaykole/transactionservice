# transaction service

The project uses fiber and gorm in golang to build a transaction service as per the API specs.

Requirements:
- PostgreSQL database.
- .env file with `PG_DSN="host=<host here> user=<username here> password=<password here> dbname=<db name here> port=5432"`

Additionally, the postgres ltree extension is enable to store the path of each transaction to maintain hierarchy.