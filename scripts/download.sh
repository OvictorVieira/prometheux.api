#!/bin/bash
set -e
version=$(curl -fs https://api.github.com/repos/c9s/bbgo/releases/latest | awk -F '"' '/tag_name/{print $4}')
arch=""
case $(uname -m) in
  x86_64 | ia64) arch="amd64";;
  arm64 | aarch64 | arm) arch="arm64";;
  *)
    echo "unsupported architecture: $(uname -m)"
    exit 1;;
esac
osf=$(uname | tr '[:upper:]' '[:lower:]')
dist_file=bbgo-$version-$osf-$arch.tar.gz

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

function warn()
{
    echo -e "${YELLOW}$@${NC}"
}

function error()
{
    echo -e "${RED}$@${NC}"
}

function info()
{
    echo -e "${GREEN}$@${NC}"
}

info "downloading..."
curl -O -L https://github.com/OvictorVieira/promeheux.api/releases/download/$version/$dist_file
tar xzf $dist_file
mv bbgo-$osf-$arch bbgo
chmod +x bbgo
info "downloaded successfully"
