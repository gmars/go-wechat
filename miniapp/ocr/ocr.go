package ocr

import (
	"go-wechat/core"
	"mime/multipart"
)

type Ocr struct {
	request *core.ApiRequest
}

func NewOcr(accessToken core.AccessToken) *Ocr {
	return &Ocr{request: core.NewApiRequest(accessToken)}
}

// AiCrop 图片智能剪裁
func (s *Ocr) AiCrop(imageUrl string, file *multipart.FileHeader) (*AICropRes, error) {
	var (
		res AICropRes
	)
	err := s.ocrCli("/cv/img/aicrop", imageUrl, file, nil, &res)
	return &res, err
}

// ScanQRCode 条形码识别
func (s *Ocr) ScanQRCode(imageUrl string, file *multipart.FileHeader) (*ScanQRCodeRes, error) {
	var (
		res ScanQRCodeRes
	)
	err := s.ocrCli("/cv/img/qrcode", imageUrl, file, nil, &res)
	return &res, err
}

// SuperResolution 高清化图片
func (s *Ocr) SuperResolution(imageUrl string, file *multipart.FileHeader) (string, error) {
	var (
		res SuperResolutionRes
	)
	err := s.ocrCli("/cv/img/superresolution", imageUrl, file, nil, &res)
	return res.MediaID, err
}

// PrintedTextOCR 通用印刷体识别
func (s *Ocr) PrintedTextOCR(imageUrl string, file *multipart.FileHeader) (*PrintedTextRes, error) {
	var (
		res PrintedTextRes
	)
	err := s.ocrCli("/cv/ocr/comm", imageUrl, file, nil, &res)
	return &res, err
}

// VehicleLicenseOCR 行驶证识别
func (s *Ocr) VehicleLicenseOCR(model RecognizeModel, imageUrl string, file *multipart.FileHeader) (*VehicleLicenseRes, error) {
	var (
		res VehicleLicenseRes
	)
	err := s.ocrCli("/cv/ocr/driving", imageUrl, file, map[string]string{
		"type": model,
	}, &res)
	return &res, err
}

// BankCardOCR 银行卡识别
func (s *Ocr) BankCardOCR(model RecognizeModel, imageUrl string, file *multipart.FileHeader) (string, error) {
	var (
		res BankCardRes
	)
	err := s.ocrCli("/cv/ocr/bankcard", imageUrl, file, map[string]string{
		"type": model,
	}, &res)
	return res.Number, err
}

// BusinessLicenseOCR 营业执照识别
func (s *Ocr) BusinessLicenseOCR(imageUrl string, file *multipart.FileHeader) (*BusinessLicenseRes, error) {
	var (
		res BusinessLicenseRes
	)
	err := s.ocrCli("/cv/ocr/bizlicense", imageUrl, file, nil, &res)
	return &res, err
}

// DriverLicenseOCR 驾驶证识别
func (s *Ocr) DriverLicenseOCR(imageUrl string, file *multipart.FileHeader) (*DrivingLicenseRes, error) {
	var (
		res DrivingLicenseRes
	)
	err := s.ocrCli("/cv/ocr/drivinglicense", imageUrl, file, nil, &res)
	return &res, err
}

// IdCardOCR 身份证识别
func (s *Ocr) IdCardOCR(model RecognizeModel, imageUrl string, file *multipart.FileHeader) (*IdCardRes, error) {
	var (
		res IdCardRes
	)
	err := s.ocrCli("/cv/ocr/idcard", imageUrl, file, map[string]string{
		"type": model,
	}, &res)
	return &res, err
}

// ocr请求
func (s *Ocr) ocrCli(path, imageUrl string, file *multipart.FileHeader, query map[string]string, res interface{}) error {
	var (
		err error
	)
	if imageUrl != "" {
		_, err = s.request.JsonPost(path, query, map[string]string{
			"img_url": imageUrl,
		}, &res)
	} else if file != nil {
		_, err = s.request.FormPost(path, query, map[string]*multipart.FileHeader{
			"img": file,
		}, nil, &res)
	} else {
		return core.NewError(500, "请传入图片url或文件")
	}
	return err
}
