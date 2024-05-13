SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Examples

curl-health:
	curl -il https://api.predictionguard.com \
     -H "x-api-key: ${PGKEY}"

go-health:
	go run examples/healthcheck/main.go

curl-chat:
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

curl-chat-sse:
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

go-chat:
	go run examples/chat/basic/main.go

go-chat-sse:
	go run examples/chat/sse/main.go

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
	go run examples/completions/main.go

curl-factuality:
	curl -X POST https://api.predictionguard.com/factuality \
     -H "x-api-key: ${PGKEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"reference": "The President shall receive in full for his services during the term for which he shall have been elected compensation in the aggregate amount of 400,000 a year, to be paid monthly, and in addition an expense allowance of 50,000 to assist in defraying expenses relating to or resulting from the discharge of his official duties. Any unused amount of such expense allowance shall revert to the Treasury pursuant to section 1552 of title 31, United States Code. No amount of such expense allowance shall be included in the gross income of the President. He shall be entitled also to the use of the furniture and other effects belonging to the United States and kept in the Executive Residence at the White House.", \
		"text": "The president of the united states can take a salary of one million dollars" \
	}'

go-factuality:
	go run examples/factuality/main.go

curl-translate:
	curl -X POST https://api.predictionguard.com/translate \
     -H "x-api-key: ${PGKEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"text": "The rain in Spain stays mainly in the plain", \
		"source_lang": "eng", \
		"target_lang": "som" \
	}'

go-translate:
	go run examples/translate/main.go

curl-rpi:
	curl -X POST https://api.predictionguard.com/PII \
     -H "x-api-key: ${PGKEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"prompt": "My email is bill@ardanlabs.com and my number is 954-123-4567.", \
		"replace": true, \
		"replace_method": "mask" \
	}'

go-rpi:
	go run examples/rpi/main.go

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