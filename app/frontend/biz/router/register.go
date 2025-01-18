// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	auth "github.com/MyGoFor/E-commerce/app/frontend/biz/router/auth"
	cart "github.com/MyGoFor/E-commerce/app/frontend/biz/router/cart"
	category "github.com/MyGoFor/E-commerce/app/frontend/biz/router/category"
	home "github.com/MyGoFor/E-commerce/app/frontend/biz/router/home"
	product "github.com/MyGoFor/E-commerce/app/frontend/biz/router/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	cart.Register(r)

	category.Register(r)

	product.Register(r)

	auth.Register(r)

	home.Register(r)
}
