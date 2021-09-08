#!/bin/bash

go build -o bookings cmd/web/*.go
./bookings -dbname= -dbuser= -dbpass=
