#!/bin/sh
sudo ./scripts/libvirt destroy
sudo ./scripts/devnet destroy
sudo rkt gc --grace-period=0
