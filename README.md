Dear Reader,
Please follow the instructions in this file, if you want to smoothly run the application

Step 1: Install Golang in your computer. (https://go.dev/doc/install)
Step 2: Open up the terminal. Go to the application directory, and make sure you initiate Go modules by "go mod init assignment1" 
Step 3: Open your favorite code editor and find the folder of this application
Step 4: Build 2 Go files "server.go" and "client.go" by typing in terminal "go build server.go" and then "go build client.go"
Step 5: Open 4 different terminal windows. 3 of which will be used as server windows, and 1 as client window  
Step 6: In your server terminals, make sure you go to the directory folder of our application. In terminal 1 type "./server 8080". In terminal 2 type "./server 8081". In terminal 3 type "./server 8082". (You can use any other port, if those are taken). 
Step 7: In the last terminal type "./client localhost:8080 5". Instead of 8080 use any of the ports that you initiated, and instead of 5 you can use any number to double it. 

Done!