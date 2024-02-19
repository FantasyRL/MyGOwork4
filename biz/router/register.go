// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	follow "bibi/biz/router/follow"
	interaction "bibi/biz/router/interaction"
	user "bibi/biz/router/user"
	video "bibi/biz/router/video"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	follow.Register(r)

	interaction.Register(r)

	video.Register(r)

	user.Register(r)
}
