#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source $SCRIPT_DIR/_functions.sh

npx @openapitools/openapi-generator-cli version-manager set 5.2.0
echo "Generating SDKs"
additional_properties="generateInterfaces=true,enumClassPrefix=true,structPrefix=true"

## Various reproducers/tests
OPENAPI_FILENAME=${OPENAPI_FILENAME:-.openapi/content-type.json}
PACKAGE_NAME="contenttype"
OUTPUT_PATH="contenttype"
 
api::generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME $additional_properties
