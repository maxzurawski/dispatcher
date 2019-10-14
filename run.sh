#!/bin/sh

echo "#############################################"
echo "Waiting for eureka"
echo "#############################################"

while ! `nc -z discovery 8761 `; do sleep 3; done

echo "#############################################"
echo "Waiting for proxy"
echo "#############################################"

while ! `nc -z proxy 8000 `; do sleep 3; done

echo "**************************************************************"
echo "Waiting for the rabbit service to start "
echo "**************************************************************"
while ! `nc -z rabbit 15672 `; do sleep 3; done

echo "**************************************************************"
echo "Waiting for the register service to start "
echo "**************************************************************"
while ! `nc -z register 8102 `; do sleep 3; done

echo "**************************************************************"
echo "Waiting for the sensortypes service to start "
echo "**************************************************************"
while ! `nc -z sensortypes 8101 `; do sleep 3; done

echo "#############################################"
echo "Ready to rumble. Starting dispatcher service"
echo "#############################################"
/dispatcher
