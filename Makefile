.PHONY: start

start:
	cd frontend && docker-compose up -d
	cd backend && docker-compose up -d
reload:
	cd frontend && docker-compose up -d
	cd backend && docker-compose build --no-cache && docker-compose up -d
stop: 
	cd frontend && docker-compose down
	cd backend && docker-compose down
fill_data:
	cd backend && make fill_tables