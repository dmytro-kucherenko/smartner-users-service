FROM golang:1.24.2-alpine AS builder

ARG PORT=8000

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bootstrap ./cmd/ecs/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bootstrap .

EXPOSE $PORT
EXPOSE $PORT

CMD ["./bootstrap"]

# aws ecr get-login-password --region eu-central-1 | docker login --username AWS --password-stdin 269733607457.dkr.ecr.eu-central-1.amazonaws.com

# # Create repository if needed in cloudformation
# aws ecr create-repository --repository-name services --region eu-central-1

# # Build the Docker image
# docker build --build-arg PORT=8000 -t users-service .

# # Tag the image
# docker tag users-service:latest 269733607457.dkr.ecr.eu-central-1.amazonaws.com/users-service:latest

# # Push the image to ECR
# docker push 269733607457.dkr.ecr.eu-central-1.amazonaws.com/users-service:latest

# aws ecs update-service --cluster ECSCluster --service UsersService --force-new-deployment

# aws cloudformation deploy \
#   --template-file template.yaml \
#   --stack-name golang-grpc-service \
#   --parameter-overrides \
#     NetworkStackName=your-network-stack \
#     ServiceName=golang-grpc-service \
#     ContainerImage=<your-account-id>.dkr.ecr.us-east-1.amazonaws.com/golang-grpc-server:latest \
#     ContainerPort=8000
