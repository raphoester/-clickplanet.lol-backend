.PHONY: proto dbUp dbDown

proto:
	@cd api/proto && \
		buf lint && \
		buf generate

dbUp:
	@cd ./tools/devdb && \
		docker-compose up -d --wait
		docker exec -i clickPlanet-redis /bin/sh -c "(redis-cli -x script load < /static/setAndPublish.lua) > /static/setAndPublish.sha1"

dbDown:
	@cd ./tools/devdb && \
		docker-compose down
