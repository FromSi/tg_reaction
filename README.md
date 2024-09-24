# TG_REACTION
Sending reaction by regex-json via Telegram.

* The bot must have permissions to use groups
* The bot in the group must be an administrator

Running through docker:
```bash
docker pull ghcr.io/fromsi/tg_reaction:latest

docker run --rm \
    -e TG_REACTION_TOKEN="secret" \
    -e TG_REACTION_DATA="dVNRTsJAEP3nFE2/IPEE/hjPYT2A3/zpmgDVCFIkGk2EaIxGv4yxhRYLlHKF2StwAT2Cb2ZrwbSE0Nmd2Zl5+97sScWy7OreUa3qOPvKceq1Ko0o0hf4N9R6Sb7STZqTTyHNKNbniOmuPqcUjgWvFBapBW9Tt3Ce82JFSxyYIdkEv5A/krDPDetoeFyzd60D++fp7tI+3ClDg74RJUq7kpiiJkws0FoKJUMuZ/ri69IUNpIApZyAk6n0p0A3dJem+QbfMWJo4MvVkgKo10d7B+b9uRwbqgBbmEHSfXReZEVRbokfMCLQKrmv9yGle96W0khKdRvl50aGBmh2cUlvY0tzOPpKeA3FORFmfH1lTqYiAFMBl8KdXdDlWpl3YnjKc039AtL7QJBee8Y0xAgzq4fed9zfqls+CNA/ERpi8DNmwrk9dKGFEiFifZbFGAtNIUdMM0VjKOZy7joPjOtmAeOg8ydUjmoD8+1bOcYA9eaoB2qY7yUTwG03RhtfDMyCfMuc1h42Me7EkYQjkWV4Za0LuG4u/wO6fxHjDc0uNMaMwqBdihKQslkHCzzbrL1xcHOZ/kDw+5kXaCYGMuucINbBMMj75S0Pe1qqtEE7fM7RBuWIXBkoeeCABikjIaWvhECeemZvtO053XxuijPgXqvH11yzXv7mKqeVXw==" \
    ghcr.io/fromsi/tg_reaction:latest
```

## ENV
* TG_EMOJI_TOKEN - bot token
* TG_REACTION_DATA - json data in Deflate+Base64

## Generate Reaction Data
1. Write a new json: `{"(?i)(\\A|\\s)(hi)(\\s|\\z)": ["❤️"]}`
2. [Deflate+Base64](https://jgraph.github.io/drawio-tools/tools/convert.html) conversion of json: `q1bSsM/U1IiJcayJiSnW1MgAc4qBnCpNJSuFaKVHc5e839GvFFsLAA==`
3. Write to env: `TG_REACTION_DATA="q1bSsM/U1IiJcayJiSnW1MgAc4qBnCpNJSuFaKVHc5e839GvFFsLAA=="`