DIR=`dirname "$0"`

# ~/.{app.AppName}
rm -rf ~/.asset-chain

# initial new node
asset-chaind init user1 --chain-id bandchain
echo "lock nasty suffer dirt dream fine fall deal curtain plate husband sound tower mom crew crawl guard rack snake before fragile course bacon range" \
    | asset-chaind keys add validator --recover --keyring-backend test

cp ./docker-config/single-validator/priv_validator_key.json ~/.asset-chain/config/priv_validator_key.json
cp ./docker-config/single-validator/node_key.json ~/.asset-chain/config/node_key.json

# add accounts to genesis
bandd add-genesis-account user1 10000000000000uusd --keyring-backend test

# register initial validators
bandd gentx user1 100000000stake \
    --chain-id bandchain \
    --keyring-backend test

# 

# collect genesis transactions
bandd collect-gentxs