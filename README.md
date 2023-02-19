# Relayer service from BTC to CORE

## Relayers in Core

Relayers in Core are responsible for relaying BTC block headers onto the network via the on-chain light client. Relayers must both register and pass verification in order to receive rewards.

For more information about relayers, please read [PoW in Core blockchain](https://docs.coredao.org/core-white-paper-v1.0.5/satoshi-plus-consensus/proof-of-work).

## Quick Start

**Note**: Requires [Go 1.17+](https://golang.org/dl/)

### Setup config

1. Edit `config/config.json` 
    1. Fill in your private key to `core_config.private_key`.
    2. Edit btc_config.rpc_addrs, fill in btc rpc address. Modify sleep_second, which is the interval to refresh btc highestHeight. Modify data_seed_deny_service_threshold, which is the interval to send telegram alert when refreshing btc highestHeight fails. 
    3. Edit core_config.providers, fill in core rpc address. Modify sleep_second, which is the interval to refresh core highestHeight. Modify data_seed_deny_service_threshold, which is the interval to send telegram alert when refreshing core highestHeight fails.
    4. If gas_limit is not enough, gas_increase will be added and a retry will be taken.
    5. Recursion_height is the number of blocks to go back and check on btc network based on the newest height.
2. Transfer enough CORE to the relayer account.
    1. 100 CORE as relayer registration fees.
    2. More than 10 CORE as transaction fees.
3. Send telegram message when the balance of relayer account is too low. This is an example config:
    ```json
    {
        "enable_alert": true,
        "enable_heart_beat": false,
        "interval": 300,
        "telegram_bot_id": "your_bot_id",
        "telegram_chat_id": "your_chat_id",
        "balance_threshold": "1000000000000000000",
        "sequence_gap_threshold": 10
    }
    ```
   Please refer to [telegram_bot](https://www.home-assistant.io/integrations/telegram_bot) to setup your telegram bot. If you don't want this feature, just set `enable_alert` to false.

### Build

#### Build Binary:
```shell script
make build
```

### Run

Run locally:
```shell script
./btc-relayer
```
