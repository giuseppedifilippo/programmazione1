package main

import (
	"math"
	"strconv"
	"strings"
)

type Wifi struct {
	ssid                           string
	channel, frequency, signal_dBm int
}

func NewWifi(ssid string, channel, frequency, signal_dBm int) (WiFi, bool) {
	var wifi Wifi
	if len(ssid) == 0 && signal_dBm < -20 {
		return Wifi{}, false
	}
	wifi.ssid = ssid
	wifi.signal_dBm = signal_dBm
	f1, f2 := frequency >= 2412 && frequency <= 2484, frequency >= 5035 && frequency <= 5980
	if (f1 && channel >= 1 && channel <= 14) || (f2 && channel >= 7 && channel <= 196) {
		wifi.channel, wifi.frequency = channel, frequency
		return wifi, true
	}
	return Wifi{}, false
}

func NewWifiDaStringa(line string) (Wifi, bool) {
	sl := strings.Split(line, ",")
	temp := []int{}
	for _, el := range sl[1:] {
		x, _ := strconv.Atoi(el)
		temp = append(temp, x)
	}
	return NewWifi(sl[0], temp[0], temp[1], temp[2])
}

func ConvertiDBinWatt(signal_dBm int) float64 {
	return math.Pow10(signal_dBm)/1000.0
}

func 