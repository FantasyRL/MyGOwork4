// Code generated by hertz generator.

package video

import (
	"bibi/biz/mw/jwt"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _bibiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _videoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _hotvideoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listvideoMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
	//return nil
}

func _searchvideoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _putvideoMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
	//return nil
}

func _vMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listvideosbyidMw() []app.HandlerFunc {
	// your code...
	return nil
}
