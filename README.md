# How to run
While being in the current directory, run `go run main.go`

The server will run instantly, it may ask for your autorization based on the platform.

# How to use
Just send GET request on localhost:8080.
For example, simply enter this url in your browser : `http://localhost:8080/fizzbuzz?int1=1&int2=2&limit=10&str1=fizz&str2=buzz`

# Bonus
You can use another endpoint to know what was the most used calls is, by going to `http://localhost:8080/mostUsed`
It should also tell what the number of requet made to that URL since it start up.

Has these values are stock in memory, they will reset once the server is turned off.
