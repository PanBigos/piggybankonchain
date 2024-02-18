#bin/sh
docker pull ghcr.io/foundry-rs/foundry:latest

docker run  -p 8545:8545 -d -v ${PWD}:/app --name pegism-forge ghcr.io/foundry-rs/foundry:latest /app/bootup.sh