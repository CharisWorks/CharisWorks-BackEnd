cd test_server
docker-compose down -v
docker-compose up --build &
stripe listen --forward-to localhost:8080/webhook
