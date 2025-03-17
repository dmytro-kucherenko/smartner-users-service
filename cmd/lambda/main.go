package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Dmytro-Kucherenko/smartner-users-service/internal"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/modules/common"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/dtos"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
)

func handle(request events.APIGatewayProxyRequest) (meta server.RequestMeta) {
	if request.RequestContext.Authorizer == nil {
		return
	}

	session, err := common.DecodeStruct[server.Session](request.RequestContext.Authorizer)
	if err != nil {
		return
	}

	meta.Session = &session

	return
}

func Init(create func(logger types.Logger, meta server.RequestMeta) (internal.StartupOptions, error)) {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		logger := log.New("Init")
		meta := handle(request)
		options, err := create(logger, meta)

		if err != nil {
			var response dtos.ErrorResponse
			httpErr, ok := err.(*errors.HttpError)

			if ok {
				response = dtos.ErrorResponse{
					Status:  httpErr.Status(),
					Message: httpErr.Error(),
					Details: httpErr.Details(),
				}
			} else {
				response = dtos.ErrorResponse{
					Status:  http.StatusInternalServerError,
					Message: "internal server error",
				}

				logger.Error(err.Error())
			}

			body, _ := json.Marshal(response)

			return events.APIGatewayProxyResponse{
				StatusCode: response.Status,
				Headers:    map[string]string{"Content-Type": "application/json"},
				Body:       string(body),
			}, nil
		}

		if options.OnlyConfig {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
			}, nil
		}

		return ginadapter.New(options.Router).ProxyWithContext(ctx, request)
	})
}

// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey	JWTAuth
// @in							header
// @name						Authorization
// @description				JWT authorization guard
func main() {
	Init(internal.Init)
}
