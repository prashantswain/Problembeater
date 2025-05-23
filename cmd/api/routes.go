package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	// router.HandleFunc("GET /", "", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Welcome to problem beater Apis"))
	// })

	// router.NotFound = http.HandlerFunc(app.notFoundResponse)
	// router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/auth/login", app.loginHandler)

	// router.HandlerFunc(http.MethodGet, "/v1/user/profile/:id", app.viewProfileHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/user/profile", app.createProfileHandler)
	// router.HandlerFunc(http.MethodPut, "/v1/user/profile", app.updateProfileHandler)
	// router.HandlerFunc(http.MethodDelete, "/v1/user/profile", app.deleteProfileHandler)

	router.HandlerFunc(http.MethodGet, "/v1/student/test", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/user/address", app.testHandler)
	// router.HandlerFunc(http.MethodPut, "/v1/user/address", app.testHandler)
	// router.HandlerFunc(http.MethodDelete, "/v1/user/address", app.testHandler)

	// router.HandlerFunc(http.MethodGet, "/v1/deliveryslot", app.testHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/vendor/bycategory", app.testHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/varients/trending", app.testHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/varients/find?query=", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/shippingmethod/pincode", app.testHandler)

	// router.HandlerFunc(http.MethodPost, "/v1/banner", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/banner/slider2", app.testHandler)

	// router.HandlerFunc(http.MethodPost, "/v1/product/category/all/product", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/public/varients/trending?pincode=", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/product/category/subcat/all/", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/product/category/forbrand/", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/product/varients/bycategory/", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/product/varients/bycategorypage/", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/corder?page=", app.testHandler)

	// router.HandlerFunc(http.MethodPost, "/v1/product/varients/related", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/product/review/byvarient/", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/product/review", app.testHandler)

	// router.HandlerFunc(http.MethodPost, "/v1/cart", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/cart", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/cart", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/cart", app.testHandler)

	// router.HandlerFunc(http.MethodPost, "/v1/tempcart", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/tempcart/", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/tempcart/", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/tempcart/", app.testHandler)

	// router.HandlerFunc(http.MethodPost, "/v1/cart/placeorder", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/orders/wishlist", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/product/varients/topdiscounted", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/orders/wishlist", app.testHandler)

	// router.HandlerFunc(http.MethodPost, "/v1/aboutus", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/faq", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/cart/applycoupon", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/privacypolicy", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/tnc", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/returnpolicy", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/payment/walletrecharge", app.testHandler)

	// router.HandlerFunc(http.MethodPost, "/v1/orders/config", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/orders/discount", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/payment/verify", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/corder/status/", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/corder/action", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/product/brand", app.testHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/product/varients/bybrand/", app.testHandler)

	return router
}
