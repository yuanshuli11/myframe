#!/bin/bash
#
# Code coverage generation
COVERAGE_DIR="${COVERAGE_DIR:-coverage}"
if [ $# -gt 2 ]; then
  OUTPUT=$1
  OUTPUT_DIR=$2
  PKG_LIST=${@:3}
elif [ $# -gt 1 ]; then
  OUTPUT=$1
  PKG_LIST=${@:2}
else
  PKG_LIST=${@:1}
fi
# Create the coverage files directory
mkdir -p "$COVERAGE_DIR";
# Merge the coverage profile files
echo 'mode: count' > "${COVERAGE_DIR}"/coverage.cov ;
# Create a coverage file for each package
for package in ${PKG_LIST}; do
    go test -covermode=count -coverprofile "${COVERAGE_DIR}/${package##*/}.cov" "$package";
    tail -q -n +2 "${COVERAGE_DIR}/${package##*/}.cov" >> "${COVERAGE_DIR}"/coverage.cov ;
done ;
# Display the global code coverage
go tool cover -func="${COVERAGE_DIR}"/coverage.cov ;
# If needed, generate HTML repor
if [[ "$OUTPUT" == "html" ]]; then
    if [[ "$OUTPUT_DIR" != "" ]]; then
        mkdir -p $OUTPUT_DIR
        go tool cover -html="${COVERAGE_DIR}"/coverage.cov -o "$OUTPUT_DIR"/coverage.html ;
    else
        go tool cover -html="${COVERAGE_DIR}"/coverage.cov -o coverage.html ;
    fi
fi
# Remove the coverage files directory
rm -rf "$COVERAGE_DIR";

