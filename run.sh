#!/bin/bash

go build -o bookings web/*.go
./bookings -dbname=bookings -dbuser= -dbpass=
