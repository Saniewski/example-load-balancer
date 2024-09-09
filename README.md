# Example load balancer

A simple example of a load balancer using Docker Swarm, Caddy reverse proxy, Go, and Vue3 + Vite.

## How to run

1. Clone this repository
2. Run `docker swarm init`.
3. If the images specified in the `swarm.yaml` are not available, you should build them yourself.
  To build the images, navigate to `example-go-service` directory, and run `docker build -f Dockerfile -t username/example-go-service:1.0.0 .` and `docker push username/example-go-service:1.0.0`.
  Make sure to replace `username` with your Docker Hub username.
  Do the same for Caddy, i.e. navigate to `example-caddy` directory, and run `docker build -f Dockerfile -t username/example-caddy:1.0.0 .` and `docker push username/example-caddy:1.0.0`.
4. Run `docker stack deploy -c swarm.yaml example`.
5. Run `watch -n 1 go run example-client/cmd/client/main.go`.
  You should see the client hitting the server and reporting the response count for each backend service replica's hostname.
  At this stage, the client shows one hostname with a response count of 100 per second.
6. In a separate terminal tab, run `docker service scale example_example-go-service=4`.
  You should see the client updating the response count for each backend service replica's hostname.
  At this stage, the client shows four hostnames, each with a response count of 100 roughly divided by 4.
7. Run `docker service scale example_example-go-service=1`.
  The backend hostnames should drop back to 1 with a response count of 100.
8. Run `docker stack rm example` to drop the deployed stack.
9. Run `docker swarm leave --force` to leave the swarm.

A `compose.yaml` file is also provided for local development and testing with the frontend in Vue.
To run the frontend locally, run `cd example-front`, `pnpm install`, and `pnpm dev`.
