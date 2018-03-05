# Support Messenger Bot

Please note that a ```conf.json``` file is required to make this work.

```json
{
    "appSecret": "anAppSecret",
    "token": "aToken",
    "pageAccessToken": "aPageAccessToken"
}
```

# About it

- Builds and run the bot server
```sh
# Builds and run the bot server
make start

# Starts my ngrok for webhook tests purpose
make ng
```