-- CREATE DATABASE usersdb;
-- CREATE DATABASE paymentdb;
-- Create paymentdb if it doesn’t exist
DO
$do$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_database WHERE datname = 'paymentdb'
   ) THEN
      CREATE DATABASE paymentdb OWNER app;
   END IF;
END
$do$;



-- podman run -d --name pg --network demo-network -p 5432:5432 -e POSTGRES_USER=app -e POSTGRES_PASSWORD=app123 -v ./db-init:/docker-entrypoint-initdb.d:Z  docker.io/library/postgres