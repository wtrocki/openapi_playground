## api::generate_sdk
## Positional Args:
## file_name - location of the input openapi file
## output_path - folder containing generated code
## package_name - name of the package
api::generate_sdk() {
    local file_name=$1
    local output_path=$2
    local package_name=$3
    local additional_properties=$4
     
    echo "Validating OpenAPI ${file_name}"
    npx @openapitools/openapi-generator-cli validate -i "$file_name"

    echo "Generating source code based on ${file_name}"

    # remove old generated source code
    rm -Rf $output_path
    
    npx @openapitools/openapi-generator-cli generate -g go -i \
     "$file_name" -o "$output_path" \
    --package-name="${package_name}" \
    --additional-properties=$additional_properties \
    --ignore-file-override=.openapi-generator-ignore
}
 