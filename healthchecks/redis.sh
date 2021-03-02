#!/bin/sh
set -eo pipefail

ping="$(echo -e "auth ingilizce-kelime-go samet123\nping" | redis-cli)"
case "$ping" in
*PONG* ) exit 0 ;;
esac
exit 1

