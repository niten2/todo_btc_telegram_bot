### Description

  - /start
  - /help (list commands)

### Details

  - Bot @coint_info_bot
  - https://web.telegram.org/#/im?p=@coint_info_bot

#### Features

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

#### Deploy (asible)

  - cd ./ansible
  - cp hosts.sample hosts
  - edit ./ansible/hosts
  - ansible-playbook setup (install docker, move app on remote server)
