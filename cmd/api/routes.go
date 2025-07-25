package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// Login/Logout/ForgotPassword
	router.HandlerFunc(http.MethodPost, "/v1/problem_beater/auth/login", app.loginHandler)
	router.HandlerFunc(http.MethodGet, "/v1/problem_beater/auth/logout", app.logoutUser)
	router.HandlerFunc(http.MethodPost, "/v1/problem_beater/user/forgotPassword", app.forgotPassword)

	// Profile Create, Read, Update, Delete
	router.HandlerFunc(http.MethodGet, "/v1/problem_beater/user/profile/:id", app.viewProfileHandler)
	router.HandlerFunc(http.MethodPost, "/v1/problem_beater/user/createProfile", app.createProfileHandler)
	router.HandlerFunc(http.MethodPut, "/v1/problem_beater/user/updateProfile", app.updateProfileHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/problem_beater/user/deleteProfile/:id", app.deleteProfileHandler)

	// Class Create, Read, Update, Delete
	router.HandlerFunc(http.MethodPost, "/v1/problem_beater/createClass", app.createClassHandler)
	router.HandlerFunc(http.MethodGet, "/v1/problem_beater/getAllClasses", app.getAllClasses)
	router.HandlerFunc(http.MethodPut, "/v1/problem_beater/updateClass", app.updateClass)
	router.HandlerFunc(http.MethodDelete, "/v1/problem_beater/deleteClass/:id", app.deleteClass)

	router.HandlerFunc(http.MethodGet, "/v1/problem_beater/student/test", app.testHandler)

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
