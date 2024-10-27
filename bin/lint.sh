#!/bin/bash

set -u

pushd .

cd backend || exit
golangci-lint run

cd ../frontend || exit
npm run lint

popd || return


