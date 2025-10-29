.PHONY: run

# The default target. Running 'make' will execute this.
run:
	go run main.go

# Optional: A clean target is always handy, even if simple.
clean:
	@echo "No build artifacts to clean up."
