#!/bin/bash

apt-get install -y gccgo-go

mkdir -p ~/gospace
echo "export GOPATH=~/gospace" >> ~/.bashrc
echo "export PATH=~/gospace/bin:$PATH" >> ~/.bashrc

mkdir -p ~/gospace/bin
mkdir -p ~/gospace/pkg
mkdir -p ~/gospace/src/github.com/migibert/netgrapher

cp -r /vagrant/* ~/gospace/src/github.com/migibert/netgrapher


