#!/bin/bash -ex

debug=$(snapctl get debug)
logger "device-mqtt: debug: $debug"

autostart=$(snapctl get autostart)
logger "device-mqtt: autostart: $autostart"

exec "$@"
