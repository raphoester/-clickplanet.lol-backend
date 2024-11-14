
.PHONY: proto

proto:
	@cd api/proto && \
		buf lint && \
		buf generate
