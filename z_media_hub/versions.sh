#!/bin/bash

APP_VERSION=0.0.1
GO_VERSION=$(go version)
NPM_VERSION=$(npm version)

echo "${APP_VERSION}" > z_media_hub/VERSION
echo "${GO_VERSION}" > GO_VERSION
echo "${NPM_VERSION}" > NPM_VERSION

echo "=== Z Media Hub Info ==="
echo "App Version: ${APP_VERSION}"
echo "Go Version: ${GO_VERSION}"
echo "Node Info: ${NPM_VERSION}"