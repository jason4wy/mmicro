package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	user "github.com/jason4wy/mmicro/user/srv/proto/user"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) QueryUserByID(ctx context.Context, req *user.QueryRequest, rsp *user.QueryResponse) error {

    userID := req.UserId
    if userID == "" {
        userID = "xxxxxx"
    }
    userInfo := &user.UserInfo{
        UserId:    userID,
        Nickname:  "匿名",
        Mobile: "138********",
        AvatarUrl: "avatar.jpg",
        Gender:    1,
    }
    rsp.Code = 200
    rsp.Error = nil
    rsp.Data = userInfo

    return nil
}
