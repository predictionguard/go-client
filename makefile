SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

curl-health:
	curl https://api.predictionguard.com \
     -H "x-api-key: ${PGKEY}"

go-health:
	go run examples/healthcheck/main.go
