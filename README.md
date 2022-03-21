# GoCLIApp for JSON Data
A command-line interface app to decode, update &amp; encode JSON data (from UberEats) using Go.

# After downloading the folder, compile the app
go build .

# Test features of this CLI app via terminal:
./gocli get --all

./gocli get --zip 94158

./gocli add --name "Spices" --zip 94118 --url "https://www.ubereats.com/store/spices/lszSibfTSlu-uoUiB1g2aA" --dish "Dry Hot Pot" --mp 43.99

open eats.json
