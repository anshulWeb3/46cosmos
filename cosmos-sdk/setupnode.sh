


#!/bin/bash

KEYS="alice"
CHAINID="sim_904-3 "
MONIKER="simnode"
# KEYRING="test"
LOGLEVEL="info"
# Set dedicated home directory for the wasmd instance
HOMEDIR="/data/kkkk/simd"

# Path variables
CONFIG=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

# validate dependencies are installed
command -v jq >/dev/null 2>&1 || {
	echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"
	exit 1
}

# used to exit on first error
set -e

# Reinstall daemon
make build

# User prompt if an existing local node configuration is found.
if [ -d "$HOMEDIR" ]; then
	printf "\nAn existing folder at '%s' was found. You can choose to delete this folder and start a new local node with new keys from genesis. When declined, the existing local node is started. \n" "$HOMEDIR"
	echo "Overwrite the existing configuration and start a new local node? [y/n]"
	read -r overwrite
else
	overwrite="Y"
fi

# Setup local node if overwrite is set to Yes, otherwise skip setup
if [[ $overwrite == "y" || $overwrite == "Y" ]]; then
	# Remove the previous folder
	rm -rf "$HOMEDIR"

	# Set client config
	./build/simd config --home "$HOMEDIR"
	./build/simd config chain-id $CHAINID --home "$HOMEDIR"

	./build/simd keys add $KEYS --home "$HOMEDIR"

	./build/simd init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR"

	# # Change parameter stake denominations to tSTKC
	# jq '.app_state["staking"]["params"]["bond_denom"]="stake"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	# jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="stake"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	# # jq '.app_state["gov"]["params"]["min_deposit"][0]["denom"]="stake"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	# jq '.app_state["crisis"]["constant_fee"]["denom"]="stake"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
    # jq '.app_state["mint"]["params"]["mint_denom"]="stake"' >"$TMP_GENESIS" "$GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Set claims decays
	jq '.app_state["gov"]["deposit_params"]["max_deposit_period"]="200s"' >"$TMP_GENESIS" "$GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["gov"]["voting_params"]["voting_period"]="200s"' >"$TMP_GENESIS" "$GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["staking"]["params"]["unbonding_time"]="200s"' >"$TMP_GENESIS" "$GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	#changes status in app,config files
        sed -i 's/prometheus = false/prometheus = true/' "$CONFIG"
        sed -i 's/prometheus-retention-time  = "0"/prometheus-retention-time  = "1000000000000"/g' "$APP_TOML"
        sed -i 's/enabled = false/enabled = true/g' "$APP_TOML"
        sed -i 's/enable = false/enable = true/g' "$APP_TOML"
        sed -i 's/swagger = false/swagger = true/g' "$APP_TOML"

	# Allocate genesis accounts (cosmos formatted addresses)
	./build/simd add-genesis-account $KEYS 10000000000000000000000000000stake  --home "$HOMEDIR"

	# Sign genesis transaction
	./build/simd gentx ${KEYS} 100000000000000000stake --chain-id $CHAINID --home "$HOMEDIR"
	
	# Collect genesis tx
	./build/simd collect-gentxs --home "$HOMEDIR"

	# Run this to ensure everything worked and that the genesis file is setup correctly
	./build/simd validate-genesis --home "$HOMEDIR"

fi

# Start the node
./build/simd start --home "$HOMEDIR"