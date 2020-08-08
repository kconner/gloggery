# https://www.gnu.org/prep/standards/html_node/Directory-Variables.html
# Using $HOME instead of /usr/local though.
prefix=$(HOME)
exec_prefix=$(prefix)
bindir=$(exec_prefix)/bin

gloggery: src/*.go
	@go build -o $@ src/*.go
	@echo built

command_path=$(bindir)/gloggery
data_path=$(HOME)/.gloggery
install: gloggery
	@cp gloggery $(command_path)
	@echo command installed at $(command_path)
	@mkdir -p $(data_path)
	@mkdir -p $(data_path)/posts
	@mkdir -p $(data_path)/templates
	@cp -n templates/*.tmpl $(data_path)/templates/
	@echo data installed at $(data_path)

clean:
	@rm gloggery 2>/dev/null || true
	@rm output/* 2>/dev/null || true
	@rmdir output 2>/dev/null || true
	@echo cleaned

watch:
	@find src/*.go | entr make -s

watch-run:
	@mkdir -p output
	@find src/*.go | entr -s 'make -s && time ./gloggery --output output'

format:
	@go fmt src/*.go
	@echo formatted
