package ocr

type RecognizeModel = string

const (
	RecognizeModePhoto RecognizeModel = "photo"
	RecognizeModelScan RecognizeModel = "scan"
)

// AICropRes 图片智能裁剪后的返回数据
type AICropRes struct {
	Results []struct {
		CropLeft   uint `json:"crop_left"`
		CropTop    uint `json:"crop_top"`
		CropRight  uint `json:"crop_right"`
		CropBottom uint `json:"crop_bottom"`
	} `json:"results"`
	IMGSize struct {
		Width  uint `json:"w"`
		Height uint `json:"h"`
	} `json:"img_size"`
}

// QRCodePoint 二维码角的位置
type QRCodePoint struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}

// ScanQRCodeRes 小程序的条码/二维码识别后的返回数据
type ScanQRCodeRes struct {
	CodeResults []struct {
		TypeName string `json:"type_name"`
		Data     string `json:"data"`
		Position struct {
			LeftTop     QRCodePoint `json:"left_top"`
			RightTop    QRCodePoint `json:"right_top"`
			RightBottom QRCodePoint `json:"right_bottom"`
			LeftBottom  QRCodePoint `json:"left_bottom"`
		} `json:"pos"`
	} `json:"code_results"`
	IMGSize struct {
		Width  uint `json:"w"`
		Height uint `json:"h"`
	} `json:"img_size"`
}

// SuperResolutionRes 图片高清化后的返回数据
type SuperResolutionRes struct {
	MediaID string `json:"media_id"`
}

// DataPosition 数据区域
type DataPosition struct {
	LeftTop     DataPoint `json:"left_top"`
	RightTop    DataPoint `json:"right_top"`
	RightBottom DataPoint `json:"right_bottom"`
	LeftBottom  DataPoint `json:"left_bottom"`
}

// DataPoint 数据点位
type DataPoint struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}

// PrintedTextRes 通用印刷体 OCR 识别返回数据
type PrintedTextRes struct {
	Items []struct {
		Text     string       `json:"text"`
		Position DataPosition `json:"pos"`
	} `json:"items"` //	识别结果
	ImgSize DataPoint `json:"img_size"` //	图片大小
}

// VehicleLicenseRes 识别卡片返回数据
type VehicleLicenseRes struct {
	VehicleType                string `json:"vehicle_type"`
	Owner                      string `json:"owner"`
	Addr                       string `json:"addr"`
	UseCharacter               string `json:"use_character"`
	Model                      string `json:"model"`
	Vin                        string `json:"vin"`
	EngineNum                  string `json:"engine_num"`
	RegisterDate               string `json:"register_date"`
	IssueDate                  string `json:"issue_date"`
	PlateNumB                  string `json:"plate_num_b"`
	Record                     string `json:"record"`
	PassengersNum              string `json:"passengers_num"`
	TotalQuality               string `json:"total_quality"`
	TotalPrepareQualityQuality string `json:"totalprepare_quality_quality"`
}

// BankCardRes 识别银行卡返回数据
type BankCardRes struct {
	Number string `json:"number"` // 银行卡号
}

// BusinessLicenseRes 营业执照 OCR 识别返回数据
type BusinessLicenseRes struct {
	RegNum              string `json:"reg_num"`              //	注册号
	Serial              string `json:"serial"`               //	编号
	LegalRepresentative string `json:"legal_representative"` //	法定代表人姓名
	EnterpriseName      string `json:"enterprise_name"`      //	企业名称
	TypeOfOrganization  string `json:"type_of_organization"` //	组成形式
	Address             string `json:"address"`              //	经营场所/企业住所
	TypeOfEnterprise    string `json:"type_of_enterprise"`   //	公司类型
	BusinessScope       string `json:"business_scope"`       //	经营范围
	RegisteredCapital   string `json:"registered_capital"`   //	注册资本
	PaidInCapital       string `json:"paid_in_capital"`      //	实收资本
	ValidPeriod         string `json:"valid_period"`         //	营业期限
	RegisteredDate      string `json:"registered_date"`      //	注册日期/成立日期
	CertPosition        struct {
		Position DataPosition `json:"pos"`
	} `json:"cert_position"` //	营业执照位置
	ImgSize DataPoint `json:"img_size"` //	图片大小
}

// DrivingLicenseRes 识别行驶证返回数据
type DrivingLicenseRes struct {
	IDNum        string `json:"id_num"`        // 证号
	Name         string `json:"name"`          // 姓名
	Nationality  string `json:"nationality"`   // 国家
	Sex          string `json:"sex"`           // 性别
	Address      string `json:"address"`       // 地址
	BirthDate    string `json:"birth_date"`    // 出生日期
	IssueDate    string `json:"issue_date"`    // 初次领证日期
	CarClass     string `json:"car_class"`     // 准驾车型
	ValidFrom    string `json:"valid_from"`    // 有效期限起始日
	ValidTo      string `json:"valid_to"`      // 有效期限终止日
	OfficialSeal string `json:"official_seal"` // 印章文构
}

type IdCardRes struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Id          string `json:"id"`
	Addr        string `json:"addr"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	ValidDate   string `json:"valid_date"`
}
