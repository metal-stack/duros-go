MAKEFLAGS += --no-print-directory
BUF_VERSION := 1.55.1

_buf:
	docker run --rm \
		--entrypoint sh \
		-v $(PWD)/..:/workspace \
		-w /workspace/proto \
		bufbuild/buf:$(BUF_VERSION) \
		-c "buf --timeout 5m $(CMD) && chown -R $(shell id -u):$(shell id -g) /workspace"

.PHONY: protolint
protolint:
	@$(MAKE)	_buf	CMD="format -w lightbits"
	@$(MAKE)	_buf	CMD="lint -v"

.PHONY: protoc
protoc:
	@$(MAKE)	_buf	CMD="generate -v"
