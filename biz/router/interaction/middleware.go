// Code generated by hertz generator.

package interaction

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

func _interactionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _likeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _likeactionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
	//return nil
}

func _likelistMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
	//return nil
}

func _commentMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentactionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
	//return nil
}

func _commentlistMw() []app.HandlerFunc {
	// your code...
	return nil
}
