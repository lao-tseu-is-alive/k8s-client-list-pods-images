#!/bin/bash
SOURCE_CODE=main.go
echo "## Extracting app name and version from code in ${SOURCE_CODE}"
APP_NAME=$(grep -E 'APP\s+=' $SOURCE_CODE| awk '{ print $3 }'  | tr -d '"')
APP_VERSION=$(grep -E 'VERSION\s+=' $SOURCE_CODE| awk '{ print $3 }'  | tr -d '"')
APP_REPOSITORY=$(grep -E 'REPOSITORY\s+=' $SOURCE_CODE| awk '{ print $3 }'  | tr -d '"')
APP_NAME_SNAKE=$(grep -E 'AppSnake\s+=' $SOURCE_CODE| awk '{ print $3 }'  | tr -d '"')
echo "## Found APP: ${APP_NAME}, VERSION: ${APP_VERSION},  in source file ${SOURCE_CODE}"
export APP_NAME APP_NAME_SNAKE APP_VERSION APP_REPOSITORY
