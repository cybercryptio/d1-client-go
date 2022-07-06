#!/bin/bash
set -euo pipefail

SCRIPT=$(basename $0)
CLIENT=$1
VERSION=$2

REPO="https://github.com/cybercryptio/d1-service-${CLIENT}.git"
TARGET="d1-${CLIENT}"

CURRENT_DIR=$(pwd)
CLIENT_DIR=$(realpath $TARGET)
CLIENT_PROTOBUF_DIR=$CLIENT_DIR/protobuf

# prepare target directories
rm -rf $CLIENT_DIR
mkdir $CLIENT_DIR
mkdir $CLIENT_PROTOBUF_DIR

# checkout service repo
rm -rf checkout
git clone --depth=1 $REPO checkout
cd checkout
git fetch --quiet --tags
git checkout --quiet $VERSION
COMMIT_ID=$(git rev-parse HEAD)
SRC_DIR=$(pwd)

# copy client source files
cd $SRC_DIR/client
find . -name \*.go -exec cp --parents {} $CLIENT_DIR \;

# copy protobuf source files
cd $SRC_DIR/protobuf
find . -name \*.go -exec cp --parents {} $CLIENT_PROTOBUF_DIR \;

# clean up temporary checkout directory
cd $CURRENT_DIR
rm -rf checkout

# perform text replacements
REPLACEMENTS=(
    '"github.com/cybercryptio/d1-service-generic/client"=>"github.com/cybercryptio/d1-client-go/d1-generic"'
    '"github.com/cybercryptio/d1-service-generic/protobuf"=>"github.com/cybercryptio/d1-client-go/d1-generic/protobuf"'
    '"github.com/cybercryptio/d1-service-storage/client"=>"github.com/cybercryptio/d1-client-go/d1-storage"'
    '"github.com/cybercryptio/d1-service-storage/protobuf"=>"github.com/cybercryptio/d1-client-go/d1-storage/protobuf"'
)

for REPLACEMENT in "${REPLACEMENTS[@]}"; do
    OLD_VALUE=${REPLACEMENT%=>*} # drops substring from last occurrence of `=>` to end of string
    NEW_VALUE=${REPLACEMENT#*=>} # drops substring from start of string up to first occurrence of `=>`

    # NOTE: "|| true" is to suppress non-zero exit codes from grep when nothing is found
    FILES=$(grep --recursive --files-with-matches $OLD_VALUE $CLIENT_DIR || true)
    for FILE in $FILES; do
        sed -i "s|${OLD_VALUE}|${NEW_VALUE}|g" $FILE
    done
done

# add file headers with license and "DO NOT EDIT" reminders.
add_header ()
{
    FILE=$1
    cat << EOF > $FILE
// Copyright 2022 CYBERCRYPT
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by ${SCRIPT}. DO NOT EDIT.
// source: ${REPO}
// commit: ${COMMIT_ID}
// version: ${VERSION}

$(cat $FILE)
EOF
}

for FILE in $CLIENT_DIR/*.go; do
  LN_PACKAGE=$(grep 'package' $FILE -n -m 1 | cut --delimiter=":" --fields=1)
  LN_BEFORE_PACKAGE=$(expr $LN_PACKAGE - 1)
  sed -i -e 1,${LN_BEFORE_PACKAGE}d $FILE
  add_header $FILE
done

go mod tidy

COLOR_GREEN='\033[0;32m'
COLOR_NONE='\033[0m'
printf "${COLOR_GREEN}Client '${TARGET}' is now running version '${VERSION} (${COMMIT_ID})'${COLOR_NONE}\n"
