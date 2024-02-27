stop:
	docker-compose down 

start:
	docker-compose up

restart:
	docker-compose down 
	docker rmi plata-backend && \
	docker compose up