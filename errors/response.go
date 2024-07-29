package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ResponseError interface {
	error
}

func NewBadRequestError(msg string) *BadRequestError {
	return &BadRequestError{
		RtnCode: "000400",
		RtnMsg:  msg,
		ErrStr:  "",
	}
}

func NewInternationalBadRequestError(msg string, msgEnglish string) *BadRequestError {
	return &BadRequestError{
		RtnCode:       "000400",
		RtnMsg:        msg,
		RtnMsgEnglish: msgEnglish,
		ErrStr:        "",
	}
}

func NewBadRequestMsgError(msg string, err error) *BadRequestError {
	return &BadRequestError{
		RtnCode: "000400",
		RtnMsg:  msg,
		ErrStr:  err.Error(),
	}
}

func NewInternationalBadRequestMsgError(msg string, msgEnglish string, err error) *BadRequestError {
	return &BadRequestError{
		RtnCode:       "000400",
		RtnMsg:        msg,
		RtnMsgEnglish: msgEnglish,
		ErrStr:        err.Error(),
	}
}

type BadRequestError struct {
	RtnCode       string `json:"rtnCode"`
	RtnMsg        string `json:"rtnMsg"`
	RtnMsgEnglish string `json:"rtnMsgEnglish"`
	ErrStr        string `json:"detailError"`
}

func (r BadRequestError) Error() string {
	return fmt.Sprintf("BadRequest! 错误码:'%s',错误原因：%s", r.RtnCode, r.RtnMsg)
}

func NewInternalServerError(msg string) *InternalServerErrorError {
	return &InternalServerErrorError{
		RtnCode: "000500",
		RtnMsg:  msg,
		ErrStr:  "",
	}
}

func NewInternationalInternalServerMsgError(msg string, msgEnglish string, err error) *InternalServerErrorError {
	return &InternalServerErrorError{
		RtnCode:       "000500",
		RtnMsg:        msg,
		RtnMsgEnglish: msgEnglish,
		ErrStr:        err.Error(),
	}
}

func NewInternalServerMsgError(msg string, err error) *InternalServerErrorError {
	return &InternalServerErrorError{
		RtnCode: "000500",
		RtnMsg:  msg,
		ErrStr:  err.Error(),
	}
}

type InternalServerErrorError struct {
	RtnCode       string `json:"rtnCode"`
	RtnMsg        string `json:"rtnMsg"`
	RtnMsgEnglish string `json:"rtnMsgEnglish"`
	ErrStr        string `json:"detailError"`
}

func (r InternalServerErrorError) Error() string {
	return fmt.Sprintf("Internal Server Error! 错误码:'%s',错误原因：%s，详细错误:%s", r.RtnCode, r.RtnMsg, r.ErrStr)
}

type unknownErrorError struct {
	RtnCode       string `json:"rtnCode"`
	RtnMsg        string `json:"rtnMsg"`
	RtnMsgEnglish string `json:"rtnMsgEnglish"`
	ErrStr        string `json:"detailError"`
}

func newUnknownErrorError(err error) *unknownErrorError {
	return &unknownErrorError{
		RtnCode:       "999999",
		RtnMsg:        "发生未知错误",
		RtnMsgEnglish: "An unknown error occurred",
		ErrStr:        err.Error(),
	}
}

func (r unknownErrorError) Error() string {
	return fmt.Sprintf("UnkownError! 错误码:'%s',错误原因：%s", r.RtnCode, r.ErrStr)
}

type StandardHttpError struct {
	Code int
	Msg  string
}

func (r StandardHttpError) Error() string {
	return fmt.Sprintf("StandardHttpError! httpCode:'%s',错误原因：%s", r.Code, r.Msg)
}

func NewStandardHttpError(code int, err error) *StandardHttpError {
	return &StandardHttpError{
		Code: code,
		Msg:  err.Error(),
	}
}

func NewStandardHttpMsg(code int, msg string) *StandardHttpError {
	return &StandardHttpError{
		Code: code,
		Msg:  msg,
	}
}

func HttpErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	//自定义错误统一处理添加在这里
	var badRequestError = NewBadRequestError("")
	var internalServerErrorError = NewInternalServerError("")
	var standardHttpError = NewStandardHttpMsg(0, "")

	if errors.As(err, &badRequestError) ||
		errors.As(err, &internalServerErrorError) {
		keErrorHandle(w, err)
	} else if errors.As(err, &standardHttpError) {
		standardErrorHandle(w, err.(*StandardHttpError))
	} else {
		unknownErrorHandle(w, err)
	}
}

func keErrorHandle(w http.ResponseWriter, err error) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	bodyBytes, _ := json.Marshal(err)
	_, _ = w.Write(bodyBytes)
}

func standardErrorHandle(w http.ResponseWriter, err *StandardHttpError) {
	w.WriteHeader(err.Code)
	w.Header().Set("Content-Type", "application/text")
	_, _ = w.Write([]byte(err.Msg))
}

func unknownErrorHandle(w http.ResponseWriter, err error) {
	unknownError := newUnknownErrorError(err)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	bodyBytes, _ := json.Marshal(unknownError)
	_, _ = w.Write(bodyBytes)
}

type emptyErr struct{}

func (e *emptyErr) Error() string { return "" }
