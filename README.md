# Guide
## Building and Deploying
For the server and logger
  1. Build and push the docker image to your repository
  1. update the `server.yml` and `logger.yml` to point to the correct image
  1. Deploy server and logger by applying the manfiests, e.g. `kubectl apply -f server.yml`
  
Once deployed you can check its working correctly by checking the logs
   outputs of the two pods

## Using telepresence
Now that the two application are deployed and running you can use telepresence
to switch the server to a local version. 
1. Update the `server/main.go` to return a different response
2. Run `telepresence --swap-deployment server --expose 8080 --run go run
   server/main.go`

The traffic should now being forwarded to your local server instead. Check the
logs of the logger to see that its not receiving your new response back.
