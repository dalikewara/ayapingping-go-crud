package rest

const RoutePing = "/ping"
const RouteAPIV1 = "/api/v1"

// User

const RouteUser = RouteAPIV1 + "/user"
const RouteUserGetAllActive = RouteUser + "/get/all-active"
const RouteUserGetDetail = RouteUser + "/get/detail/:id"
const RouteUserRegister = RouteUser + "/register"
const RouteUserLogin = RouteUser + "/login"
const RouteUserUpdate = RouteUser + "/update"
const RouteUserDelete = RouteUser + "/delete"

// Profile

const RouteProfile = RouteAPIV1 + "/profile"
const RouteProfileUpdateImage = RouteProfile + "/update/image"
