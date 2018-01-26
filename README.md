**Setup**
---

Bot @coint_info_bot
https://web.telegram.org/#/im?p=@coint_info_bot

#### Feature
- show current btc by schedule
- show current btc wallet by schedule
- add a reminder when the price changes (poloniex)

#### Development mode
- cp .env.sample .env
- go run

#### Production mode
- go build .

- cp .env.sample .env.production
- edit .env.production

after you may use ansible script
- cd ./ansible
- cp hosts.sample hosts
- edit ./ansible/hosts
- ansible-playbook setup (install docker, move app on remote server)

### Describe about bot
- /start
- /help (list commands)
