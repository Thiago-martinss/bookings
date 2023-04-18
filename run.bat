go build bookings ./cmd/web
./bookings -dbname=bookings -dbuser=postgres -dbpass=root -cache=false -production=false