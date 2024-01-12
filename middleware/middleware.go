package middleware

type arg func(param any)

func Routes(at string, f arg) {
	f(at)
}
