build:
	docker build . -t delivery-much

run:
	docker run -p 3000:3000 delivery-much
