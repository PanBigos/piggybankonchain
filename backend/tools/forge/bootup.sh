#bin/sh

# on arm image the nc is not preinstalled
arch=$(uname -m)

if [ "$arch" = "aarch64" ] || [ "$arch" = "armv7l" ]; then
    echo "Running on ARM architecture"
    apt-get update
    # for pulling git repo
    apt-get install apt-transport-https ca-certificates -y 
    apt-get install ca-certificates -y
    # for detecing bootup of anvil
    apt-get install -y netcat
    # update curl for health-checks
    apt-get install -y curl 
else
    echo "Running on a non-ARM architecture"
fi


git config --global user.email "you@example.com"
git config --global user.name "Your Name"

# Check if FORK_URL is set
if [ -n "$FORK_URL" ]; then
    # If FORK_URL is set, use it with the -f flag
    echo "FORK_URL is set to: $FORK_URL. Using it with -f flag."
    /app/deploy.sh &
    anvil --host="0.0.0.0" -f "$FORK_URL"
else
    # If FORK_URL is not set, just run anvil without the -f flag
    echo "FORK_URL is not set. Running anvil without -f flag."
    /app/deploy.sh &
    anvil --host="0.0.0.0"
fi