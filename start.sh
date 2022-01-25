export NODE_OPTIONS=--openssl-legacy-provider
cd core ; go run main.go &
cd ../app ; npm run start &

# unset NODE_OPTIONS