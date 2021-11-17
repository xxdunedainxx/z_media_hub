#!/bin/bash

HOME=$(pwd)

function buildMacos(){
  wails build -x darwin/amd64
}

function buildLinux(){
  wails build -x linux/amd64
}

function buildWindows(){
  wails build -x windows/amd64 
  mv z_media_hub-res.syso ./build
  mv z_media_hub.exe.manifest ./build
  mv z_media_hub.ico ./build
  mv z_media_hub.rc ./build
  mv appicon.png ./build
}

function getVersion(){
  cd ..

  ./versions.sh

  cd $HOME
}

getVersion

rm -rf ./build

mkdir build

buildMacos

buildLinux

buildWindows

cp -R ./scripts ./build
cp conf.json ./build/
cp VERSION ./build/