package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/go-micro/api/proto"

	"github.com/jason4wy/mmicro/user/api/client"
	userSrv "github.com/jason4wy/mmicro/user/srv/proto/user"
)

type User struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// User.Info is called by the API as /user/call with post body {"name": "foo"}
func (e *User) Info(ctx context.Context, req *api.Request, rsp *api.Response) error {
    log.Log("Received User.Info request")

    // extract the client from the context
    userClient, ok := client.UserFromContext(ctx)
    if !ok {
        return errors.InternalServerError("io.github.entere.api.user.info", "user client not found")
    }

    // make request
    response, err := userClient.QueryUserByID(ctx, &userSrv.QueryRequest{
        UserId: extractValue(req.Post["user_id"]),
    })
    if err != nil {
        return errors.InternalServerError("io.github.entere.api.user.user.call", err.Error())
    }

    b, _ := json.Marshal(response)

    rsp.StatusCode = 200
    rsp.Body = string(b)

    return nil
}


