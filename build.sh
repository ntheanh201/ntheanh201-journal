#!/bin/bash

RED=$'\e[0;31m'
NC=$'\e[0m'

git fetch
CURRENT_COMMIT_TAG="$(git tag --contains HEAD)"
LATEST_TAG="$(git describe --tags --abbrev=0 | sed 's/v//g')"

if [ "$CURRENT_COMMIT_TAG" = "" ]
    then
        echo "Please enter new tag version for current commit. Latest version is ${RED}$LATEST_TAG${NC}"
        read tag
        git tag v$tag
        git push origin v$tag
        LATEST_TAG=$tag
fi

echo "====== BUILD IMAGE ======"
IMAGE_TAG="ntheanh201-journal:$LATEST_TAG"
echo $IMAGE_TAG

IMAGE_ID=$(docker build -t $IMAGE_TAG . 2>/dev/null | awk '/Successfully built/{print $NF}')
echo $IMAGE_ID

echo "====== PUSH IMAGE ======="
IMAGE_TAG_REMOTE=ntheanh201/$IMAGE_TAG
echo $IMAGE_TAG_REMOTE
docker tag $IMAGE_ID $IMAGE_TAG_REMOTE
docker push $IMAGE_TAG_REMOTE