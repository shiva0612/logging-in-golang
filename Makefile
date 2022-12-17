.PHONY: simplelog zaplog

simplelog:
	cd simplelog && rm *.log && go run main.go
zaplog:
	cd zap-log && rm *.log && go run main.go
