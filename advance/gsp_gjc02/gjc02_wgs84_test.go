package test

import (
	"testing"

	"github.com/ForrestSu/go_learn/advance/gsp_gjc02/transform"
)

// https://www.zhihu.com/question/26112618
type Point struct {
	Lat float64
	Lng float64
}

func (p *Point) ToWGS84() Point {
	var wgs84 Point
	wgs84.Lat, wgs84.Lng = transform.GCJtoWGS(p.Lat, p.Lng)
	return wgs84
}

func (p *Point) Distance(target Point) float64 {
	return transform.Distance(p.Lat, p.Lng, target.Lat, target.Lng)
}

// 分别使用 gjc 和 wgs84 的经纬度, 计算2个点的距离
// 结果误差：在千分之一米左右
func TestDist(t *testing.T) {
	// gjc02(home <--> office)
	var gjcHome = Point{Lat: 22.618873, Lng: 114.144797}
	var gjcOffice = Point{Lat: 22.547379828559027, Lng: 113.94241373697916}
	d := gjcHome.Distance(gjcOffice)
	t.Logf("gjc02 距离: %f\n", d)

	// wgs84 (home <--> office)
	var wgsHome = gjcHome.ToWGS84()
	var wgsOffice = gjcOffice.ToWGS84()
	t.Logf("wgsOffice = %+v", wgsOffice)
	d2 := wgsHome.Distance(wgsOffice)
	t.Logf("wgs84 距离: %f\n", d2)

	t.Logf("误差: %fm, rate %f%%\n", d2-d, 100*(d2-d)/d) // 误差在 0.1% 以内
}

// 快15倍
// BenchmarkPlain-12        8130693               132.2 ns/op
// BenchmarkExt-12           604148              1953 ns/op
func BenchmarkPlain(b *testing.B) {
	var gjcKeXing = Point{Lat: 22.547379828559027, Lng: 113.94241373697916}
	for i := 0; i < b.N; i++ {
		transform.GCJtoWGS(gjcKeXing.Lat, gjcKeXing.Lng)
	}
}

func BenchmarkExt(b *testing.B) {
	var gjcKeXing = Point{Lat: 22.547379828559027, Lng: 113.94241373697916}
	for i := 0; i < b.N; i++ {
		transform.GCJtoWGSExact(gjcKeXing.Lat, gjcKeXing.Lng)
	}
}
