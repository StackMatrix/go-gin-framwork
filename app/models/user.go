package models

import "strconv"

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Account  string `json:"account" gorm:"not null;unqiue;comment:用户账号"`
	Password string `json:"-" gorm:"not null;default:'';comment:用户密码"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}

type Log struct {
	ID
	Account        string `json:"account" gorm:"not null;comment:用户账号"`
	DeviceImei     string `json:device_imei" gorm:"not null;comment:设备IMEI"`
	DeviceName     string `json:device_name" gorm:"not null;comment:设备名称"`
	DeviceModel    string `json:device_model" gorm:"not null;comment:设备模型"`
	DeviceLocation string `json:device_location" gorm:"not null;comment:设备位置"`
	DeviceLat      string `json:device_lat" gorm:"not null;comment:设备经度"`
	DeviceLon      string `json:device_lon" gorm:"not null;comment:设备纬度"`
	DevicePlatform string `json:device_platform" gorm:"not null;comment:设备平台"`
	Timestamps
	SoftDeletes
}
