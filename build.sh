#!/usr/bin/env bash

# GENERATE FLATBUFFERS SOURCES

BLUE='\033[0;34m'
# shellcheck disable=SC2034
RED='\033[0;31m'
NC='\033[0m' # No Color

SCRIPT_PATH=$(realpath $0)
SCRIPT_DIR=$(dirname "${SCRIPT_PATH}")

GCC_VERSION=10.2

echo -e "${BLUE}"==================="${NC}"
echo -e "${BLUE}"--- FLATBUFFERS ---"${NC}"
echo -e "${BLUE}"==================="${NC}"

rm -rf ${SCRIPT_DIR}/build && mkdir ${SCRIPT_DIR}/build && cd ${SCRIPT_DIR}/build || exit
conan install ${SCRIPT_DIR} -g virtualenv -g cmake -s compiler=gcc  -s compiler.version=${GCC_VERSION} -s build_type=Release --profile ${SCRIPT_DIR}/build.profile --build=missing
source activate.sh && echo -e "${BLUE}Using $(flatc --version)${NC}" || exit
cd ..

echo -e "${BLUE}"--- COMPILE GO FLATBUFFERS SOURCES ---"${NC}"
flatc --go -o "${SCRIPT_DIR}" "${SCRIPT_DIR}/serialization.fbs" || exit

rm -r build

echo -e "${BLUE}"--- BUILDING EXECUTABLE ---"${NC}"

go build .

echo -e "${BLUE}"--- DONE! ---"${NC}"