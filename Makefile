gloggery: src/*.go
	@go build -o $@ src/*.go
	@echo built

format:
	@go fmt src/*.go
	@echo formatted

clean:
	@rm gloggery 2>/dev/null || true
	@echo cleaned

watch:
	@find src/*.go | entr make -s

watch-run:
	@find src/*.go | entr -s 'make -s && ./gloggery'
