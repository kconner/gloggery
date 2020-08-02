# https://www.gnu.org/prep/standards/html_node/Directory-Variables.html
prefix=/usr/local
exec_prefix=$(prefix)
bindir=$(exec_prefix)/bin

gloggery: src/*.go
	@go build -o $@ src/*.go
	@echo built

install_path=$(bindir)/gloggery
install: gloggery
	@cp gloggery $(install_path)
	@echo installed at $(install_path)

clean:
	@rm gloggery 2>/dev/null || true
	@echo cleaned

watch:
	@find src/*.go | entr make -s

watch-run:
	@find src/*.go | entr -s 'make -s && ./gloggery'

format:
	@go fmt src/*.go
	@echo formatted
