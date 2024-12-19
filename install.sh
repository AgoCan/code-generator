#!/bin/bash

# GitHub release 
REPO_OWNER="go-cheetah"
REPO_NAME="cheetah"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')  # linux、darwin、windows
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
elif [[ "$ARCH" == "aarch64" ]]; then
  ARCH="arm64"
else
  echo "Unsupported architecture: $ARCH"
  exit 1
fi

if [[ -e cheetah ]]
then
    echo "Error: ./cheetah already exists!"
    exit 1
fi

LATEST_VERSION=$(curl -s https://api.github.com/repos/$REPO_OWNER/$REPO_NAME/releases/latest | grep '"tag_name":' | sed -E 's/.*"tag_name": "(.*)".*/\1/')

if [[ -z "$LATEST_VERSION" ]]; then
  echo "get new version error."
  exit 1
fi

DOWNLOAD_URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$LATEST_VERSION/$REPO_NAME-$OS-$ARCH"

echo "curl -o ./cheetah $DOWNLOAD_URL ..."
curl -sL -o ./cheetah $DOWNLOAD_URL
chmod +x ./cheetah

echo "Use './cheetah -h' to view help."