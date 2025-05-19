package glog

// CREATE_UID : 최초 UID 생성
type CreateUid struct {
	BaseLog
	Country   string `json:"country"`    // 2자리 국가 코드
	AuthType  uint8  `json:"auth_type"`  // 2 : 구글 3 : 애플
	OSType    uint8  `json:"os_type"`    // 2 : 구글 3 : 애플
	OSVer     string `json:"os_ver"`     // 디바이스 OS 버전
	ClientVer string `json:"client_ver"` // 클라이언트 버전
	ClientIP  string `json:"client_ip"`  // 접속 주소
}
