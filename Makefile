build:
	@go build -o refactoring cmd/app/main.go
	@echo "Refactoring is compiled!"

test:
	@echo "TODO"

clean:
	@rm refactoring

.PHONY: build, test, clean