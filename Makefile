.PHONY: build
build:
	docker build . -t ghcr.io/fromsi/tg_reaction:latest

docker run --rm -e TG_REACTION_TOKEN="6213811905:AAHZnBTM4lmthRH--aWZlPotW8TT5uyEVjs" -e TG_REACTION_REGEX="(?i)(вечер|утро)" ghcr.io/fromsi/tg_reaction:latest