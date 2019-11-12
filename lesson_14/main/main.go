package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

// sensor関数型
type sensor func() kelvin

func main() {
	list_14_6()
}

func list_14_6() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
}

func calibrate(s sensor, offset kelvin) sensor {
	// sensor型と同じ関数シグネチャの無名関数を返却する
	return func() kelvin {
		return s() + offset
	}
}

func list_14_5() {
	func() {
		fmt.Println("Functions anonymous")
	}() // 宣言してすぐ呼び出し
}

func list_14_4() {
	// 引数付き無名関数
	f := func(message string) {
		fmt.Println(message)
	}

	f("Go to the party.")
}

func list_14_3() {
	// 無名関数
	var f = func() {
		fmt.Println("Dress up for the masquerade.")
	}

	f()
}

func list_14_2() {
	measureTemperature(3, fakeSensor)
}

// 第2引数は関数
func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vo K\n", k)
		time.Sleep(time.Second)
	}
}

func list_14_1() {
	// fakeSensorとrealSensorが同じシグネチャだから実現出来る

	// 関数を代入
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 1
}
