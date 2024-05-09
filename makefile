SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Examples

curl-health:
	curl -il https://api.predictionguard.com \
     -H "x-api-key: ${PGKEY}"

go-health:
	go run examples/healthcheck/main.go

curl-chatcomp:
	curl -il -X POST https://api.predictionguard.com/chat/completions \
     -H "x-api-key: ${PGKEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "Neural-Chat-7B", \
		"messages": [ \
			{ \
			"role": "user", \
			"content": "How do you feel about the world in general" \
			} \
		], \
		"max_tokens": 1000, \
		"temperature": 1.1 \
	}'

curl-chatcomp-sse:
	curl -il -X POST https://api.predictionguard.com/chat/completions \
     -H "x-api-key: ${PGKEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "Neural-Chat-7B", \
		"messages": [ \
			{ \
			"role": "user", \
			"content": "How do you feel about the world in general" \
			} \
		], \
		"stream": true \
	}'

go-chatcomp:
	go run examples/chat_completions/basic/main.go

go-chatcomp-sse:
	go run examples/chat_completions/sse/main.go

curl-comp:
	curl -il -X POST https://api.predictionguard.com/completions \
     -H "x-api-key: ${PGKEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "Neural-Chat-7B", \
		"prompt": "Will I lose my hair", \
		"max_tokens": 1000, \
		"temperature": 1.1 \
	}'

go-comp:
	go run examples/completions/basic/main.go

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy

deps-list:
	go list -m -u -mod=readonly all

deps-upgrade:
	go get -u -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all