SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

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
  "max_tokens": 10, \
  "temperature": 1.1 \
}'

go-chatcomp:
	go run examples/chat_completions/main.go