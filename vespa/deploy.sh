#!/bin/sh

VESPA_HOST="${VESPA_HOST:-localhost}"
VESPA_PORT="${VESPA_PORT:-19071}"

echo "Deploying Vespa application package to http://${VESPA_HOST}:${VESPA_PORT}..."

cd "$(dirname "$0")/src/main/vespa" || exit 1

zip -r /tmp/application.zip .

curl --silent --fail --show-error \
  --header "Content-Type: application/zip" \
  --data-binary @/tmp/application.zip \
  "http://${VESPA_HOST}:${VESPA_PORT}/application/v2/tenant/default/prepareandactivate"

echo ""
echo "Vespa application deployed successfully."
