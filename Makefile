.PHONY: proto dbUp dbDown

proto:
	@cd api/proto && \
		buf lint && \
		buf generate

dbUp:
	@cd ./tools/devdb && \
		docker-compose up -d --wait
		docker exec -i clickPlanet-redis /bin/sh -c "(redis-cli -x script load < /static/setAndPublishOnStream.lua) > /static/setAndPublishOnStream.sha1"

dbDown:
	@cd ./tools/devdb && \
		docker-compose down

dBuild:
	@docker build -t clickplanet-back:local .

dRun:
	@docker run \
		-d \
		-p 8080:8080 \
		--name clickplanet \
		-v ./cmd/api:/home/app/config \
		clickplanet-back:local \
		/home/app/api -config /home/app/config/config.yaml

dRm:
	@docker rm -f clickplanet

test:
	@go test -v ./... | grep -v 'no test files'