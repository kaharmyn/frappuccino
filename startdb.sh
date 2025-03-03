# /bin/bash
CURRENT_PATH="$(pwd)/db/init.sql"
docker run --name pg-container-name -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgresdb -v $CURRENT_PATH:/docker-entrypoint-initdb.d/init.sql -d postgres


# source: https://medium.com/@marvinjungre/get-postgresql-and-pgadmin-4-up-and-running-with-docker-4a8d81048aea
# https://medium.com/@nathaliafriederichs/setting-up-a-postgresql-environment-in-docker-a-step-by-step-guide-55cbcb1061ba

# docker stop $(docker ps -a -q)
