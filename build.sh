#!/usr/bin/env bash

# GENERATE FLATBUFFERS SOURCES

BLUE='\033[0;34m'
# shellcheck disable=SC2034
RED='\033[0;31m'
NC='\033[0m' # No Color

SCRIPT_PATH=$(realpath $0)
SCRIPT_DIR=$(dirname "${SCRIPT_PATH}")

# shellcheck disable=SC2034
GCC_VERSION=10.2
FLAT_BUFFERS_VERSION=1.11.0
FLAT_BUFFERS_BINARY_PACKAGE=flatbuffers/${FLAT_BUFFERS_VERSION}@kapilsh/release

echo -e "${BLUE}KRYPTO_DIR: ${KRYPTO_DIR}${NC}"

echo -e "${BLUE}"==================="${NC}"
echo -e "${BLUE}"--- FLATBUFFERS ---"${NC}"
echo -e "${BLUE}"==================="${NC}"

echo -e "${BLUE}FLATBUFFERS CONAN PACKAGE: ${FLAT_BUFFERS_BINARY_PACKAGE}${NC}"

cd "${KRYPTO_DIR}" || exit

mkdir build && cd build || exit
conan install ${FLAT_BUFFERS_BINARY_PACKAGE} -g virtualenv -scompiler.version=${GCC_VERSION}
source activate.sh
cd ..

echo -e "${BLUE}Using $(flatc --version)${NC}" || exit

echo -e "${BLUE}"--- COMPILE C++ FLATBUFFERS SOURCES ---"${NC}"
flatc --go -o "${SCRIPT_DIR}" "${SCRIPT_DIR}/serialization.fbs" || exit

rm -r build

go build