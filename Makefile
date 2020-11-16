.EXPORT_ALL_VARIABLES:


run: 
	docker-compose up 

stop:
	docker-compose stop

test:
	go test -v ./...