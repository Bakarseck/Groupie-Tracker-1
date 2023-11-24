package app

func Errorrac(code int) Error {
	var rr Error
	switch code {
	case 404:
		return Error{
			Code:     code,
			ErrorMsg: "Page no found!!",
			Image:    "./asset/img/404.png",
		}
	case 500:
		return Error{
			Code:     code,
			ErrorMsg: "Internal Server Error!!",
			Image:    "./asset/img/500.png",
		}
	case 400:
		return Error{
			Code:     code,
			ErrorMsg: "Bad Request!!",
			Image:    "./asset/img/400.png",
		}
	}
	return rr
}
