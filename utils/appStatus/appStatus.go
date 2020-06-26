package appStatus

const (
	Success                 = 0
	StatusNotYetImplemented = 500
	UnknownError            = 600
	Error                   = 1
)

var statusText = map[int]string{
	Success:                 "Sukses",
	StatusNotYetImplemented: "Belum Diimplementasikan",
	UnknownError:            "Kesalahan tidak diketahui",
	Error:                   "Terjadi Kesalahan",
}

func StatusText(code int) string {
	return statusText[code]
}
