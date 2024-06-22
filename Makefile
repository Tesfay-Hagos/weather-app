start: 
	docker compose up -d
stop:
	docker compose down
restart:
	docker compose down
	docker compose up --build -d
logs:
	docker compose logs -f
ps:
	docker compose ps
exec:
	docker compose exec app bash
clean:
	docker compose down
	docker volume prune -f
	docker network prune -f
	docker image prune -f
	docker container prune -f
	docker system prune -f
	docker system prune --volumes -f
	docker system prune --all -f
	docker system prune --volumes -f
	docker system prune --volumes --all -f
	docker system prune --volumes --filter "label!=keep" -f
	docker system prune --volumes --filter "label!=keep" --all -f
	docker system prune --volumes --filter "label!=keep" --filter "label!=keep" -f
	docker system prune --volumes --filter "label!=keep" --filter "label!=keep" --all -f
	docker system prune --volumes --filter "label!=keep" --filter "label!=keep" --filter "label!=keep" -f
	docker system prune --volumes --filter "label!=keep" --filter "label!=keep" --filter "label!=keep" --all -f
	docker system prune --volumes --filter "label!=keep" --filter "label!=keep" --filter "label!=keep" --filter "label!=keep" -f
	docker system prune --volumes --filter "label!=keep" --filter "label!=keep" --filter "label!=keep" --filter "label!=keep" --all -f
	docker system prune --volumes --filter "label!=keep" --filter "label!=keep" --filter "label!=keep" --filter "label!=keep" --filter "label!=keep" -f
	docker system prune --volumes