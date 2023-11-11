To spin up a dev database:

```
docker compose -f docker.compose.postgres.yml up -d
```

To connect to the database via pgadmin

```
at http://127.0.0.1:15432

sign in creds:

email: admin@pgadmin.com
password: password
```

To connect to the server

```
address: host.docker.internal
host:5432
```

```
run this to enable the uuid extension for postgres
# CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```
