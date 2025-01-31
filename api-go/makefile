default: hello

hello: 
	@printf "Welcome to the Fancy Todo written in 3 different languages.\n\
	They are Typescript, Go, and Rust.\n\
	Enjoy.\n"

down:
	docker compose down

db\:up:
	docker compose up -d pg

db\:clean:
	docker volume rm api-go_todo-app

db\:exec:
	docker exec -it pg /bin/sh

go\:run:
	go run cmd/main.go