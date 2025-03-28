SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Examples

curl-capability:
	curl -i -X GET https://api.predictionguard.com/models \
     -H "Authorization: Bearer $(PREDICTIONGUARD_API_KEY)" \
     -H "Content-Type: application/json"

curl-capability-completion-chat:
	curl -i -X GET https://api.predictionguard.com/models/chat-completion \
     -H "Authorization: Bearer $(PREDICTIONGUARD_API_KEY)" \
     -H "Content-Type: application/json"

go-capability:
	go run examples/capability/main.go

curl-chat:
	curl -i -X POST https://api.predictionguard.com/chat/completions \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "neural-chat-7b-v3-3", \
		"messages": "How do you feel about the world in general. I think the world is grand.", \
		"max_tokens": 1000, \
		"temperature": 1.1, \
		"top_p": 0.1, \
		"top_k": 50, \
		"output": { \
			"factuality": false, \
			"toxicity": true \
		}, \
		"input": { \
			"pii": "replace", \
			"pii_replace_method": "random" \
		} \
	}'

go-chat:
	go run examples/chat/basic/main.go

curl-chat-multi:
	curl -i -X POST https://api.predictionguard.com/chat/completions \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "Hermes-2-Pro-Llama-3-8B", \
		"messages": [ \
			{ \
			"role": "user", \
			"content": "How do you feel about the world in general. I think the world is grand." \
			} \
		], \
		"max_tokens": 1000, \
		"temperature": 1.1, \
		"top_p": 0.1, \
		"top_k": 50, \
		"output": { \
			"factuality": true, \
			"toxicity": true \
		}, \
		"input": { \
			"pii": "replace", \
			"pii_replace_method": "random" \
		} \
	}'

go-chat-multi:
	go run examples/chat/multi/main.go

curl-chat-sse:
	curl -i -X POST https://api.predictionguard.com/chat/completions \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "Hermes-2-Pro-Llama-3-8B", \
		"messages": [ \
			{ \
			"role": "user", \
			"content": "How do you feel about the world in general. I think the world is grand." \
			} \
		], \
		"stream": true, \
		"max_tokens": 300, \
		"temperature": 0.1, \
		"top_p": 0.1, \
		"top_k": 50, \
		"input": { \
			"pii": "replace", \
			"pii_replace_method": "random" \
		} \
	}'

go-chat-sse:
	go run examples/chat/sse/main.go

curl-chat-vision:
	curl -i -X POST https://api.predictionguard.com/chat/completions \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "llava-1.5-7b-hf", \
		"messages": [ \
			{ \
				"role": "user", \
				"content": [ \
					{ \
						"type": "text", \
						"text": "is there a deer in this picture" \
					}, \
					{ \
						"type": "image_url", \
						"image_url": { \
							"url": "data:image/jpeg;base64,$(IMAGE)" \
						} \
					} \
				] \
	 	  	} \
		], \
		"max_tokens": 300, \
		"temperature": 0.1, \
		"top_p": 0.1, \
		"top_k": 50 \
	}'

go-chat-vision:
	go run examples/chat/vision/main.go

curl-comp:
	curl -i -X POST https://api.predictionguard.com/completions \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "Hermes-2-Pro-Llama-3-8B", \
		"prompt": "Will I lose my hair by the time I am 64?", \
		"max_tokens": 1000, \
		"temperature": 1.1, \
		"top_p": 0.1, \
		"top_k": 50, \
		"output": { \
			"factuality": true, \
			"toxicity": true \
		}, \
		"input": { \
			"pii": "replace", \
			"pii_replace_method": "random" \
		} \
	}'

go-comp:
	go run examples/completion/main.go

curl-embed-basic:
	curl -i -X POST https://api.predictionguard.com/embeddings \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "bridgetower-large-itm-mlm-itc", \
		"input": [ \
        	{ \
			"text": "This is Bill Kennedy, a decent Go developer.", \
            "image": "$(IMAGE)" \
          	} \
    	] \
	}'

go-embed-basic:
	go run examples/embedding/basic/main.go

curl-embed-ints:
	curl -i -X POST http://localhost:6000/embeddings \
     -H "Authorization: Bearer $(PG_API_PREDICTIONGUARD_API_KEY)" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "multilingual-e5-large-instruct", \
		"input": [ \
			[0, 3293, 83, 19893, 118963, 25, 7, 3034, 5, 2] \
		] \
	}'

go-embed-ints:
	go run examples/embedding/ints/main.go

curl-tokenize:
	curl -i -X POST https://api.predictionguard.com/tokenize \
     -H "Authorization: Bearer $(PREDICTIONGUARD_API_KEY)" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "Hermes-2-Pro-Mistral-7B", \
		"input": "how many tokens exist for this sentence." \
	}'

go-tokenize:
	go run examples/tokenize/main.go

curl-embed-truncate:
	curl -i -X POST https://api.predictionguard.com/embeddings \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "multilingual-e5-large-instruct", \
		"truncate": true, \
		"truncate_direction": "Right", \
		"input": [ \
        	{ \
			"text": "This is Bill Kennedy, a decent Go developer." \
          	} \
    	] \
	}'

go-embed-truncate:
	go run examples/embedding/truncate/main.go

curl-factuality:
	curl -X POST https://api.predictionguard.com/factuality \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"reference": "The President shall receive in full for his services during the term for which he shall have been elected compensation in the aggregate amount of 400,000 a year, to be paid monthly, and in addition an expense allowance of 50,000 to assist in defraying expenses relating to or resulting from the discharge of his official duties. Any unused amount of such expense allowance shall revert to the Treasury pursuant to section 1552 of title 31, United States Code. No amount of such expense allowance shall be included in the gross income of the President. He shall be entitled also to the use of the furniture and other effects belonging to the United States and kept in the Executive Residence at the White House.", \
		"text": "The president of the united states can take a salary of one million dollars" \
	}'

go-factuality:
	go run examples/factuality/main.go

curl-health:
	curl -i https://api.predictionguard.com \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \

go-health:
	go run examples/healthcheck/main.go

curl-injection:
	curl -X POST https://api.predictionguard.com/injection \
	 -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"prompt": "A short poem may be a stylistic choice or it may be that you have said what you intended to say in a more concise way.", \
		"detect": true \
	}'

go-injection:
	go run examples/injection/main.go

curl-replace-pii:
	curl -X POST https://api.predictionguard.com/PII \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"prompt": "My email is bill@ardanlabs.com and my number is 954-123-4567.", \
		"replace": true, \
		"replace_method": "mask" \
	}'

go-replace-pii:
	go run examples/replacepi/main.go

curl-rerank:
	curl -i -X POST https://api.predictionguard.com/rerank \
     -H "Authorization: Bearer $(PREDICTIONGUARD_API_KEY)" \
     -H "Content-Type: application/json" \
     -d '{ \
		"model": "bge-reranker-v2-m3", \
		"query": "What is Deep Learning?", \
		"documents": ["Deep Learning is not pizza.", "Deep Learning is pizza."], \
		"return_documents": true \
     }'

go-rerank:
	go run examples/rerank/main.go

curl-toxicity:
	curl -X POST https://api.predictionguard.com/toxicity \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"text": "Every flight I have is late and I am very angry. I want to hurt someone." \
	}'

go-detect-toxicity:
	go run examples/toxicity/main.go

curl-translate:
	curl -X POST https://api.predictionguard.com/translate \
     -H "Authorization: Bearer ${PREDICTIONGUARD_API_KEY}" \
     -H "Content-Type: application/json" \
     -d '{ \
		"text": "The rain in Spain stays mainly in the plain", \
		"source_lang": "eng", \
		"target_lang": "spa" \
	}'

go-translate:
	go run examples/translate/main.go

# ==============================================================================
# Running tests within the local computer

test-r:
	CGO_ENABLED=1 go test -race -count=1 ./...

test-only:
	CGO_ENABLED=0 go test -count=1 ./...

lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vuln-check:
	govulncheck ./...

test: test-only lint vuln-check

test-race: test-r lint vuln-check

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

# ==============================================================================
# Image Data
# base64 -i imagefile

# https://pbs.twimg.com/profile_images/1571574401107169282/ylAgz_f5_400x400.jpg
IMAGE := /9j/4AAQSkZJRgABAQAAAQABAAD/4gKgSUNDX1BST0ZJTEUAAQEAAAKQbGNtcwQwAABtbnRyUkdCIFhZWiAAAAAAAAAAAAAAAABhY3NwQVBQTAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA9tYAAQAAAADTLWxjbXMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAtkZXNjAAABCAAAADhjcHJ0AAABQAAAAE53dHB0AAABkAAAABRjaGFkAAABpAAAACxyWFlaAAAB0AAAABRiWFlaAAAB5AAAABRnWFlaAAAB+AAAABRyVFJDAAACDAAAACBnVFJDAAACLAAAACBiVFJDAAACTAAAACBjaHJtAAACbAAAACRtbHVjAAAAAAAAAAEAAAAMZW5VUwAAABwAAAAcAHMAUgBHAEIAIABiAHUAaQBsAHQALQBpAG4AAG1sdWMAAAAAAAAAAQAAAAxlblVTAAAAMgAAABwATgBvACAAYwBvAHAAeQByAGkAZwBoAHQALAAgAHUAcwBlACAAZgByAGUAZQBsAHkAAAAAWFlaIAAAAAAAAPbWAAEAAAAA0y1zZjMyAAAAAAABDEoAAAXj///zKgAAB5sAAP2H///7ov///aMAAAPYAADAlFhZWiAAAAAAAABvlAAAOO4AAAOQWFlaIAAAAAAAACSdAAAPgwAAtr5YWVogAAAAAAAAYqUAALeQAAAY3nBhcmEAAAAAAAMAAAACZmYAAPKnAAANWQAAE9AAAApbcGFyYQAAAAAAAwAAAAJmZgAA8qcAAA1ZAAAT0AAACltwYXJhAAAAAAADAAAAAmZmAADypwAADVkAABPQAAAKW2Nocm0AAAAAAAMAAAAAo9cAAFR7AABMzQAAmZoAACZmAAAPXP/bAEMABQMEBAQDBQQEBAUFBQYHDAgHBwcHDwsLCQwRDxISEQ8RERMWHBcTFBoVEREYIRgaHR0fHx8TFyIkIh4kHB4fHv/bAEMBBQUFBwYHDggIDh4UERQeHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHv/CABEIAZABkAMBIgACEQEDEQH/xAAbAAACAwEBAQAAAAAAAAAAAAAEBQIDBgEAB//EABoBAAMBAQEBAAAAAAAAAAAAAAECAwAEBQb/2gAMAwEAAhADEAAAAaPKLvG9Vl4MhGearLO+Nw+dBwk1S35zKKHSbCZyJnoxzBz8u1gGA2t5CFF7muc8rz4PXHTiZyio0slJU2Z2HBxcxcWBsQUIcRrgOo44NM0W9Q6KaLVLBTRmNopg5W0Q4Y66JUBio3jo31XEmLpI3sHlzdN8h5qYhG5nphtEzVTPUlk24CafJP4uP81+w/JavoaC9dDo+Oa3KP8A1eLViKHHmWxCdwo9LlDKsb0MSZT4rMgGAUnsEOFcdcLnmBGXdi4pYOV90opIjSUKCIEepvgdyFluM/SmDOmzy5zAgNM27Z7h6oXRuw4tZStOoA8BkUZ36FiOiS4optRM89zetlZvhftnxSfpSZ2sOzyB2M+TFBnIVjOPaCLq7eqRpkjTrBHokyOydL9Fx9iSTmsFGo1GeqiGUpdfON7tmwsoXkDlDG5qrxy1NV1Ra5gsh0bSd73h6ez74bt9N5C4EgLqg+rSbDlq0lVJ+T5r9ET0Cms+TMyvSnhtYE86eCdMbZW5cv8AHGRCI2IsWXAmD8s2QL2eXhXYUqmpOjJwRUbbPPWtZ1zXNJF9mI6eDDLe1tmbGcf5dgfc0ZxbPQ1Vs2yq3fIKJou3y5HH6TMYOwuzDFgs1PfzvH6fS8NAxk+Ltz6/PINT6HIOGyM65U3S4rD9IrGqsjaRT3scZ186DCBMTpJXbAj59Vvl6vnDYjZriRuRfbvfl7Dm6d9Urv5ewiuHiVqdwqZ3jVYbHEUxmBxK9XVlpOS7BYdn7aHZSwxitmq9Hn1C9lmJjPWlM/V8qpjR7dRgy5Yu0kcsrC7cfAWsNxTljNmcQLMSZrwsHPM0GT9AcfJGGX6p3GuSjmge85as1VCPlrraFc/6F8623D2soy5xdyda2C1WBoxSr2fOleLmIDI/7HsZe93g3ZR6Bh0+hzvq87xZajGOCWp6K8EVsOiQNOnOdcwxtWnPJ5UIPrK8t3Npb8pzbZkYOIH0mn56Rho1JDBky9O3GKpGga3bWGYtrJnEq4Rrpn2D3nL0sOS5zdSoQ0RbMiRiine96FgMTFkP9z0V77nsJejzbM5rSIvT5qcj9Q+W2QM9Pf0S0YCCRDMOForR0nuYfxFuYTrKeZbxn7ZNF1Rgr4dQVpjbzLw4LmDc3Mx09QCv6QQSubc9nP09ay87t77npVVjkUirAocgpL3vBYehcUJ9zyLL0fbSGvAsmcVlR7ubShDETp8yRfdPjPZzMYH2OyrrfiVWkND2ZNfoyBXMT008+W9sI7ZSjVAlcqv0iY8ykTTTK5ab+hcpoeilR6bDmRK4U/QZU3N8oeR6XZQtUqId4lWBA5DJ33uZYQuVvN97nk3fe9t4cgeiYvAb+70ImFKdBAu/hn3L5KUXXjMe+crE5ZJkY8JvZKlubaXYPWp0NvXJw8M2k01OSBtc9OAJV2KMLWDAYJkaU8xnC42PUT9I+a/SODr96NvI3J89sk9GSVPIGIafveiRL5n9J+a9XN9Z9KHPT3veB72PcMiWEZ3cVWuxuwhYP519A+S3OVcgbL0uEq9b6XUSCo682N6Y0oUtY3pZuj3ICdmFsO8/Eno0d7wy7jQ1JYG+gqdB7SawwBY5AtZt0jvzrdnX2Ws9zmyK2i1KHkCkFO89wiPzv6J8/wCzm+yhtKn5ldbMXn6Bu95KuWvjLs4w9lltRG3fkP13FUfNk039sTVrH2aJkSGEQ2MtlS1p7O25KM+3Ms+Wvx8vlECqsvgaPpj4j0s16iskcwPpCoS82nvc4mt7GR2cIFIVzyBiCvee9hzCbzMdMPpsh7LcdlPaMaKiaOfrzcq7ejlp02Z00L+XsYLT5qTQZ6MmHILGo4sz1md5WtJIGINFzm8Ikt1NtV5iupanFM7PVWHZMp4umww5I6pEsZli1sFJ82nvc8mnZRNtnCwDAWF4pAE/cltxcwFdNR73urju6PcVHFahJXKXUl0kFpczpY9Ha5q0pk5XjejIhiIaeg3wQzEoQA4acQBsNhxTUDy7OTOexuIGoITd2dC0wXehy5ZeLVpXlviUzfzenvueRuyrmdmDlrABgQNeun6uW0o+9s/7Rd2cc4d6RZ0e0rkia7Xmv0Wc0XP19HvilcWLpM/6ELDFV5qWBC9mJvlJlTo9dWUQyItCgszDTqLrfJQZaSvKRJHZTrQWKSFXri+X4StNmdDwdl3Oe5OmcZhOo99ZjLDwQuzbwXhjaOjkaa2izYi5dc6F0d4Uz14ZNuYDR5vSQ649jyNleZfIO/ntslbTQmeGTC/N4+0vq0cg0WruA3RrxGNoGbzP0BS+RbYyE2qrLWzbtZIuVLNWd0cLN5k23N0aLqxjwd1wZgLL6Fc2WEir8LVTdOCcnboqT2fI+lX0+GEBXQhikIWs+jgXaLOaWPXXTYNOy7OaZV1Q4QsZdUy/DGpTwxECynRqLXXUVJCxMtZcJqBE96l7+eks6QpTy8VkoHmvbhsLcScoakNp9N862XD2uRCwIXqt94ggmi4G5S1V5Ss++zdZ7cucUaq/tO0ociGybdI46OFTqM1o4dnKLRpdAtBFdp5hrm2XocbQsKKu2rsiKD+jwOVwWx9ZyvwPZxInrY3CZAZSVFIKmaZ5nl8ErydzxvmwGkz7tH4pYopW+o24fc8vSRbVZGl6tkvwtzOly1o/Rk42VLvh8vPPogV/k6C9jgvoFvJzT1Cdzekz8JyHVIP1FoZgHyn1PI+gGZ1oLPTUTgGysmtOmu+28miBVQI93ILorSVLc8RZTDczuvz7IQJl3FuSoionCRHI4qKGwAoJ9Z+L/ZOepdtU+LquBMD27kdXjrweIdJjd0EwrpObxhGPVLd/PfoHR42bLEt5/TPxO2xZ1VsQ+nhDXXx6+XUXh3z6WzHPk7aklGUKtuiSDlwDhsSMGvCkCU3jW3cNXTWO0TLg4Dm9HHeVMhLQNGuy26YI/et24U9G9ZHvPOYhNU50YzX4x+azpocuskcoqXQJVC1Wq3OG3fX4uUGJz+6NNTmhqyaJjSumKHmpQsxbTMbNcuiwEWviweYs5qOhm3lVoa3t5S6knxIFh/SgRsVuvldudM3UbauQT0pyd2Z1PTt0es7wXonLhxkauhbp0TAszj3ol848FfxcbItezjcf04bIV2wzHp+YruTt+iJjtvOPorbDhz0Vp9AsaeH+gYYinkfRoQO5u1RS4B2HlP2PiQ7QS+wtU2nCsdjLaMayaH5dUT18DzU0kR+ijd6A6Ikj34XThYGh6XDudrltP3J7clHg3yu0Ink8c9znro00eW0Sysz8+/ndct9As5fpqpt8L0ilQ1B6SulPnGhUay/jY3eocis/so3z/Qyo2jRWrMLEUCNERi1pT6VncJykiwZOazHYRfamljO3l+pq9Zfsus9YRfyFgPY2U487yO0ra+4W9qt2+SGAk8njmyqhJ3EBtHUzt53q6LIzpNR76rMBroE6ofCY5MiY1U189wK14nb84XfXFVPN+d91INOFJ1l0or68YB8qXuHKdqAxxlF6Vmyze4MI+j2PtTnX7Yeu+gg/ldwPo+nsP33seclXtbcNdh8hrWePzjTYdar22+hYeuXeW6shCwcLYWxzcnSZmrqtqyzzukTNFtIQ8VqjZ3ND1ns/o9kNGz0MZS5zbuG0uc6PB0DwM+Xp+l2Kdfed7jWMWNlKsFJ27VPuMqPXYVUkV5hzAS8P/8QAMhAAAgIBAwIDBgYDAQEBAAAAAQIAAwQFERITIRAiMRQjMjM0QQYVICQ1RCVCQzZGRf/aAAgBAQABBQKy8mG/y137KL9n0y3nm8F3xj/ltPP+Ywf5p28xaahV/j9Et50WIZWCZXS8oxWLX4HmowP21WERLMbhZfTXCKOHPGSdXG361Mp96/A9FKk9o8vs5+fgeevUkAoxdjpmd9a30/8A3/r/ANj+t/Y/rf2v6u37pfp9vf1fDhLxxNVt6mRxeCt9lpsgxrJpOO1eVZkD2jG7ari2rXrWnOPzWy3d0bdrkDYWFbZXms+8waMbpC+lmry75jW7oyvZMvqKcjeVnZWby+r1ptFTthHoZRc+xE/u/wCsfqNN7VahnoUpz1rxcluWUR7j/tt7jb3+3uP++3uD9R/XVf3QH7X+0PpqrrBi5lKHH3qA5rBYZ1HmRdaleg6jj2XUfy2OOeuKh/M3Uq1R81bdszBpijefhrAW638SaIMejTWL1UPxam2alXMmE7eGOu5RdjWu6/2f65+p/rf99IKKmffhstyK4HLcj3e3vOPkPzNvdH5m3k/3292vzdvcf9/6q94Tzs+3+32+1ymynRMHhqVB/wAxiH/PafyOr6iPfV+tb9mZejWpa3Tbn0/UtcyKDpWnLxrPe2q4q2W3PGyuxKPZEwc12rwsiqNVcGp+V/Y29wR+429wqc8vpBVZBsUMCGyONofWfcjs0Hr9uPcT7f7Ee747I52Xaf7f67RCVNV9mNn1jjrOIN9d1d7sXLObl2r18iytntfHxcg15DVCvUPxPSGqpsWxq1VHHzXxnd61DSqnH9trs/fB7GzO76gC5lgWyPQOTYXub696v+GDVvdYm0FTNHr7VVFb9WrFWof67e8/0294/wAv/qPkge/HymHvv+W3vGt5qx7/AO0+33+3TFlpP+awT/nsvA9uzcitMTKxK+czsGzCayvaaWtmXVk4S24w07pZea1NdWAUaoEtMYBrcVCWxe9uH5nxju6dovmClxOowhu99Zta2l5lEYqD1UhtWXCxxq38iflH6j+ufqH+Rt+6/rD6pfpj9SPkf9x2TymD4/8AWDff7VfNb+YxO2vjIqpyxh49r6bh41U1RaymrA476LZ1MfUstMWnoe023UXJn4tJpoYgV1nhioGrxqyasMe7xa9xikcIOQc+qkcWPB6nJt41ZMruvxbqsmpwcmric0hbktuyDjtwak8zR5DX5ukmxWvl7jjzx963oE6tE9orEXJnXmXYQn+x+DbvNpX80nfV8f8A9DkNadbmPbwmQ/Ul+J7VSMlcBMq8ZOFa9eJVXY+TnVgbsA0tYGzIvV7ci7rDOu6kuya1sS6neu2lxY6rOpvNRrIxKrOTpcpAb3VyvjmrJeVnEsg09WX8tSHTKJ+WYs/LcKLg4QnsuHNWooXAxcdSKaEi1rAsAM4byzHGRmezidFZ06pwqm1cXhG/laf/AEWGf8r54BYZlZNOKc7MyLqsfHyBbnNirYqW2zFSqmNZduPV95xMRdpttCO4jeYWV1wgE5L25VeVgsamq8lft1DUap5m8rAzFyrKDianTdNlI2SHpzemGygTU3rsxMYcBXYIhGxbadWdaWg9XoLOkk6aTisHaCZP8mv/AKLTF/yO01PVwsbcw1P1LrsvJxasZFZxygrAHDswVYe8MPYltoO877Tu42UT7EtHrV5bp6uMjDYxcrIxwbMd4DyHEg0XGY6Yt9fs9M6FM6Nc1OsLi4g7Jwg47doGnIy0t7TzSdRZ1JznNpyeZJ/yn/0ODy/M/wAR55rG/GVhiUxVVIvy/uDszP36iz2nytkoofNr5PmoD7cJ7eOAz1n5mAKtVlGarObg0U7pz4qXOz1VOmVhcgtQEsfg4YccfJarJrsW2rw1b6TGQcVUQAbDxyfmzvO87zYwCZf8v/8AQ4JC57g5JXEtWU1tWr8zLHuVfaHWNbkrDmXIbNQMbLucomdZFwLzF0houl1z8pp4rpHJrtLvR3xMpC3NYLBKcy1Di6lKspLG6yKA44kkS6pb67RKUdbBQzzShtg+GqD9vR8K+v28cz4P0iZ/8yrb/iK9ypY8F6rQ37GzKbezNbazOYwHKuiaZe8TTsVJX0wGsNZsy6Nn1HFMXUcVZ+aY+41LF5JlYjQlCemWS3Hw7Y+m7y7HyaYlzCU5VgFec8XMGyZO8yxfc6Y+UDR00TTfpvDUvkY/oJ9vG0cl/VqqcdRq76/qT8LesDLLeItye9dGTeadPSJTi1QOdrraVZsysLblZFkPebeO3juREyb0iapZEvxr4zuJ0cPIsv0wRqcmqVZTKa85yUy22st5Sh9ppm/sfhqnyaPhHqPExPX9Wrn97i99d1z1ZiIQ1jVGmudbaPfjqWzDu9ttk2/RtOM2hEI8Npt4bRHsrKZZi5FZi5b7e4yI+CAiiysq007GsyrlVUTw1T5eP8K+vi3onp+kkCayynLxkCa0cMZS5CWU2ciHe7ZDyacIFgWcZtAkWuCqdKdOdKGmGgxqiIVP6NoRBuILH4iwCdVnlFe80mtatP8AHVfgo+EfoaD9V3rqv1yM351hnzZeNTm5ur6bbgZFi72ClWToQUNOi0FDRMVjFxJ7MBPZ57OZ7OZ7OZ0DHr2jqJYVEZQZtNjFWFIR3RAY6gAATBdt8VDVieAmq+lPwj9AHfJc1436crq8tSGQc7UltU6cL3rHVTV/xLhrmaYgByF+HabTaKsrAlYrAHCbLFAmwh2hhA2sC7Mi8kxHaHEKw4xM9nYTpy1PKHi1i9banqs/C+nFofETVZT6D1/RqltdeH+m/wCL8QfUWO5v0Xf2f/8AWIDDJw2ws6ufcLvFp3lgIiu0UkGtorwGCN6ZV1dKvqQsOO+eLGqz7rH093FmkVCNpcswMxVU3dXJ3OcKawekldeHTRqNnYDwA8NV9aPQev6PxZ2zv05CWMNVWz2xMPHD2dX2rHwXryTPxZ1K9SwwbJkFcUfmdcSzU7R7JqjwadkqcinNtj05Intefirg5NeTTVxK52QtOLjVtlvUnGe11hvamgyLjKslXjEGHfjkCvl0q8i+qvhfnIWp/DvbJPrB46p8dPov6DPxZ3yv1a59ekQb6lM644+GlQuybnsrbTsWu2ikV1g28Z7cC1+pkEalufaQYXQjTarJTpFXHUtPFaVsa6S9jL8R/aMFrBoxkayuEbg44Y7bRU3ycgeXQq/3HgPHVPmUwfof4PxSPMR+rW/r1+JHA1WanX1dOw8ixG1HHdGxa/dPvW1nvBaLIWqrtry7bqlrZKOmrJpm3tCHtlrzrxUZ0txDYwww0TAQFcRALKK1HCxZW05e/Msco/xV6TXwT9OqfOqi/of4fxQPIyxlhH6NdX9wu++LX1NanrMqg0ZidpgHZL1VwA29WDvBhLsMVFllNe/oNNVuoD2cx6h7V0+48s6zRrHMCziNiBtWvIzYFSNqsReON4jw1P6iqL+hvh/Ew/bmARlnGD1++u/MHrhfy3hrdH7hu0xhxc9W2dAhU6k2ti1kziJleWpO0TvLBHr5hG5L224LANvGwndVCI42ZASXlfyvEeGp/VVRfXx+2q4rZdO8+28f0PYD11/5iTE/lvDUk54tw74lLXX37CwTbfxteuoV1NbYo8yR/QdmZVY2WZVR9scT2u0zlntFTN3HPwvmP8brKTvT4jw1L6yuL+oGbwjwbYz7678xfTE/lfCwcqrpV8D29Vmy2qg1CDKymipqlspw663fdhSO61yxNgw7ovkV+LJbOoJzjkS3cRpb3lQ2sImJ8n9AmofW1xP08fAGAwjeOs9Drp3tVTMX+V8LWC05PZR8CVgyqqsTyiFto9kZ+xOy40VtpZYONhG9fpx5ux6cS0NA05dmaGbbxuzu4VMftX+gTO+trieG/jt47wGestWa38xZjfy3hqZ/b5Y8gPkQxWhMseFizBBxsuRRVeN0uBW2yZOUKTTnmyY0dFYW0vUUt3nUhbwUbzUmYUB7epp1vuf0LMw/vqjEM3m83m/6CIDFMcbzXhs6AbUfy/hkV9bHc71nsEPYNHsjNzNIVJuJkVcocYLZ1WWC4tEqBleONkG02JjVki6vYo2/gB3TtNSPuDWfaMA8bB38RA8sxq3uWlBOAE7TdZyWbiXuUUN4bw9/AGfiH1X4cb+XPr4alWoRwDF7TlsNy8oVmbiAOM22a2oOLKdnCDausbjYT0BdZlZVdS15JsKV+9ZNoNy3wzNJIuO1h8tgPikPqr2mecBqr2e57K7a/aHCNa8Zba31T5W8VpuIDGm/fXzE+DH/AJc+vhqrft1YlOMCx07VsKlF68xYDDtOPYKDFqCNY1dcbPoSW6kTM3U8maZpduQ7AKESXHaJH3mU59psO4qbeqtxuO03lfo/pjx/XYxawZR8vD+bm/O1g7LNuwgac+w7zX/jr9Mf+Xbw++TsbKGg7xOxVdxfUjTM0+9hbfqOO+JXk5EqwcotRg5rquBltbdo9aLZplaJmVUVtRhK1qjiPU+gs8zDsHG5azq51hPRDbIlldkrPGCL6NMQ7iv4p/wq+Th/MzPqNbPl28e/gDNf+KuUn/MTeWNxryOwt93cnp5SKSRLU5RPTIoquFVXQbGy6urQyNU92PVZdnY5TMtNyHvYs+wEv+Gkea47TJt4VYq+W3y1Y7eV24tj3mI3YHtML0rgjfJr+Rg/Hl/U64fLuZsdk9eO0bsdzvrje+qIiELrH2MyO5+POzq/d49imL2IPFuXKH1b4fWVr3FSueBjqDLAnLaDbcd/C5pWCEtsHLObesDaMOVbAJcGATGeYdu8X4bNwuIuy1eDfJT6fB+PK+q14+Ud5z2Udhy7N3az11v6muXKPzAnwf6jGEzByorOzIdxylPifUkzlsCzbltiW2nMRYO0bsp7vc2yk7y075nGVNujKrO6lTSdjVZtLLLLIuo5uO2mXpkYtXg3yv62B65H1WvShfLZ6ld5x4weWWWbzWj+5qlv17+GQeMB6eTlMOTt0bKrjF8y1+Vt+3Zg2833h237GPvOM4ype8uPY7KLH5GDYak1TCMem727WPuYF4zn7ncomf71fwm7pdV6Rvln6XT5kfV66fNkZK4tT6tUsv1y3ddcsYW6veVGp5/K+98i2r0yjvqXVJnUljck3NiXPymotNPulNoMI2lLbqnq1fI9OdMGLSs6YE4ziNxtsx2jney9ptEXc6gu1tOULqcn4bOyry6PFSzdsO0t061D4ukdSrVK/hj/AAN9Lp8u+s14+bW7TwLFhd5TU+y77vaw3rO9te0y+2p1nsDOaidRRZmWIa859xS3G8HY12sYh2NR3AMJ2K7QeDnsNoZa20d+02gXYZ1e7FnruqyQ4ZN4G3ffavIPJTSbnycbMWYHNcvHblVH+Gz6TAlv1munvrBacVlsr228m1gQQEG1PTN/kK/T7ZFlvVrewxHJGS+7jfmvcI20RpU8RoO8XtCd5vO0JOzN5LbD4KIo7he2Sm8uHKxV3VfUbiDeCpZgJu1TcS2Pi2GusIsf4bvpcCN9Zrvr+ImBXqKJ1NwjiLYjTIZAavmpNQ/kKiNh3R6y1xqVJdZ5W7xW4jG82PFfaV28SLCYlhEFu4VxDYIziO8e7eesUQStIRNUfpYqEq4BNiBR4c60nNrzRXxBXtWdpv27777xgXrx0NU3/ea53llL3WLQk9hUj2NElVFGxqqJKBLF9Mwj247l6xYwtqZZfsI47it7DbRwGlsHp4wibkRLCsGR2GR3XLQD2tY2QsdyxCkxV2gErQxRtCJr7HiTMcKwJZTX1LXow/eUVADbvGERoCYTAe6zZWmp6R7SBkdJva95Tmtu+aNuu25dtm5G1JqB/wAjfdxdsm0D20bpfziVbFzxXUt0p0ezjdx3jJGWcZxmx8QpMrqg2Hgiyte0Ppq9vU1CperbXjKo9jEoorUFREn249tpt3BjT7b9oDLr/NU3ZWm6zmOd16VjK1Pkb9TzLCmdkK/tfUOP0nfoNypxlRiw34HjrS/tF3SYV4vpWPXCkI8ABFXuo8BFG0qHde3hquZXiUtuzaFjAkIJtyLev3U+B7wCfde0Xv4Dw+1yg21oplFKNYMJC7Y1YYY+ObbsDFU6pTTWs07CuyrKcarHBBnHYAeGqLzwKlDvjW2YWTj2LagliRhOInBoVeBXgHYRd5VvDsi6jryrEZ77z2mFX0cM+o5KWbv6RGn2/R38e02ljN1a2siWWq1Ft9j5SZNVen3XHMOSjzVuJTSdNOQUWtK3WAd7IgjS5eePWJqmMb6MHLsxmws2jJAjpChnfceAiy22nHXK19QMvKycppjLtVi19XKbvB6kkzhNu/cT1EPcDfl959x4Ad7D71CsJWVXcGfMF1GPWqo1NVYqxA7eFhEQdrJtGHcL5bV4WYLEY+padzg7HG1PLplWuVmfmunmPqOBDqWJPzWlZZq7S3UMuyH18MXTrbqePbRqt7OO52EUdztsx2KQCDv+n77z7mWD3lU7xRvMbs1ClrlRQfuIvp6ttD3aMJxmrVccjTu+Mm4Obp9GXMvBycQ+O838NptMfGuyGwdKppmrW71t6adVwxPAQyzaJ6z7w94YPT7QwS/fr1q04NPNvRTbcaa1qrEA8ydo/osY+VRuT2B7wTPoF+NoR5IQOIgc7ZGm4Vxt0V4dKzhDgZggwcyLpecZVouQZVpGLVBsq/bOfnl49XUyG9dptBt4NB8Sw9p99p38D6t4LMh1F/tvGLm3M2Nhkj0UCbT7iWnuvq0BCgerwRpXWKtSB7/CePfbaH0WeabnZZv4Ay1uFS7l9LQGyAT7eHeejD07T1npCNgPT7EQRD5WbeVo1tmm4dWGvLuIsB8Ps3xDePFHcdofRT3PpmqeOHeMijvPTw2M7ztNpxM7b7zZpn9q9u+AvDF7zad94R4H1SDabd9oDxGw2j7w+tR2P//EACoRAAIBAwMEAgICAwEAAAAAAAABAgMREhAhMQQTMkEgIgVRFDAjM0Jh/9oACAEDAQE/Ad9L2JEXsS3R5Ih031ybIK41YyLljEaLESXzsYlip5D4PRTd0RimdTSq0aqjDeMjeDL3QyLE1ra4oWJqxcvovlUjpUEynI6X8tKlS7a4J11KR3mOTeikzvSI1v2QexcqfK3wekFlydmSrXKkVZxKMFFXuQs5uSLly6NtI2PtHgjX/ZdSO2ds7ZiWNyaerYxSsVuqUdj+Q+ImMpeQklwX+OVjufsyIt8oh1XqQndaJaz02HYlyVG8dhUpye5GniYlrfHYa0uRmcnSyle2keB6S+E+dHIdQu2YHbMEYGCO2OBY2FTiztIo7PSPiPST+FTkYxISLGJiYmJbSxJFiO5ThbSPiPR6z4N2RXonGxBFjYjuYmJJWJPWWlO7dhLbReIxkttZ8Cq4JoplZeyLVjMuyNRohVUhzih1XPgsvZjFjpxFBlKIopPVeIxnUv6/CpyUuCaIQ+xwTm4kKjkRX2JQuhIlCTZCJYUSmraXLnrS51HiRqWIVblypyUuBljEcLip2IQ+xbYlCzMBRGixH4etayvEsRRFlTkpeIxkON9YFySMDEsSI/D0X0lvpYpz9FXko+I9I6MysXMtLlxs72JGWSv82tLEij4lrkuSMhyMyTMmKRnsZF9hlTyOne1tHsXMi5lYaRKBiT5KPjpU5Ik0N4iafsxX7MoF89oihbWfJQlpI9GOk9G7HJW8jp/DScfeuxKjGR/HZHp0WS2WjJPYfImUalyWtiY5XZsxFbyOl8RokthPRPS5fR7aVONYuzJS9lOWUb6zIwIpJ7mKsdR5HRx/xmJOOxJ/c3IvcWrHuyxNbFjHTF2Om409lXgjuWF4nU+Z0X+orza4JVZHsXGimZGQ9y2k+BbsxIwbZRo4xHSiOl+hxaKvBRmlyZoVS51n+wo1MaZOrcxcycMYlKV46WQkWRfS5XlZHS0nPc/jlOkoltLDiToKXJEUiU2itGT+zIOTdkQo/sVNIrw/xspTxE76IuX0lUUeSpUzZ0tPt00v6ERkVbuRKEp7FGgqaLblipHY6enk3Bks6DsyNeLO7E7sR14kq8vQoymdNT7lVI/8/oQjHIjE9ENGKlh1F/2TpRmrSKv41rwJdNVjyjtT/RHpqsuEUfxsm/udVGHT0bRPxlPdz/p7r9FCEv8AojHRi0ZVj7Iu6v8AG5+Tq5TxOjp4UlqxfH//xAAkEQACAQMDBAMBAAAAAAAAAAAAAQIDEBESITEEICIyEzBBUf/aAAgBAgEBPwHA0iUG57EVu8FSHkReiSNOiWRzzwVdUdyMtTIUxUiTS4PkYpMbZNNop+ou1iRqY2ylhw3Ka8mavPBXgyjOWNz56kK5VSkaNEiksonwSyhWbMkHmyu+2lLGzP1lHSx4K2FEowztM0Hwp8iWLNJioxJ0V+E9mZKUt+3JnsXIlyVW4cC69fD8ePIdSU3Fo3b1YJPbFsGBoRODW5HE1uS6f+HofKj5UfKhSzuSmjWinNGTUKRDdMrexQ6GVTyZDptHAlpQ+5Gn+DjkmlwyfTfwaads7DtTMKy5IcEIx+TMjXE1ocjk0GDCNJoMXccmDqYrmzd49lL1NJGm2KilyLSjU7bGEaTDNTRqTMfwk5IdVlXiz5uuyj6kBJ4MGyM2yZMmb4RuVFhlWWp2fPZi1FbnA247ooVdccknbHbjtrNKOWN2fN4rLMGCnsyOCrsdHLfBg0GDBKJgwomTUxVGN26mTzi75FbpPcxePBX5OlWxKWxkiiSspWTQ3bI2V3mV/wBv0vsTppkqRghwVvY6aWHgzbJq+mXN3zfp3hmRsktyPBW5KbxLN9Jgd9jxJXn7d0HeUSPBW9iCzKyM/ToyiosS+uryLZ5IS1LNkYti2LO0ODqY+WbRRhGg0oUM2UjIuCrzah62ixLNs2W3JKeborwGiNsjKfNkrR4KvsLkpS/L5NZg4G83jZrJWpYIod6XsY2uvVFfkiRfl2Kz7I8itJZIR/GVoaZYvR9hD3vX9hckfYivGzOO6PYjq1vejycW/bVvYpRIwzbFnEwY7I9mCcdXJKh/B05IorcnqZmfA9RT3iShmRCiJYIklvfJnspLJLtxbTgkkNYIpZKTgtkOMeTFkSjnvjTciFPSSe/0SRKJQwluQ5G83Q/6OKkOkzQzRI+JipI4JPC+mSMFOAl258TIqn9NSMo1IdQj5Mqv8+lUl+kqcfw47l30kTeX9P8A/8QAPRAAAQMCAwUHAQcDAwQDAAAAAQACEQMhEBIxIkFRYXEEEyAycoGRQiMwM1JigqEUNLGSotEFQFPBJEPh/9oACAEBAAY/AvMvMokqZUclMCVV9K7b7LtfTF77zCqU9MrVOEwgXaSiGXumO36FPa4brIO3I7Q14puZ7flEF7VsmfZaH4WVoXe7phdy90DJmlZx5s0LLuyps73psNFk4/rTr6NTeqd6U31I+lfuX7V1cujV1chyan83KmOATz+pNOzfihmIyNfFtMNFotFnPBNpNuqnRdtngF2p7rSFbB7T+VFjZAUEWVSoWl8fqhZWBgj9Up8ZRHJNNpXncOhUZ3/KuT8qomI80MGuAB2N6Ld3eSv2L9yPpTHR9Syim9f05YfPMqo79Kp8yqh4BU/Un8mpnNyf6VS9SqcmpnqTvShzeujFPF6LbJtIzMyrUl+GF5Wrd8Iv3Qu7cPtjvVT0rtcqrlUOsggn1qdJodHBQqjqp2WGzZ1X9R2YWZd6fKA5YZuOBwnCeCHpX71+1D1J3pTe8ByzdQyk6eOihjHA9VUkfSqI5qt0VEc1VPAKj1VXoqI5qseSodVWPAKiOaqelM6oDcEX7ty98AinMB1CFXNMcE/0rtQVYZjlA0xCdJGnFBrdSYCNOrDGu1lVWl4+1YQ2N6eEzosrk0zhFOm5xPAIMHZqknkg2pScHHcr0niNbI9UPSv3o+lM9SLeSgYW3J7d5Cojgq6oquqAVcqgFXK7OF2gqgFXKotUOUBFDpgUCOibQ7ppDz5lfe1dq6LvaDyx54Inv3Zxqszar8w1WdtR4I1uqdUklh2XXVL8peCEztDdRYoUxfiqgboqfRBzIvxX9O90EarJ3eYNbN1sMADWrXQIfpbJW85inOcwcAs7TpsoFnmFy1d8weXZcqfqT3+yiFotLpriLSmMGhuq5VAclXPNUByVY81RHAKseapdFWPEqiOAVU81SHJTgU1HAJs6tMhD0rtPRNGcNa1V6RZmIMBwQpM7PL3HUJrHDzhFv0lUQRtU3ao0zeQn96crWmNYVFo+oSBG5NcQJHJAX42VV/sq1Q+VVXaTaNVVrHMRp5VVqyb2GyqMu56JkuaNrohB1KzFzruXf0I7zRwOlQIns4LQDJpO1au6f9m/nvUuIXLorNWYNVHoq3VUvSqvqVL0qp6k3k1H1Icmp3NybyanccybwhFw1nReaF7IIoYjoq/REVHQTou9qMkuM6rvadMBybmY1x5hMNKA08k2WgEDgjtNznQIvLy7gCbv/wCAnd9UFZxE23cl9rAJ0Eqo9oP5fKjmzbXBy8sE8pVj5v1Qj9JP6kRcEn8yec+gy+dETORv501ppulozeWVTHG6JMiSnfaQ5psodFOqfg9F3VUfOhWpngssO+FsUj7lCs6J4JzZ8xlNdOgTmTYoOziQiC+xus3eCeqjNZTnuuPsvL/C/CPwvwXfC/DTnZSLJqPVHomooJqrdE1tPLYb0MIXdmnmgoUKbu+r6bO5Pc+pNcVg3VNYcub8o1PVCo+AdTA3BN2mtsXnIJn3VOlJJ8zsx0TKQIaN8JjNrLN1lbTcA03ustNlQwfZQWvMDdomgtq8XL8Rtz9QhOcXNBOyIes3ebIbFwmwBmF7KfLU3j8yyVWEhCl2kmpR+irG0zqvzDUELacY/lf37mng5sKR2pxHJXr1F+LV+V5qp/ctH/6l+D/K/AYnFlJrTyCFl+GF+E34Xkb8LQfC0HwpqO2aZ0XmWpXm/lf/AKtFp/CaqnRV4bNl5F5VFZ2Z+6m3Vd20tpB2jG6lOYWmk36jvTP6do7xn/2IvM8yVqXvI03L8rRwRfLj/lA6hDhwVwVF1YnmuaglWYz3CEOM9U1pyyLSQtuldu8L7WXcL3X/AMckjgV3XaOyewQf9DvKYRkm3FTTeQP4UVdg8dynNbqtQvMF5x8rzhFlNwJlXx3fK0XlVQsJ2tQvM5b1v+VorDBid6V2npgaHYiHP31Nw6IuLi551cUHtqODhcEHRU6dSptRw83Mry5+asbb1a61wvYIQoutVvupWq2grCChbfh5RmX2YdT/AJQlmYD60O9p/wBRSG/eF3lKoPQ8QVsx8LVARPuszAbaidF+Gvwl+EFIaAZQlaBaD4Wg+FotCuow3/C0d8LyuXkK8n8piPRVsv5UOx0n7TvxI3BZQoAvwU1dqdBuUuJiLK/wspKgG6j+VlPXVRuCaDp/hAiOi2d+vBRmUSrkK0nqtoLUZQbymw7/ANK91mbmUOacvFTGUzuUviqOQuFla7KeBUVGuzLMNr/0hVpgnjCbUZo7H3Wi0Wi0xpH28bUfSu0Pdo1kqpuz3V2Enkge763Xk/lfhz7ofZP9gjFO36isxDZ6rULZzFfUBzW3VhbVRwUd6flTmqqBUeFAfY6SvKDC26ZW8IEPlbbjdSYI0JJR2uQurGed1ob8kRtB3EL+nrjMRx1Wo7rfa6PdOyhBszBxHVDxNPB3jCPRdpA1fDVlBWq1IWq85+VZzlstcUDUqAe6Gd9+Z1UUaRPQLayU+pQz13v9IX4dU7ruX9s75X4H8qYezo5QKz5KPd1qWXc0j+Vt0JjeDqg3K1h+PlZqVS27mtphjitSFZyjLZeQ/wClDZgcwUx7KT5CBJa0cCV3bTLjqUfXiOqHicPHRf8AmT+iJXErVWVhlHNHO+SoGVx5aoZKQt+fZWepXud1NZaPZm9XXKu/KODbK5nx6lWqFRWY146ItZUdRc7msuxVA3t1+FD2Gm4qaFXN/j5V2SOS4YaqZUysx+ozi3qh972cKr0U88Msr7KnLxvKzGuGcQp7yo8+pTSYGraefv5Y4j3X2rG1ArEN4B1wjnGb0mB/hbdOmP3XWdlTLfSZUPwyM0+o8E1jfK0QMWofeXXZo1VTmF2mm4jaGyeBTqFZpa9qkLgpdv8AutPBp47SFe63gqCrKnk+q58DEPDHjC7OiS0gQnodnrtsWWO8LuqsEG7XjemtXT7rRaLRaLRaLTxQrqxUcVSpHVrb+Bib4qlRuoHiHdtaepXZw402n5TT3zLclmFZv+lUM2V88BCcfrpbTVPx47nDRaLRaYGVovKflWBjDSV5CtFKi6kGSoc1DttcbA/DHHn4WIeKo1zwHEWHiCoXj3UX144U/Si12hsVXpu+k7PRBDDeFDVdaz4s1SoGDmsvZqFSseOgQe7snZY4VDKFQ9ooUI3UaIX23/Ue0HpZf3PaD7qB2qqFFPtTnDhmIQpuc/NMQSjTvBAsgwFpI1i6zEgQqdM+UOkoNaAGiwHhYm+IRvb4vsss81QbXDVm7oKlRoVRRz8pQr1e0d4RyVlLvrbZWEoGs5reW9ZaNB7yvs+yMpji9bfbWM9IU/19SeiBqf8AUK7o9lbttX3UvyV6fHehUZpv5J2VzYb5riydWzNIAttI9q7US+dAhoBwWRjS93Ja0m9XLZax/RyvIPBSESm1qkN3SndoYCGkQ2V74N6+Jib4qZ/T46GFGcKlVuoFkx3aD3s657pzG1HASREoGuwRMiNSopMa3oFcrK1rnEcFEtn5UCqw9RCiozKURxVaix7g0viJstp38JwZBBHBMEblrlbvKawSylxGpRYKdQOH1ZplZmbL53LM54zj+VH/AKQbCALQRrcKFKKaeF/Ezoh4qR5eOioVGcK7N+Vd1w0We5aSmR+VWULumiGzqmwxrgG6OEp1Gs2lkAhuxCZlh06tJU5S0p7eBwIhSWAsYb2W06eSGbN7Ke7XlACsNpX2ltMhRuy4CBqiU5/t4m9EPFRPj7O7AEzsNkYQd6J4OhXQvEWWl+QWQX91LzJXlar92PZS0hWiRyVas/V5WmA4G4UFWlbvlaqTqt6uqj+eUYFAJnz4h08RVI88J8PZuuB9GObc6+B+VyWqjPhtF2GTfUOVQFfCN4uFfUa4aDwZGeb/AAgzcEVAQHBM6eL2Q8TWNIEHwWx7N6sD6cZ3tum80wNjhdFjTocNcftHhg5r+oqWH0DwXWYWcv7dtQfpKv2Kr8LY7HW+Fs9jcOql7Y91ceCU3p/2nZuuH7MXt5IclwKDqkTFyLSiHUWmeDl/bVD7r7PsDv3FXNOgOS7yqTVqcXH7y1/D7/8Aadnjjh+3F7idybjdoW7C9lKnHTE8MNVr4AtFKHiP33ZvVh+3Ec3IIeGFChxhbDgemFl5S8rL3RHuirhSwyOCvbwWUt80qHPJQnxO++7N6sP24up79y3yNfDAXPDjKluyeShQJV2r6UP+FAWq5+CxTbXzpqNI+I1C8yVYlalarfhvQcPH2b14fs8HfZdreVaZxsiGLeitFNlAWl+SzwflafCkSVOZXeFai/KtkROFsGidVTP6VTqjep44lOPBWpj5W4LOakT8LIIKtlUZm/CAe4EFNHj7N6sP2eDJvKjNbCVYwolayhrhqLLM6VwW1V9pRgSrUisrYahX7Y5x35SsrW2U78Jm5wZTnytlN9KYTuUKDgU/oivOV+I5TqindVTTOnj7MP1YH0+AkiwsFB3LNjtBq7zstctP5VlqVKjUJqOk80R3r5CMdoI6ruqnaFnzOqHmmRYk3QptDeK7x7RbTkgBhdWwmVUfu0CJ9gulio0O9eY++BTkeqd1w90UeqZ0TfScNfD2b1YH04kplL6nq+i1VsJMrREPaD1QNO2XcnNcHC2uWykEarPUq02iNSVNJ/ecI0TMtU03AzZd4bv0zLXwTvWqJOsLMd6E+ycCd0oOGiAN4wKcij1w90Ueqb0X7Fqp8PZ288NdW402fmcn1P8Axiy6KLqYWmq34c1da4XMjorlWMqw8EKZXILLveVlA0Qd8oNIWWdn6TwWUnosu/gjxT8D1wCKPVDov2YQp8FDrhR11xpe6qRxVvMEYupV/HqoUwuGF1r4Gj8rZVkRE8lB/LbgVGiDpELPtTpZE069VhPKyNOo/vOIKFWnocQj0wHRO9ClRhdZsKHXCl1xY/g5ObuKncBKvotRCsFc+3gutJRglWHh1V96OEu32UKVINkJ0KfCLkXLvWi5sn0HWDhmGLUemHsn+hAuBPRZjRqqaNFmX9Rutvs49ioZQvzK2ms6QqLntgygmM/LdaYFvFNMw8b1JtYhdAsh3YZlfwaLTDTHReayytNsJUhAO87dV+cIxc7pVNscysw3hTNllkhpRa+3BUSDvi2LUemNX0poK0VxqojCypdcJ5Y3IWvNS07SPwmdVIUICDBXRa+HXHXVcz4JRIMQVlfqhDp4oQNFbVNa2AslPKSAialF0chZMLPMDKDuODUemBVb2TQ9eZ3+pRrhcLRUo44e2L4cruUlyncE3qgeOFkLdVGNsdVZRbwu5Ko4EQrLehmVl3r7AJ1Q7yoJKzGizMN4EFBrTYYNXtg5VuqZdWClbkNFuVK+9Be2Ltl3wpcTPBQLBckTvQ8FvlQtcOCiVwC2beJ5UrW3BWWYwg956BaWC0U42V1k0UOTlV9aeXF72tNhK/Dj3Xl/3L8H/cpcwD9y/Cn9yow2L4SVZ2VZRmd/hbeWeS3TuAUuVmps7yiODvFd1lcLeQthp91fwStMAwHUprd8rNCtdXWd0k81pgcAt3g0TjRrQ4mYcqwynXgvw3/CjK/4VmO+F+G74VqbvhUZaRfCwmyIcYg2WZlYeyl9RxX2bYneVLgrA9EHRd1kaZ+rxaeC608TyJys2UBFybLJwXBdVGIV8BZaKfBUA3larVarWyzFyZ3TfLvV67hyCzZy7qj3rA4H5CAbUeTuaWrZaJPLRS7aKsr2ngmkDRyDwbgoOGu/C2GmGmO5WXPwbR2z5G7yi52pTu0OtGy1aIBWx5eE3xvg7XVfV8rLtfKja+VEO+U2k9s5ua2aalog4bIin9TlFKn+46rXCyhP5Qgw/VZFp3GHBB7DIOOi4K0FaFaq+Gi0Wd7g0cSjT7FtH/yEWTqtVxe7eThTp2nU4SNYV9k4CBhIxmcNPAQIWg+UCGieq0HyjVyttzX9U8tO5ahTIXfVhFEf7kKbBlaNw8VQfmas3BN7VSG0Btc1LbjeFsuh35Trhbw6Lcs3aKzKY5m6y9kpT+t//CmvVc/luwniqbN03wm8bsNFskj+VLh7hWw4fcytVYosM6XWRjZCnIJ6IVK7WxubC6eCMeYMqqzg5UzyRr9lF/qZ/wALgVGYPH6lFbs7h6TK81RvVqtW/wBpX4hP7VZtR3svs6AHNxUGsWjgwQpNzzxPaKn2VAbzqei5KpV4bIXJXxspWYWP+VP8KFOOmE4HDRSoOhTu70Ck3d9zzTaoEB9j1UcHYZvw6v5uK+1p7P5hp9zlo0y7/Cz9oIqv/L9ITaLdDg2bTcrRaEY88dVOMYxgRK85XnK8xUMeY3mFlZPXE/cFm/cq1M+ZpmFphlKksdSPFi+x7Qxw/UIX4M9Cv7ar/pX9tU+F/bkdSvtKlOn/ACgXl1Y/AWRga1o3BXCI3NEJrOa1Xmx1x4+6urm3j0RVmKAwHkg/tGz+lZWgNaN3hC0w0vwUlcsNFnFhVt0cocYPHcVcQr4aLUrVXJU51dXWiL3GIRPFOqcBGN4UYaq61xg+Xj4NfAGMEkqfNWOp4eC4wyg4WU+HUox5tQhUEfqC2NOBVw5vS4Vi13ursK3/AAt/wrB3wpyq72/K3v6CFup/yVlBJJ3lXQt5r+C61nwRvxjVv+FI+fBC/8QAJxABAAICAQQCAgMBAQEAAAAAAQARITFBUWFxgRCRobHB0fAg4fH/2gAIAQEAAT8h235R0isQ0sSgu+4XmGP2sqXVkXm+P1N7tqUslzUNvdBLL2KzuZ9LxLAlosqW+Yr8QxzHDEtpdJUYVLtY6QV2rDUM0w2gIWTGYMO3Zm+MFs2QQWnOSUOvRwiYNFOY0JPCvCXK2pyTI5tMGiA6TDgeMahaurLW2q73PxT/AIcxkTEr/GV9Z4RTojWWIlhgKavhB2gZmtNW1M5QgraWVhj9kdpzr31Ka0Xj79rmpg+ssz3SRiGDCV0niYSyiocGVZUEeSqCusxcL5vQyfmbBYqcadVqMc/NQylq5XOuv2lA6mMx1r1LQBHSpQ/M24EoRGFbQd4jaOPx+qFguy5cVupdDvZ4gEPwNqZ6Dcb8GISk6TodY2oBjrcZnp8ZGQ5QXjpVEyOpxhcBvrKFA0lIrP2jUomoCi04ETsHc+Y5CN2QdRcVKsKnIExDGWEdpC2tzzLBYKdnVJgEB625YXI8y2XqIGG8yhCYg21O0WRcVwTIzxNycankQjiPiRvlhZrzVjCBnuR0EltBYYeydnfIHfTMcKTlkGNH1oBRUHS4YU6yeXpFUPWb0+oy99MRT6xWfiG2qjHOA2st+l4jDRXU3UvHNM8XKlLZzcJ9y907wP375NSLsdsRNKgK7TrZzCkHLORG5U+hNm6mSWc7LTPxdb+5YueNo6SPUmq2YRqydJ1hy8T4qheLMA8Z2iKJcxUdk0bd4ljqszYjrziZPgn4kER5Vwp7DFCqlv0KlnWwvAkDiJoZnVuAww7m3jHkCzaIi9hxlwhfOQSg+dxft5hMfGkO5wqEvMZo16ojC3AeYhQ5mtxBHqGajEQ89Z/OGd5E0mN5wVpzF3KLtXAQQG2oQVbd5DUTm74JvopxV3lywRgxk0vyvyQ7woRM0cMRgeoW8MSpojW8vqMMfWclSk3QzK9ycBDxIfBx7KdSI8yaCvU5yKQO+qn0U7PnTmX/ALXKw8VidcrH0YH4QPvgrOy9JyeI+pTr73vHz37/ALiPWVufUJiluCaWcMdonHrCYo3KRstfqoGeznNBZerzqMJY2r1VzIHR9ID2QApiVRo90x4h0o5jl01DZivRrKzLlMxE7cPcoG5yIykaRvnERp8reI9+8FVgLk5rqQhh43X2gxCYMW9kDMn1CiKOhDbe8GacAh2QMXxEsMfy5ymViy+4o3QLWDaWr+E3ul4ViW/FxfaLTDzM415nA7IWcdwIbr1FU7p9TBMlwplvl4hFbyru6J0k9mqFjjUq/sgDoQ0D9RC4Bn8swWLvLHVr7PqVA2Wfh6S+N2nrSzDUY/8A3MveA/FRKU7ng9bmZh4IYlbfVm5QLWuF3NSNZ/0lHMVmDuJABqU/pEWjfNr8OJXKypizbKuWBr/tTqv4jmpXDwdpTAKSGport3rBg6uDBU34ERV3VqaJtuKIoalZWzC5QYDFximee46K13cGrI6QXYb2ihZbF16ekXdP6nTlTU/LA+CBPegvKpgD+YEX7njRFXcMGfssLyTMRdoJsYlYjSbqoaXimxcHdhKGEAwfkTI15UosnABAKlrYY90Vxo5d/jD2G2s0H7YEjauaJSh68VXxF/aCDVRMHOsYo+GysTEV3/WjHiZIjovrpdxYu+3+Jj9etHbvELHunJFf6EhdRFll2X2dRjKkESoQEmPSCPMpyPRFG5FDloX0fMP/AIcvJ3DBpbvMOk6WIhNfAkbMwBEc7zUorLHaImWveYA7R7I/0uUJVWalGJh4MVtMNtS+DZzEggSdw89JZMNL9F8zmBCKZyXwQOZtGnwRly6qOF02orKUgV0iO3dxG0GW0KdqDyOeYAYSBTI5MNKiYcanQbAYiNnvDCPdOcDp8bP2SyVc3dwKC7z+GXmL3dhPc25Q5PT1CPJ5r/RZRnOjoZ2U3dqVVn+3xMYL1EPJ906v3ROLaH7nCQa440fqaoP1KTEAa+iP+GU6vUyyHdOp1fxf1DrL6gfH3QDn7YDgepa+ZQr3mOsvCId8wfNfcV0Q/wBJYrSe02sSh/kiMhwlGo5O8FUufQTADzw7yyJ8pRyAuMxZLtOZUVYNSlXPSDYt+5t0aieY8ZwH2S+mRNVw7k3PRTlspOtjOJZ3wv8AMbWBz1QJVvIc3KBNAn/2gW83UPHWYQQVnA4x+4CzDssCTON281O6+52E6A+o056Eyz/O4gn6Ypyxr0R3H18RlKf6Z3vxOhCkl+JdD7ISHrSGJ1RQCLxOe7L6nuBAPEp9YBq1CToExos13VmFcgKwvAEFyFEO3tqCyynFqadOdQZMDjLUpGXc2mMrHliFADuytVetkwap0ghWg57wEnED9T8ucAf3Fe/ldVFdBdZloVwnJemMwhx8EkcZaN/3K4HvNR9dyqz9iDsEx08TNmV5NS431mQJobgOh9TgEHZCWzgesZ6H1PB+Nf6/M4xHljkWKIeElY+Fw5d+QIXwyttP7mALGscIquvaVra+kblfRnSzpV7RRax4JhQV1ROk8DcSX5I3FZW1eYEv9PWDfw3ZqVRIWuiKsHamX2l6lrD4WITA3XWA3Qac/wBwJndkX1G834dy/BOvlwrV9wAXLv8Ay4FXojJ4m9H/ANtA7aPK7Y/cHweZOXS/ipg5OPvPJA8Plg8OfvH/AGwjQadYyJWX3AsFm3rKl16T9ZGCp14bg8pP5KTJWXMPA3jKorTo8enTiWv9ttzCGt3Zg7V3iw72xG9e2xmqfFby4NvCqmPv3lr6Y/3Kaf3Ksmv3eG3zG3BWw3yQ57E6CQrr3WpSptqv+nNQe3HhmKUS0C/3Al9+ruaig6TyCru/zEtBq039xV7Ml/Uv0ZNsNHzb4fiqSzdQINfCZ7TD94/5Ibgj9RDYrj8X0FfYw/ySXdSZewry+nxfqHod6LfcMtDnxSxnmxfmNq3rsPoRqQOSeT8AqU/BMRJwi9z+Sso4u5tBvQ8fPaa6NW/wNIjbYcvt0QOMOQbEZWi6rlOV6sTIEqNtTxc/tEXQlhCrf+AU18zCND5xIcr/AOhn3KfSI0pBCr2mfJUwISe2V/1BUsXFVxsPI6grj61n8xa233nU5YEqVBQ+08flKmXwqM3A/aLNd8GZRlbr0QOB8Ufii7ve/sSZSnJR9kqufcRe/QgbuXMGA3XwA3Pz4J/ZCHy4W/624IK5hbxEB2LC1Ku0amOCyPPiZo1AKst1U2xWDRqP0nPUKR4hmI8RemIFanjK6tK9JjwTpPhguJSGZUT428QLl4GfyDMxUWDUBqYO3XMoprfW35NzS7z1Ze9eXrDXwRW0gxX/AF+FF9U0FraT7OE6XUmXhJYzoCH+5yHK5RhLhhnBqGmickDpMg1GeagTr6h2Tgp+JRqTjE62/wARrheYnVn1MwKNWh+pVwxGPFBZEYcB0zZ+yB5NNc8RUSp7S5cNs0IfJN/lm0wUDWH/AJ5jp4TdcQxeS/2jQs8P8xpQ9/8A1EQqCLiSqbXeH2TEmKjGBf3Ds4jT7jJXPeGx/EwQmso+t/E1aTsvqCuB9QiLR9Qf15YV78sZmnY1meVHDB6l0W9J/K6ihypf2oY4WO8qihiU5JZ5zvP+BFaxhBUXxv6oTic/G0BE/wCzdMHfhM0X55XKKOJu85QJLN8DLy+X1WmZV9spaLqJwZyPIzfmukft7rMtn4J0h/qKYr7lmS7js5qU6og4UtM5G4QeI6sPoZVrFBRnzK92lEZLJ6wQZ0c5+JtxV/ykwWHAbRVlYPLywF3EvSEKlg5Q49yjotAwHxV/9BUOPm41fBz/ANA1UyhienDK6hSHHTJ+JqrCJuxcq2UBVYxwS4Wx6S4D6Ta9EW+lKv6hyycn9sC4p41/gghuW8a/mc4+KCn1P1SXNV5zWhLR7YDy7y603T9yCNexDb4lwM+f/wAiXULgR8fflgmK6yywwXfhtVO2amQJqGDmXlQXH9SpuVDg59yqwek1L1C3v9qL5GEf0T8Sbbhr5VF/92kz8MSuzUKm1cdyobnOsZC2mhnzL5ngpi5mMFMD79pU4+CCf5J/oidsDAez/CZrPRoG3vxMQHDUW2CFwq6zDk9a/tCTxrRGEsWE1Ks2IuQIcwdn7/oIEZNGUFcp9YyZrcIJQl1rdhDQ6zITFwO1UXoLOO4RfLxw9e0eZcX43TZ5Q2R/4NzqY+DjpARbsjuMbtc9S7Mu3pFIFCtaXrL2JQMEytqhdyF7hV33Zc/sNs8+Y1YbBRXcjDjulQOwOYnV2MczZzkqbZ5+C4LY6QF4tg9TahDEXx+UpR4yJ6b6hr31hx8dNkqtn4lIzyjb16mNO3+Zz8Hn45iiePCLM387Z9wILxK3EY4+TcA561KtPcBHoHEdwAOrBlbN2H4hUmVue8TbBabM7obAXWA8JHEDUu7kBUQisHqKY+2Lcx2jytnW45umJaezAZqXbi+ZhRiXc6cWIG0ZYqvmwSl/aawPTEHRGT4/9nqN3MjOqTKB1/b/ALc8j2+Jr/gXUOcRL8pfxGq4lLTMHJOrMxm5+Kem18m4KYMwbrRdPc4FENpU9QjDwsSUGz1Mv69SqwL8sZXf1nP4gCAB4lQUhcMxfGd7yirxCUqJOjIcP4mOT8fF/u8p92EfVidozqMnd9U/G/8ACnM0xoO0rBr55RzxyzKDUB5lFQDXM/hQ5mPxbhrX+FzKeCkU8q2vQ/7+YR1XEVNdmJbK0usXMrll5eO1EhJ46Z19waa8SiZKxM6VXSOygPWAED5rD5JRTgX+6cWvaBuXPFWf/ENCTPa5bheqYAqiEE7ivrLD8GERCLHwTCvEx1L+eI79/AWmCNkqZjWnxPBZ/wAAFOXLXd6NSgaXqEfF6j7dK8TvlKK8L7QV6safxfbPLApU6ZEWcxtYJoHNTcYFG99Z0Ilze5dobhQyERHUilxcxUK6zg3UtItHr/5Uf4fIahPPxT3/AMAMJDlgBjUGiYTWXAQAOCzWWCg88VHcvqVz7ZKpoPqAHJBq6dEGwblqmLm4FGi/MS8BnKRs8VhBVXAhBLMU+51GArn7ivmfkQ4qepfzoiMsljr4Z7+GcOuesI0wfKXMGVCm/iQEpgpIWDNQZ5l38IGLQ3LHC+s+4nFxMXuP1idZUjcBYRQLrtmQ6vaZy5l5UUw6OcSvL+LiwJ2yjcXwDD9ggOsRnO4ywWnMxpx64iBjLi5YNHjcPCfJ8LC7wOvr5iDvKRHWX0hBlmSIfCA+UcMefismtn0YdSHGejBhSpfsEpIByx4pWli7lgNoBtKAMByN+ICinnE2ZczxkO6jmlzl4pptvQTPVfRUMVi9VeoXVMuVmp2B5mBUUXjcoVcXLqZHESN5IsuZ6icgTolRFv0p/apjuO5HcfqNGotIDmdyHmYSkGJbjmbkPaLI+Q80yuAPTCcOpmXcQl3BDwgADXMMwWOpuWBFfUoB+TKHQ6M6i/sIUUO0QXLGyOivfaXipBsCu8rCb7bg8ZbaQV2HKqmCpiENL8SzQb8TqfINzidU3iy2r905+Fx6SuDdLlbt6gbqdKMhrogYyVcyCw7RDGbpIE9PEYds49Q7p1YdaE8/G52Qljoj3h/D3Bvf2dMwYoONr9EsLcaIrQvSURux6XD3mSNqqqmcyC95dC8zEyhpwS8Wg6c3KM075I4rL0ls+dcefSlsIGXRMvmASAOk5dek6tSHpcGIskHzUU7Z06xkvBJaghQSF6llgnKYX7qKCv3mdA866R0Q0FXVgrHNiE+lntODE5XEyiUdp2JqpDS0/wCNLp9xKyjy9WU6srCaVdRYmDjqBVmdZR0OjD4nseTEGrhwyEz47mHs6ZYD3AhWCmYOASwyDymRwsagABUoQv6hn38MWBj4m3RiVGrmpQjtd2PgMrmtiFBbJypJrF3ZOpzPw5qe0aFuQZxR9RxpP7vw2r8LBvYmBpwmLwytJk+BZJy2TF1qV7MaGRi0bdYhfNVd4LnccTO2nSBWa12zMBfqCaYQWIuYmwblpKreAWFast7mMY+EIneeLFVp0n8pkooDfpvcooDFkZqYaUO6Z7XMWgfmW7DsmsLzWojTCzmIf5rmfyq44WLXpgsNLG+GJxf4ICd+esKkYuqhPz5+98NMdb1+EfbmoHdFoGwSzQxKUmfiAaTKrvOYCrGqjii9UeQYM51PG5irsc8oedg5ZOOZoA95jbAO0ytr74auRMMCpdxUhTA5iV4RBYPkhiOwcI5k+4GnqSl2B2lGQRWmsty5cqvBLABjmMsjyZ5jEDR3xLJY19zstjLrEYF22aqsy4S//C8Rzptrii5BYUy4h6lic/eYM+SE6DmfivxGUdC7TZauIWkL18qnMFUCVqVmBioprCcGZJeSOk7QgYLVZfWWqTiHiVF1NUQxT+ZTYfNaj4tlibWKizOyspYAIdb5jOlMgaO5BFLeCWdQ8RHeBFiZHDvzNlCiYYu4mOJWtfsMcLmxhh/6joqyy/AxdaF3Qg1vQQQ6z1vaBQF4f6Y2mTJ0ek552wjmY8ib/Px7yYEhmaEAYiQ2uXo2W8HxB6iLlihqLk5npU+HEsIx7qZ00un6iIjLLGVTrtOWW7zG2U9I5bKrdTYplHUIagQZlDzLSs5cy7mg60/cVtJytSr69xl6PpUFgLDco6y6NbjFonagp3VblNR7Z/Et0Z3TwxYgu3WmDsYLcQVgLc6lYuDKTAsdfLfaNepuY+8vt8xtn4X0Ill2xUL2X1/cF7n/ANJl6b5v5lJRwkENHwfzKGXrfwMWe+HES+GuZZ5pAvoaOh7zUy9Amxd4I1LXsSgFcw0VV4LlbU3jM6kyFVMWqJ1hzBdU6CUMR1cBv0m0+GYFbTnmCth/eHZ6mbcEf2niUrGDr8+JZovtzAotVohCv8dS8FrqfcpoFg6sx7Bsx388uukWAjmdlM1fFJ4Qcp+iJ4hKCFXKUD7hVIGCiYk41K4nrKAZSg7nqJgYZwnnMvsaeqVOUdETicJfjsjEnMzqbKDrqJdioB3hMxX4lpzKVhlwRcpgtH63FVt4mjZ7JUsj0Cd7th/7gfwypyesZ8+CLzCM+w7IFjTbzG2AVivMRUZu/wAxOejf/ZVLrfadKjs7SvbM7TLJ+/j4fAr6iVNfRFZJ7grgGLuYgN3Fw+rlUYWUcUmJtSYCW5jM5alyqL6Q9OPDWpkCTtLoPFAgDOUXZUZH4iGwR8DDlLNr9zwd55ddNReVmWH9y1yFS9obnIDnrHbU5lly8ubz+4NrhaMQiEtxxKLK/Eboh6ygqZlpYavfSbqhx3jWI0VH+wGH+Fq57IuDUXG+n1KNHTMVP4c/ER34pHijnctmyaKjrDg3FcfGDXaukQpTxix4RYE7BqMXOK+olCZh5C6UAw/zMyZ0h4e2I1NyowpxG6pKgRXAxKwWsLuzNEtDgvvA7by1gTrcy75XM4z+0u0JYLKOs60ECgmJRxAllzWpcU31FLQuuk1gadpQWrjdzA+EEqbk9HGomffEAgc3KopGvD8RnYdY7TXj3JlFWo/C52aKz8ookdeky8k8pa+0o0wHGUPyvqM0TExYTTodUfEVCZdqDsLtOZd9OiUNA8an7qeCdcxGjv4OomuqomN173Eu4BiANlbJwV4sXns5t2UJziUsj/cPEZpFG0FrpK3KvfUVaHm6ROwdGG0LrEFQPioT5/WoeafUL6eY1QJ0KwRhtxMeMWB7ywRUvecwX1lCF0035J7UxHT9EkFIm4E9ZFWtnlTDZ73IMDcYGZ0QbY8mYYHIcy/6Fg4LyhlPLCpnrNQltNcI2QGpnIS/WMjzKUl4YikZtwxCDWoDeCY6UZROAoBRZ6j87lJgHSHaWMGBnaRR01B1WGaY0xAEuhxiWgM644JrtQWVR4ZVWf3HSC3dse12lMjC7hULfcutxaaJcAAoS2ymjzOIuXhY2zqMOzbPEoywYBSNBXZhg27ExIjeEgNsOu4K1tw/RF/iCFGFyOLJJwO/WQjkNovLMMAB+2LBU1Gxoe8Fn9zFYmbVShzT4f0QOUA5hAtjYM9XSFszkviFXbMFHGP+K7wMltb7gPAdvflnIxfzMdbVq+0qXODAHBL7s6zOt0wdYPUkbNfc56cd49/5iOJUSLWzN8altoSzo+prSmuJQbTugH8JNrjzGQK94SWE7ojB8GH1H3OyhRmZIOTRKtx5TKcweA3My18TNgE3k++k0DLX1iOZhelwTX2GCcgyP4gOQiHVLNLDsoeZwYYOs5xuFvTUs0oNBRPHdvQQoxz4Xwczlk15mK41Xh9rmZXH3cu0p/Qmi7zpiUjI+YFBCKciIJwTvTKE5mLt9oK8TmvtMTKqqN8LcOdt1mVSQNYNQH3cRVuBnrikkjRdRHNz3gpUYRWz34ThFqFAUj5lxe8xa/mU28XqDefExLvj6f6SxRijh1lhLd7hgAHxsa5upfnKK7MAZJ9TbdwK1ZL8/ZnSC8Qt0A/Q3Ey14T1Fa5xqPBqOJRsznLgG1/AzHTGJzcOpzLnHqIzhXc3BCfwaDQ+jnsghYOIPbE4BAsc1XWNu1RscOYY3c8Ta5jg1LBnzKGLMS2Li27TMsZjc8FGBOolJQC99pzyXleYl3opgqULgI7Bd8zWXEz4X+4U0sQ5TNofvCOaVbVdGdYHXDols/wAsW5A+yU1L4N/mHWDqX7QHzOrj6itr/rpAOPCcf1sEV+cf4IT/AIMesVpnlVa/Jmobq+l/MMFB0S6Qa+zmGVh850jf6joTUroh7JscpRkPQJQ2nR4jJTFRf9o2bQb49wzxCv2MtrUYeXmcHDZXaUCCyFMdo4oXJGgqkuUhbdyS7bjH5lDKrgeO40HfmY/WVx/E05Y3bv0Z+qwKHFS2Ix5xF0JToY8iZ4Z4zwQe0vtPCeHxaoLrO4RvDyx6gXH+Fx1mXNdCYW2RTj8tfEqh/AdTWLHmbtp6iN3+YgNLD6seZxqIlrP8xzq5lyYlacfmCgF+pdX0sLq0L4ibKToopmPLEdqjxDm7eJQY0+JUxPbELFTlcwl2p2lmWuYHIZ8xpZKgxd8ysJnr8aeVPULVv/yE9dYRyqWq5NSsOAoVzB2hSiscXL5RNUmIhS+XH1qDz08jl9x9NiNMite6dGDskzfuL/jE8s1fwn9wA00OpRnd2lX3iajNFj+PMe9b0Y1FtOHuG132ViJdpHBVE4hl+VdpvgcuYtSjO8vzGl5zAvkKJ5HMyuxg5PE3kwd41GHe3ctgXthtw6BmN62Lrv3AA1AE34gwOr2mdivcGC7x8RwRgIszKy6tw3u9UdeInpXZ8BQ7lQcifI/xmfQdPsiUfkYfDEaPVxKwvuXaU+4ePVUq9L9MbR9iDl/7lt0Y23bPZhdDBo+NVuLLuTmMYa+5EXQ7iYcSrydkS8EiKUV9xC6XMtbnrM7U+5e0GlJiDzenp7MfBn3CqMDcMZLLC8vUWa48RsamFz1j71UEymG0/UjtteuYhVLB1MA2QDYwC6VhZcJ4jsYnQcygipvQS26tu7JXAbfEaoslG34gvEIbHNnomoN9tIaeSaZHbWRPz39SWXfaMYL+iXUaGpcL/JDC7XZyjJn0fuNsD5v+oBsDpfzmZ1IHdx0zKagIFnLL7mW8QQXeu8pgr2xaw89o/MUOlQueFuJs/rUvMPKpjhf9x6Ar9xdznOZafogMZPxuDIqcDU5hjdz/2gAMAwEAAgADAAAAEFYBC5C4WivHMwDPnvDo9GetjCTRQESSdOQo2xrIJBMSfex45YSVg/ARBKwUo9f2UXTxSJMqvk119dnpa1NScG/AZAEtl2uBVCdOkkjqNpZt3JZS19MCX+hJG7Tuu1RYF/QF0Tbdt9gif/cZQVA8bgxacY7wjuSvzybebM0pA15hNfOSRtzhqskhGQEBsFFoS8N6kYAEaYRVc2KLdgOgc2NmthMK5U4kJsTghBisXVluSpWEBRjpuHyboZKpAEaZQ3tOBh5js5YecKiAMqQtFGatp11qlGlSKjLzUigFgtgGllxTP14d4OjsOh/gA6hAqrPAUls3KOqF99ESjnkO+XBFoAHLk6riX1+bhTykCAo2EoxLJqAhzCEy2+C2aK0f2/I8O8PNe0YPIVVxshc/hexy/IjuMZfE8xZ8qRBNObGHcbidbCKtKnaR7ZXw2HN2NHA/MYZHI7+3XTFOwF45XHRPHlihgiLvMFWUEdEfRVzXzqUMdMx/CkWjW0c9V45bU+0vilUYvBkH450va2NGOt/fHCfV4Cilzl0xrmnj6b7F/r9AUOSaH44EIilLwplimNyj845purXmxFAEYqi9CorjrvlHIpzr13T2ZF767jGqrz9kygv/xAAjEQEBAQADAQADAAEFAAAAAAABABEQITFBIFFhkTBxgaHx/9oACAEDAQE/ECAyPV2yA1LWHEgoHYPnBe+XaKH7DZl2S7/Bi2+mAfIFsCS0T1AwZDGNTEZ9+d/8jrCkZF+4F9mMsxRerT1wZv5NkTS+Ekyfzb4trXR/1Pn7ZmEjrDmxex4vnRelu8b+B+Tw29EAncfLPHuWrAesPMIB7JuqGM31avaPy/Bat/qP5lHWQz5ZM2WlpwvEjN2i20N2QYuv1E5GbBzqG+ph8gdkxA7HHVeS3csLOC+l0Hu6jsOBLl+rCyzlr4yJakOm2ET8ljznqzhIoC2Av1ccb2x+1nYl5Hy/TLIF6cQBDllyZZJktn7Xb234yi52uZJoX8Wx3ENfyXLLpE8bU1aMeAhiH4PdvI62O/IZVljjy5FjWzg7i6F9lvceEN2sfL7BeuS2eMIq3ZdCddUWVacI18n2c+H4BHfLZvSF0G0WweCK0msbUbG5lp0cRUMiZXrl+8m9Z4HtP7gPfD24zIITfRB4tgvz8Y2ZdPx70l4Zgm4O2xyR6hphNDi0nuGWPLbq18hwchl3Pw3kESceX0fgDxj1ZwxM+wd7t4PGSczuyRzv4Z8OHSeu8h0y8WJrAbbyyMktQe3eO9rfbhTccH3le36YR7dYevH1laX1EFs31gB7aGGQHCOiWWts3Id4ZsMcLrnZm6YvCTiHUOncChOuoH7CdWw8CskzmXZeNlHLyxYcJd/l6w29Fs7VyHOMtQpZTeFNgyfYnzDZnBvMD22FnuIY4ncl6u22PSSeaNhGWPbxEYOO4eGR2cl2I4el0bPaZPSWQw1kQ7bN03eZ3pfusWJPUgBws4rpYougb4jhOdnEJkD7M/xD/wCJ62YAZoKjRLCjgPjR8t3h037k0eQ97vGOGbeuIvI8Z4JSXV4dZlEzIej9SdnyIacjgWO1T/zh+pJ+bd2wYix6gJ/exqEBZ+4a5+evQv6SH2B5K66SPRvEJ0cPv5LgbWcmQfYId3RvJirS7h7/AC65b+7/ABeo2ueE5PN6t0f7cZPv4PCrkxauALzDqL1Ah+kRDg7sOEziflnPr3/on//EACMRAAMAAwADAQACAwEAAAAAAAABERAhMSBBUWEwcZGhsfD/2gAIAQIBAT8Qa4MtwPa/x/1Dn/gWT+if9Y94gpmg1q0+zxht0j6SA08TRZLDZyLT2N/vCV5CH9JuwajQ43yR3dNrXqey6kN8Y2tRUYIkzcg8hiHhTfQWmr4N6tj/ACJahetGhtL6HvhaxEEkawOpxHJrCki8GGEkx45H+oO6+2b5/wCPYhypPa+idXZ6S+DpY3wsbDEhIVl4NOp7Qkm3oX0F9MCU2Lu4GGT9GNxsPwfk9AoZEClLpLdeL4cQlD0ENaSmzFMQ+SNifpo8DSEQXf8AoegqR8GfqIE3yV7J+n7kfSvRS6so7GkJZ6I2QuH3h4ZNmvgvZHLOkE5otwr9kwPyxJ4xd4u8bPo76GPRezL4bxkIJN8FkP4KmJQhz1lwL4hRxi9BUEUNIuLDZrCYK1/susJHtguDKUWxZJlGILO+GhqkiiiaM0JsZI6x1eo3b0JvZP0a+mNWxOzge2J/SGrATKNwcVueSXocdNiOYmgq2ukFmCWJXCxM2jf0SQoo29+A8Y1m2hwRy4aHPFewJing2Y3qE2N6KXFrGPXeF4C7sRoiA4Y4AX0avCjToxBM/SwZesPgkZZa3mOEe5HLBMQ1kuG2JNbKMbuE1EPcnl+CdwnSRIUO1IIiwtwSweDWhITDEIIYynEl4Jzh9B/A1TMnA0TcOtMaaxrZO4cx0SFihREBTjHgXXFCxnDITYhIxOtopm2zjwybERQ714T0JoKaI0OYjRBtkkZNVDWSSEGLfiFrGIsNFdOY9CaC6FSCnILiOWC4WYpRRkYlBj7xS4ajxNY5EPRGfSEqCPYpuFiiwX6fEsbkg3hdjWfoSkH1CnaxPQHyMj7KRRksC9ToljYNhuN5qok1h4mEEjVEAkPwyAoS4HTuH2iXXR1aZcNDRBhEIUo2J4w1FFoZuNkq0N6z0PEkENQjw/Mt6E7ot0UUsv8Ah0YPexC547T8E6dQgJvs/QavYv0O6v8AiEop0EvJ8jxfCKpQf8P/xAAmEAEAAgICAgICAgMBAAAAAAABABEhMUFRYXGBoZGxwfAQ0eHx/9oACAEBAAE/EEfS8ojlriZlglNWEUjz4sKV+YRMIADLM8SFOt8C487GszH2J6sQTh+LzqVtYrjJ8L+4vOErbans5+pQlBYRZXzGQyDuBOX3UziUNjhiX6q615+oi0uC08wDI3ELOGJYlDdk7IOVmHCdS5NlxDLVG4aIEWz7BpzEkmELccJyVEAEoBJtH1YeUZ8l/wCocAGuhC7qYNLNQ1cG0jRKF+SINvBX4gYe8/lMwkqexi6uoIlPBR8y4HIY2cyo+4y815YflYUPzv1LCwAPwwq/RWwA+JSRaxoclYmq+6VIyR2onI4oIUkrfM8sDeggOSj+TgpqcbI5mVbKc7TEEUPUccA3Lo7iflr71Kpb5wp4QlCpWqwGso2gB1Br5iNgQazfxHqUFKhdcVLhynYfuILgfKLz+UjSydD1NhjlkbVGV4zIGMuDdb/ETOT8xT+I7f8AVRv8/wC47Tr+EzWYBvDGmEywHxmW8FuircJu5fsE68kb0r9xHrP1BFH/AKYJT/UmNf6s/uxiD7lv3Af0uI9HRfuJdyBMf5PuHYlm3H+5ewZFvnGJQeIOKI3cAcMyIA8lnMrZPZA3pcCM9KUI/wBSo69lCcDwxKnEao+/ghbYkBeVniwLKKkzCMgJAkSDh1llIZKIPOwwy/KDg1caGVlorNIBWL8YI+WCLyucmI9DEuo/Uz6HZAVc3+ZdHky88yKuVqAF3bIrPWF8zwFfvE/scxIp0/UzHuv3Hu6n1BsygbM8TvF8B+MscpZboLPFErLVfgR0XCxH9XEMBzb+Zgbgr8QfItA8En1Lu7Igi2B9Rg9gszv/AILDF7/ufV/qXf7fucXUr6ndl+LmAM1BVPLIwtmYGnqmUaK3CH4Y9xiHI824mU8wLuhX8QORQvY5qXQmBpTDwb2IGinA6lwamdqoPmYLTsdbIeM4lRlWkxZXjthNc/pKfxPqLdTluNY2KVRcn3EsRAIX8E0UC6jLmkxqlKk16uHvDZ17JM7g/ml39jEp/UZiBiD9R9R/lNYIPqszGegHCZYh3BNByEf3oNBQ1LN3+UT9HApm/wDSWcXU9kbmD0U+op4V/UwLtWUv19UsPg+ojt6a9EbdH9cU7aGO0NUmmaWmzzmYOyjiVeYHy4J0nkMSMIttzMlE652BxO6wKlUbH6j5IwxY68za0dBv4SONaYaB+sYA1vFxy42ZDS0fMJE3Jghs9R8I9Aqy4XNqwqHJTx5JYK/nMRsNDpD2DUob69kSc8QoHpTzuW8pQgvC3j6img+sBeqqpX5cLy0wZuIShkUZCi63AsHHZTVATMGEGVPDdR/P0NWt44hfB8sNZy59kB5rhGz6Bw9RgHX9wmtp4QTCryXnFeXiH4OnMCOhfEc3DtVlFjMUt5rLHEW+p70vqfPL6n4dj4S1/Es7uR1tlr4lRumA7Av4g9jR9xqHGb5rEE+L7RiXwdwpWyrqfl1/hElPcsM8Tj9Dk+HpjWMNQHMbVDLe7X98xaxMQ+HMoDIXKXh0a7lkAA6t3xiGi2S6YL9F8refEzG4NG6x+InVmXoufNS6hRlNIGwUa7m9SbAhoeYQN1zQs8Zx5WDPAvO0ywl7nNPbwNltTDEqZBS3SqhaUCuay1ekCo9vmDqFa9hAW4mqjdfMpYA9ITk034Ybg61484z9Srk2zZbTqpg4CmdgQorqnuMXRhX2Z9ERvpVhuTgPuDgHWvqXYIK43/MEyIbtR+8bHs1fuC27x+8/2TCNCitTmv3DR1bECrr/ACmKwWQrMv1MYmmH2kYyZBzBa8EeLjz+Fq6Tv4jWUS04PcVCxEDpWIroWJlhuFF6Y+4h9DFz7mpuXH4jyl16GJUIygesFMqYs3PqFWrjtkQMMmzqZLw42cOyFxlTinJQM7lZ4OFKaN0b86hfRuoXbU5O+Rizo01kUEwDoGA6g6gnXKGLbDxMf0FG1RlzB/EZxEXhDbV7THEd2Yqqcm8D7ju1vRULdiF/EvKOoXvlaUMJREEu1ZE35lykomqKZ/6Iwmems+xl8wOyz7JjLu/LDACr2tXtFQrSJrN7E3Pl2EmRuHe6WC64C8/InFrix5D9QQIZcj8S9tu72/cUVUwQPq4YFZwqfMQVWR/WZQSxGVztkbyMS4okqq8wJRbIbhH1GhrqIrVrGl9x1OAcldTkM3xLBZydkMCT0lx5rjRGkfvI+G4NRVxAVnQq4ZIEXkhB5hCqcNx2kLpGeaUlYiFMuVRKdkXyDXccXdtk0gxpY7W2BY+G/wCINs5XA4c+U45jmFGmRsAw0HGJraoxObyt9v4l6AFQBsF6MFsTaBKWEDyzP2McD4K/KCkIYTILr9kplIdr6gbdUTS+BGmlgzVQlIUqVjMm8kAac8CwBLpZtHVMijr+GVXlYXDQUqpmEjCVscOfEMUlNsRwvMsVGlBtzb+xGLCaXHYtikxOSJW9l0iR2A5nhH7eJhOh6ZA+4UVYGKPiO6h263bfm6gT8zgfuIwb6CaTvQRJuvH/AFPtlcsg3zMeZatZX+4GtHefmeYDJ1Kc5Yc/MgD9BC4C+EAs7uksyouYaxDK9TCPQxWKb/EsTdVaQFxr0n9Tcl73dojuMG2PmPlgYnzZDGpi1dcYi/8AMH7iFjXjD/MOWFo36tdD3KYy1C5wtk9FXD/lqwFr7E7damNSVUUKANnubBy4S+rd+ovKSlPRS1gs+Zgs6pbl+YrZos8ref1KTRBdNzEorsWT1M2Z5NFer8eIWQU2Bj7IyAQUbIe5X8pAuqua7jMANJtg5xAF1FkF9eYjGSwZw8NkqLPkCrvsUxmYTaVaHd4yJDDBNoHzlXuBUyM2zjKn1BAIhazcjJ6MMq/QoZa27E6ZvFTqqrCJBbDX7kOITfeNlPl/8QDcsgkfxAYB9oTgHwnGnw/xwHBoNeEA8DqHVMKHZ5cEuDYcf7GHdfqf0yhEavaVFCoWVbE7jtRA5T+O/QTQf2ukh1Cex/bAKQeBE8ia/wCRG7T90buz9y4tAbWPJjX7z9jR5l+IboO7eYEUtAQbEeyAKBjhxNtVZrXiGqpSWOzEKvM44/LVHE2XlGQL4iiq1TAPLDvh+pn6hxQMeA7lhULyM+7YgIOLivXuBg0Xi+rgZkpc7HiLVeHB7MwVuywP5vEDDbsxCS65dZw+PcpsFmrK/EImR4Pq4M3Wtew/BIA5iDYB46m+gLAHrJHZy13FmsfIikZAZjvBwPUx2ZwGWI7su9XnlqE2bL1hJqcWqS9jx5gLVnygOm92xH9og00DS1BfyRQEuL2lJ/Mh/GROn4IOG6EFBN92qId/Sf4lOi9KdLf15l38Il2xDUuvv/iIoCW/mN9D+5YKFe15xMHK19ummnl4nlsGAuGhbjteCX/Hgwe15aj0vYxY1oioMzRkWea+4rrrQ43iviFNhYdV34iHlHa+YR2Ghyd5aiNbAVsJpe/Ux4ZXIdVzMZB1CutpvPMAc0g7FLKDJ531iseIXI2w2Pb5l4tfW/5Ju8FMLCn3Ph+dxfoWwo6oPfMtFtheYOrtfiYTG74D5fTuYuANtm7eMagbVFwFOOMQLoi0NHrDXmXOomHtZ1GfDXcOuR9SyEdFAxi3PxMLgVfVP6ZgaSu0NKOn6ZpiWUK5E7GWDEagvZCFztpUaDAgmEA4Ca9xkARt+q4dS9Eh1hS4fzLo7MfMay2Epds+Ef3MNQvimypPlbrc1QDLRPCZmSAF2C/I5lWBdakdXeGfEsVhrlw5MAfMurbKuAcrnqKXAadVcY3GNRDsHriXcF4C7iedQJwaOYhXy1oD42XM4d3P5KXKEVaxGBbrj3C6tuhhQ1+fHUqFtUVEeFXmvuLP+2tuHsREC3oUh670Cr6loocLavzKNQ7gWobVZGrvMo4mKAVSt2HlmKVtFihaijwRox6GgckBGWeeP7pWceCJ6OWqg9OEHgcNBbwoHHuNWtF6MgwRvFmyg3b0fbUS96XK08EAa5T3u4jr/C7sW/EOAPB7l5gKJVEuY5lHJ/fn7gy5cWDiKmLwYeiJA0ejmzF1U9S33xBO1mv8rxFYWbLmu1WAteB6ef8AsZhMVZsecwxFL1SUgXV3n5xNyEhZZovQQovC6Fhafg5jr6NWFulr2ULs4msAAAqNuG3OCkiMBbli6yWUOqCondy6qLl6MviXJYsriHxdfzN5hetWU0D1Hl5Rq7BRYNh1qEQE+N24a4DBCZ0BcD2G3jIIqwHKAHtHFB3BhC5tUzhjOguPYoyvo0uBk5qKHBWl9Z/M9TL/AOzDp+F+WpmffKA+KbhzUyN42+1X8zL7iUF3rTMQuloC7jKvUuLgZNPYhPoZjYfivER3Bt94mtw8Q42ALJqwZSEhAImEaQbp7M/xLhj/ADpHHfAPZC4dKltpg1bUcjjLkb4ih554rlhusUbZlyNuD+HMtxMAdmTPJaguw4sDzRWFzWOqCtOR51FHnDuN0bydS/atri/PHxF64DCA4xFLQ83P3B64lqopLH+AN3UONvSlYVjhsRJQaNX7fxGopnIXgTkHAMq12GY0qj5c9wBb13DkgPvcuckZCnCaIM4wVBHmjMGpDRov5hq8clbv7lqeQWx4vmLy6wf8cSmGi/F8QYGG8l1f1KJRcw88OhtylbO+Zk3iFp/hKrCUysly4MFl4iXBG8K+o1HUF53e8YuCcTm+oHstqugOYTYVeY+W9DoI8A6AQ6EF/LH1pugrealXTVrbJ3dn8RDsJyzA3+QjuX/AvFajbYvUJVMAjLDFyRNQGABPSxMsFtL7Rgc3Jvi2vqB2uv8ANLlx5iE1vDntQ/c45sCc4pLuW8HZwfTEa44fyTDdh5oX28EoDQXgKnv/AA3B7cvRzMBBaJlv/CBjzB7Uv/FkHEGOg6+5kYA7GPmd1iZRxZyZSfDpj4kceE5XI8MfHRTnEIMgp2QAZjZ0HqKHIm3lF2g9oT/gFKcw1lokLwhuWQbUG3sYMsnIFJuYK6S0GKdQxTtLn5Rj9Ro/p9x71AyXKdx7ru3GYIL8z+0P8nMTfwaJdxY4di7Mw6X7g9CA4f5OOSUk8f4ucS5yRJVNZ+UxkCaw+ownE6Br4EyN4eNSq4XlvV1wOSeNgPmGRYIJ3GosirIcsRxIEYRAeqGDeHqivqDGVOJm1eWOMqab3Ialb+tRSW8wbgpuJ9pzIJ9Sj3/hwPHCjn+JVsbMYYI07xznDPBta/kynUr2FpFSzMAQV5zLqItWEMstxEJkSLWpahLly1hBg1lN6PwCy1oIVh9IsE5kFPliyevL/EqSX2prKisQ92zSj6khHR/VCdS4COoIEcYQWu2DwLfEZA5MJgwRSXiBoovWiXZMJZdCFf3RV5kVsnMG2VtWhjr4/inEqInQFegg9sau6juU+B/MzFDnA/Mc+w81EogKIaaOUP8AmBhBlQU3rHc3Q8pdPzFhCd8+V/B25jOc3McwqxlcYB7m+QjnGtEHaXpM3M24QLURqXKlz1FzNs6mnMcwddsCtsCFv7vMPAzC16P0ylpY7Cn6mLxRyKsPFYjW/UEVs5mNPcSneJuJ+UGwswN0d8D87lvWfAqPQ9gg6tu4ElSCO7ZSi2CYGPfe1C+jb8ShBUArfdL+ag4VdXs8uVXTHCji1yzkvlVgKj5CP4LhlUdG/wBSoLsp0+QSIBg3+KFpMVwdoeWDvBFQLC6jZS0z5AGCZDY1QRaBWYOt8P4QDMNKFgA8QR3H/RAN5huosHiEsermNifIliYipCzMbll+EuXBlkrM6TrsD8w06vHIP7h0pMm6HwalqbXQLOCzFxsETjZgD6GN7H1AMLual0wHVpTjShUNmZH3M7k8D8FmX81hNnlAfqVLYF38v5IN8YcSMmFR1Ie0FVdUmr/kB9MQwCUoHmqr5GVPO1w2Q14CoRbV2iIUA8eijdX4l5wR7qj1wNAS5dFpAeiplMOSwPbojqalUw9DUtpuSCv4nCEqJlFse4F7KQF+XE8WqNtWFwOUnyJVY4sY8Sm0ZGKaOZ4eR9SaWXT7TOPMM6gKjpjLzLD4nCzU4Xy9xO8WaJYwKHIjPr/PxK/w6YLJ3lZsIzU5Ac5rcduV0KC0FgX1EvtV2o5rTxVVGNN9hDKQcyietygBT8yzmVCOqCw8u35iN8hdINjC3Gj2IE5kiL0PKbfMa5HBYPnNR6wGOfKM0xIXyIwZRwgcro5q8RpaJYKn22uKBQXAGshCBq5wazOdOaGvDGXi3UtKXW/LHmCTUrxeFY6hihLQtPLCrnqujiwofcTakDYolDJ75/mX31dCMMJcPxodsMW38THNEz/SPKo+2O2zHMV1LuWajUcwAQ5ssGYnEu5cWC85SL1C2pkly6jlgrwi1GGUZxVX2k2TaqPblFHLHy9ffiZkxmLrbeYVRr9QSvzOyOtiYvAQw46qedm/UK6ggG4aYw2PDKUbhrYxC0HyzaCtM+rxFS1YHIfiIXBznK4t8y95krtzKGwNPcekJaLZkMovB8y6+SVD6pOMXYDBxMoq9/6EAjGgwjkTES/yhHfWNtdXCIYLxm/zC22LPnEteCOPJPqDM5pUvDYF+/8AAscDcXaY/mn+Bg4wYU4l9xFc7hX7YLmFqym1jMM7mmBDflcAVFFAvZatJk5ngJHwlM4ynMDDt8M3sSjq0+fwmRpsiu0peJ93U5PNcXzNy70pdevlJo6L1MVI0C/eCYRFzYD4iBYgOBVtopPDK6FM4UVbowb0QHaLB1hGAc8jnwu+YjOh55gHC1K9m+Qz6HbV9ES7jmr6OofFCbtMKSY5IZcvoN/2lopwagZi6MIjbUXnxY8q5ZLhNvEcNRFPQgnCDXmMKGObmCC/RGU0D+ZXupnwBtRbXCGw8S1FLN3gRX4xUfOXdINJXEsTr/jh/MWn2PcGbx+Xaq93DjosXoa4/wDIYV1q7NvgIZwOGfllTmvm7+6gOWOKiOHfu7/kw1EHYsv9PmE8aYOBHzpxKCYuuJyavBGvh0xEV5rsTc8WM11OteZZDAGIRY9FUIIHVRuSLOIX+gczihKdryvliu5uJ2TC+iAwaHLuUDOP1Ra1Lgypp/weE2CBoQ14eZ+iUzHMXH4Y2cyvV6gxTLeo0hk7VFF5MNNylQ8oNHOLC5cusxgF2WntiA4gfqM+Y9A2fImPhEwmFsH0OSa9mHnMsiAJcWt7zEHggmogofRl+CFqJbh7R3DgMiHTxIDQtdOob+QDqF+DDsPWKz7iRB2n6yIYzvdzKOu4pU1+90K/M2RcAH+vU5YpmK9q5ZnxBjjgo+Ec6MENkofJiL/ksZj6l4dSAOdQ/ITAIIzzFz8k3rFMIUx3CJalT0idcKmXvxNRCo72qczqaQT/AFAC5B8I6eDZgRM2MC5KNLOTdjS0X7l1JULqU7LDLEb7jFLtedqD9RAON0fmbZXGukXgdRsNOAjKzaoEPwZzE8hmZAI8uUJ+52whXx3OwnxFo+ZmHXv1HrII+CJdYu8akuCgeUI66n4gerit+ItEu4WNkpZavQizXgmbf+ZgDQi3nUbF2mI7IphmOuIDjcMNT1HM0KDi6BjmC0e0yJLXMqJj2UQWOVmEB4eycMmGEVdTmXaa9CU0kOvKOWN8AzYAlwHkVZ0KexBt7EhFyviYl06hkrBLCoH3LDYPUMUZgYio/MATYitRDgPEUzC07+rY3/2GBWCt8N5v5lTtM8/4DGTFS6SOl8IWoJef8QFcMPxGKYfqcDHmS9X+UOwpA8xIDdgjXPJKQpeIegR/G5m+FwTRyPuHgaczU32l/pfMDieRjmpyh56yQEYYVR+YFgR0HHxPHhEmXmFzjPd78RlpYRbjhQtudpmpOW4BNpp5jAvaWH9xBG5F8izDkxHta15jCDdwuuAhas2vpqEz8V2S6W8xr/Da4XQSy6ygjO4lwRiKsPMFuArW08JUwDyRGmZ/EtZPlSMK7lSrhRwpdszVJ4Dkji4nmlI4jhg6T4hUlyoYhxNJ5J2ccwR7Lgi/s/BODYkBETh2v/UwnCb0xCMPwV8y9Xb2ap7uHxWoGofYOHR/ED7nd4YJpuFq/n/cUMUZwZOMbl5WDcZKEnmcRIAD8JQGiDTo9wAqYs/A/TOUQ/OYXTedxmliuIULsC/MFsThhgiWnRBKp8ZQ3cXpF6CYOCIckZ4SWkNS/wC4Q0wxgOLAGNxLgghStuEPIR/eTGOMRoKhM4lW6hg2XizlgnSb4RUN+JjA5xCrA4YoVrbxvrGZ60JAbqvEKsDNDSHMZ8zKgimN/THWddefll0hZua5trMDIgQ2q7JgLxWgHsiMCCqMYealDZ8/98xVUWQV+1a+Ibl6Kr7tllpl9M8/maFQN5Fq9IkpHeke9rsM8V34h3oznz5/E1aYpjA8MPFJe5myWXmVi4OVo4H3LUA6Fyy8rPDT35iWDtK2roqIqJa+YWYxfwj8B8y/UY6NOePieYV8IS86SjEXIaKgqWf4Iac/6xAnRKGNKiLWXeNTv8pblExX9Gszhfms0rzqU569dHcHaSZYQ3oV6g+qJnSd/mV52Q6eA3DAy9v8GPLDVnQiEosDv6gKQlaDFteYFDJsxdG3iI1Ny4BPsZq5/UEsGh0Km+Y0P4jiURKUvPtr1MBNFcI9lpQe3if4uJTwAPA5h/HK2qXYRi0Z7XMpR9zEx0q+XH+5qUMOmouCCx7JhZllbncyJxAInKUL6jWoGmGFqIiIvBOxvB/Uz70Is4lk+lfqBzMJkDGF8Ew7lKqnKGuQhYI8Stbx/OVuiAxqDAbRV5fSk/4EyNx5OGoUVDvG99zou8XyMSv8E/fsMe5VgG+hyuJaGXVkrPDVRBdaapriiB+S3RXiMRCryHxcCCD3Vd3BdcA8swiIjBgq2mMxXQ5Ro92wtkquGW/cqYv59xwcXfKN6JBzsDTrR0733LZxVKM/cYujzqqaGnMaD7zOQrf2aJX2nHhqrhT0AjQGxI6GD0ZOIHGSnxK3vEqu8YM8iPglbOcozJAKLpaj+U/SG6tQnwFLE4E/iZ+FyQKpZKh0g4CV13ML2RajlxuGiy/kjNsL9oEoAWV5qaKAc7lafi5vpYCy7qIv7RNSq63ASgYpOUK5Bq+4OQq0DxAm6VWpX1GMT49+YGghtSuQcpiMTKMlvAOSMr6Zct9NMeeioBnl9y4xl0CfcqvVxTExYapQYeggJbcuDgODwRgyn+Ihq28QtR+8PqjfaIzsY3wRKoJhxW4uZ1AwtGvzqfF+vKpx1FBy9odR8wuLhvv8Mv5In22a6YDZEU8q6lj2ExtD84rMJ4AlV+2SomGO+JlbDy7VG9BIR8DROpodRgMO4y50QJbWXoektDtLcEy8WGIW5j80hlMllmkvJ6ibUZ6MoWuyr1dHiMd74pajAexiG1Kssiud7maB0cNHa9ylTZrcDuMhW4D/AHogUME4jpdw1OCvq7+Yx8dfmEYN5cVxRWhQRVPeIPxg6H6hqk5P9olCHFDxFt1CnnEsCyyu2G3pasX48S3wGDTwoHJMEhd1Mv8AE1dWZvG4C2C5NDyviC0QN+xVfUbnNvOHl5fUfLrmGGHh5GN6Azbz14g4IB5/UQ3Yy9suxyOAo9k3t7jqjZxA0nAqfUWHu/cxBBS1dS3NoHnIZUzMpbAmb7iidkLY9ERvohCGfRZY2hPmoq2XPReFPcukUl67lPuyZydjVT5lt51e19+ZQWF4yYdvuOe5R1fNShLJhXGfBDjqM9kcrKFN/EzrgOuPiUnJ1jkvuMPleR9TDVMSs+RXlPEElWvBfc5qSVTpN/RF70jc1aPeLjYsDPy7mfxcHk2+iVb6KtZpjW4X8ulfcvcUhlYY6H3E3YVQ7A48lRaEMzNnGKj7BdPh4dRICULJ4wguPWqQOQ/7LFCLdlteSVpuz+/8aj7XEzRDB3Ocf4n7ioeDfywLS1I8PXuoIOJzEjxCBbohX1QmIU9SJrNrcZl7WA0acRjxivWxMVAkdgrMHmDjyoFPllz4AzFrMW1u8LHpuUC5ZdWCbbymiz3G9Gx3V7iVPk8/Ea18CNk2N+oMiNNYGNX3UF1zRdq2r+oqv5NQ5bkbCD8x4gg3GQY1FUXPBjBUdar/AOwUpFg7fLCq4gVlmNxFbosoieMTsTDT5I9EN2WmzOUY6Y2+7Xp7mXaqHa6fEoNVpcG+IRmAPJsDjzMAx38q+P8Acrt3k6Z8Zn4BSUL5Emavf9wF4tL0HNsVJ5Qbu2OwdCOD1KWdQalt91Kqjx/vDFemKH42PuUrUzU+kT+W4fgjBXU/sxxw4WE8QODoih1MjhmG3o1BGRkUvA87PuXWE6BsHIwyhOgUqn11FDtR8H+4O0XW5a5/ELRdDj3KjlqwlPyEF2/5KWOW5Zil4dwxVnFy5Q004maLuY6WPiH81ltW1cROfC4w1Pom6W3tT35IbKsXr/xCDoRwvkOooxq1fDJKGHANjp5G+mF2RYdBx5xLCQQWB778TOQfzOvwIkwcpRpHB9Q6TB+l4CFYBXqq5Mu/RM/k4C8MvHVy99a3EkS5UPtYYjRHTH8hDleWO34P7l9NDi68WolUjqxNHJxkm/5zMs4RzEAyEfSECi4IpUvrDpOiO1WYECku8DlzG+0NWR3O4FC9H/kfBhCfeJWYSCTY7rDMiis/ycTTjGBwSkLHMywHi0dtpfURrk5jQs35g7tLEltYfsgje8QAMmjh9yk2S8FTbs/qZweYxDKkOJogQwAds/3UU0t8DXImmXTSqOl+I9Z0vQHD4D8x8epNW2/Ur5sctHAM7hcgN8C1obZnIKuwUqvsjUhulYVD5QZV2j+TMvFw8XiJ9pHee7lvhk2aNRxRbwlH8wYue/8AiIuYBZ/MWQ5dkQGxxcQbl5lYykTAnIXLvjTiVeuEFU/HB/NzLLyIFj4nitmH6h3QNPL4mFBH4Q+AH8wSdRXWPe/3EGP/AJTm42KkyVljcSjZ4cyqXCrIQOMRcuoA7rsIbg/Rfll/QlL2V4igGRjVv7UaqA0NQSxCt4hI+AY7kl2KHt8EBGjjJGlqCyo7HY7JXIjHl4mDIWmrfAGYRlLrK6yxioGefru+5WoFxYcfqXJBdWSz6lIUlLg47uUN3T1O2A721DA4V9pRXmK86zFZOIWB4+jGNRzEU1X5QvUFMa7DbALFGCFZaQdk3rhHo8IYLmAWrEdy9xfpBEqwTBPmoRVGKc+7cVPIUB9niKlwyQPUwfBFHZqn4gwl5PBgVW/Oc/MudYOKX/kBxHauNzMsUFuTjzKHAVA0d9QS7ls825h+q/ZnEqelW5Z6hCswpvG0FlzBY45lPww4DsbPMDmNfcV1Az5cEReDAYayxetcFhquYH1HcBlvtkNdFxsScqK+AP8AkQceLWPHV9ww1czHcsqT8ev9zpsh5gzYjGYWprWly80Uvn5RjwO0iVepSciKvwSu2jED5phfgKyfqLjn5fqTUebWHp5oyRhAsJV/F1OQyof2wVxoHzAlBuHVep4LYhr3cA3nOz6VG6RdX2jD6RgtMHLywMIuT2zp3XAeY3SK9Ml/MJnFz0tn7iY8Ia0Ty0Vuzq6h9bd4+zmZ26JOL4L1mBOKxtcu2BJjCVgmeLilE2uBo41LZ2ZZhWcbPO4ELUcFTqWMD3FLbZa8RH+nKr8XM7ILmLDEEmfHfUx8nvdHxCTCZ0HtgmhPM+BrEfQEOXTT/s6YaL9QN8gXOojC9w4TbMqQhIRJh5lD4WXqXdbrnjzF3uyF0NH4l+MWpSZgQljyRhB+UMvXti4M2rfzC2A4pdn/AJpQ6mCGtpZ6E2MWOVvxqXojxg/dzNNcVk/GMQFe7Acj2TG0QozfLL+pwHlga7oiTU22pZr3Dog8vXj5J4ZJkmRzEPD+IVZ/xmyi7TKNBF7hAdhBUIZYroUjitHUVb/2GM0eIS9TuLQyiYZj97lVV0vwMcG7gzwqByNxruc2ymeC21L8HuAsaOPcpCgGYDbC7xFDbCuCHV6iLB52j8ktu9vmMY0jrx0VG4gr1c6y/wApl+c/9m5R3y3EL0XBH/gMFD2CJKZhmc51W+o/r9JeuZgvF1fcVr2VUnk2fDLWiulTxaH8Qg8qUfk1NpZK6uUOyuXJaIM2rB5X+IKtdvOHcwIQj0yq4+SDcBR54hNodQxp6WU5Jh19xDOTmBwg22LpkpE/EhP+2cRBfzZgUDbl1CqExxdzad1wOVLMl+FVss0Fmci/4oTyGencrhbOVnJ+CBPwPDEdjANPwnLywislPws4AAoPKAfIOULW7eoSV6N7uaaFothwA/mAZCs+JarUP24xGkSATqDYim8uz9TastllTe6wFoK3TOlK8tHO4PYUBLZoShuGsosMuh5fBCTmVef6JVql/UYCxLGlqXvM9RpN6mvZB5fnyt+5UmKl9A/mfAeWjs89Qe47P09JK0O+YNaSdzBPlJyg8R34RRgHCqKVBUDhBjB05hqKTyyp8kwnL2QHyuJp4H5RbL6Wj3GCouwS69BwGJReRwR4Bbd7H7lsgDah4Q1zXFflguoAuy9JBkqHLmISFWXHpbLLOvMX6IVe2XhxKo3EvltqUtw34RjpfHqAF1sZE4Lxp1DkD2y7gdVG9jdcoSjrG1gisSW7igq5ZGdxLBsQqgPiAm1cWIYE3W5gRIGxevDtlWN1iAfzGBQZvp1EY1As4HcNZ8EZmWHI8wldic5C/wBkDx2H5VylPXPE6bTnxP6tJnh8yia87jx38QCyS4wLWYJXqncCyQ4xHqoTm57o/UqpdeM8zimK4vRtfBAnmw/MBl+U9QUNN/U3D3VxgvqNTj/8fUN0AD5H6hUf+cZyhdKmcviZLkvym2waQH7R/p0/grZKy2L7QbPxGmrlZzDUuk024xb1DarU/bEIQ6Tr2x2ej8y9ykq48Ie7jTNjlXNxwlly/wCJqY8zFbCkqok+Y4U3CUtirD8y/fN95GnRQQOEdQL4p1MBhIOgHAEbRhXqV1bWRdB5chPiNPDDkOx/uEgw+iAXCpOF3LtwB9g+o6cgD5GI/RhJi3n/AEfiOq2PkU+xnvY0V1Sn9wnyV9QozMUdn9CTIfA0XCo/rsh30g+2E0x/QA/ctZHCH5P5SzAbT5DmLliO4oI/xCJwuLY9wimDAejUIQBTyzT4qHRpY2z4PECKuarocEwkG/zMtCn/AJRX0PHEad8NMXtg09nnzKtn76I/zD4aumYcPTiXM0/U9R2EtTtiorbOSZEcrhXtaB/3K4rDp3AGcri19kLqWriUHjc7aHKZzStPcogToHqBaeaASswVCo3ibDbaCo9tMxKthqGY+etIoMsHByswug5zf+QapPj4H5JYeoPu5zO1nxHXrDKPu9mZcmLln5NfMd5JVqDvGGoskWwoMc5+BLjE148qYCEDKTwnlcv69yqxbOpwdVcFvo/EvurWbeVDri5exwApA6mFdHItv3L1WHUV9lc/pmVQPCNkWpgO3FTRrQz59yqUGTrwiMAAx/rEFDhyyo4IkCrphyfj/wC4RVFj5QDtOI6LFulmUgXuUCTlhCZQeCYBTwizDZR6fL4JRsO6vtm/jUUGwxBmULpNqi2xT6gnMpUbYHpmMacZeZVqEaxksQ4GFPoV3zHLeT61LgRpVZvy1CygC2YpQ9kwPwdt+fEYCj/4QihWrS8I4mWw2we73+FTM1a+cEsl9KeVt6zMpbwn9TH/ACgmbQf0lmLy7LB8U+07EEFZ5yV7jUNMh+CLelr/AFBwQsKr2/coGYo6Rn6EYTStVAJwEVDt4Qqd1QApu3dZLjKaObSxxldbxHk0mmCG70FPZLgB0yg6qXVD3jVxzuWr1UfdBcwoXAwrxKLgODeDuFslZUG/WRYNq54uPNL4c1At55g39aYVfUoxCcPouP3DpjKxR3P7qUHNTpm/aV+oL0xX7iA5N/ibAD+oQdNMH/J3H3nwh6H7Y8A5HSZIvsOSKEtWvPMeo146Mn7RVWoo+IE0+Y58g5+TA2he/wCKUsh1j+IKAvIFqNgWayQEkHH80NaV8coSNCiw2rqYzUDby9RgKekBltTAu/xzL/yo0L4om/Jutlb3MaAHtv6Izewq5gvgEH8CMqCnCGJxvcsVJ6ETIiCoZkimIbTpmjubfHj3uUSxaub+J0zgqJnK7jT2GIdI3lmFY6cwhdhkIasjh5i2v80rCPUMLkxHs/bHQKHON1iQ12Qavz1GAtF33LvIZzx6j5reYr/B4iReqkVjPz3mZNNyr4PEdlgNFsfiefxXUCttf5Rb7HjExWdB2i1+Sej4s8J/Eo1IjkD6dEVXyh+wZimoaMj8MHaVaSA+GCu0N4VDpOyBP2IPzA4Y7f5Fx3aHRifAmOpRRCBgsvNa+rcHxEEgcwo/j4lSAOXiHvNHfC/ipoRXTxiMR6NmXxlvBqE3I6wr5iQNDwfgIgKtqm+4XyB1PXLYZVZdEWpo3UFS2Ozp5Ix7nf5Ds/UBkxLLT76ldLEag9P+BU6FdVKw4bMPqf/Z
