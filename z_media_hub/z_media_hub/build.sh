#!/bin/bash

function buildMacos(){
  wails build -x darwin/amd64
}

function buildLinux(){
  wails build -x linux/amd64
}

function buildWindows(){
  wails build -x windows/amd64 
}

rm -rf ./build

mkdir build

buildMacos

buildLinux

buildWindows

cp -R ./scripts ./build
cp conf.json ./build/
cp VERSION ./build/