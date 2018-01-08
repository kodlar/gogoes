package main

import (
	"fmt"
	"math"
	"strconv"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func (r *Rectangle) calculateRectangleArea() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func FloatToString(inputnum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputnum, 'f', 6, 64)
}

// Circle2 oluştur
type Circle2 struct {
	x, y, r float64
}

// Circle oluştur
type Circle struct {
	x float64
	y float64
	r float64
}

// Rectangle oluşturur
type Rectangle struct {
	x1, y1, x2, y2 float64
}

// Playlist oluşturur
type Playlist struct {
	id      float64
	name    string
	preview string
	source  string
}

// CreatePlaylistJson should be playlistjson
func (p *Playlist) CreatePlaylistJson() string {
	return FloatToString(p.id) + p.name + p.preview + p.source
}

func circleArea2(c Circle) float64 {
	fmt.Println(c.r)
	return math.Pi * c.r * c.r
}

func circleArea3(c *Circle) float64 {
	fmt.Println(c.r)
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {

	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.calculateRectangleArea())

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))

	//var cx, cy, cr float64 = 0, 0, 5
	//fmt.Println(circleArea(cx, cy, cr))

	//Birinci nesne  yaratımı tüm değerler 0 olarak atanır
	//var c Circle
	//ikinci nensne yaratımı burada bu struct için memoryde yer ayrılır. Tüm değerleri yine sıfır olarak set edilirken bize pointer döndürür.
	//t := new(Circle2)
	//Eğer bu yöntemi yapmazsak circle için değer girerek nesne oluşturabilriiz
	//c := Circle(x:0, y:0, r:5);
	//ikinci tanımlama yöntemi ise girilecek metod fonk sırasını biliyorsanız parametre isimlerini vermeden girilebilir
	//c := Circle(0, 0, 5)
	//eğer structa işaretçi ile erişmek istiyorsak & operatörü kullanılır
	//c:= &Circle{0,0,5}
	//struct içindeki alanlara erişmek istersek . operatörünü kullanıyoruz
	//fm.Println(c.x, c.y, c.r)
	//c.x = 10
	//c.y = 5
	//go'da fonksiyon parametreleri kopyalandığı için  fonksiyon içine giren parametreler fonksiyon içinde değiştirilemiyor.
	//Bu yüzden bir değişiklik yapılacaksa nesne yaratıldıktan sonra ilgili alanlar güncellenmeli, bunun anlaşılır kılmak için & operatörünü
	//ekleyerek fonksiyon içinde gelen parametrelere değişiklik yapılamadığını belirtiyoruz.
	c := Circle{0, 0, 3}
	fmt.Println(circleArea3(&c))
	d := Circle{0, 0, 2}
	fmt.Println(circleArea2(d))
	//Eğer fonksiyon ismi ile parametrenin methodda yerini değiştirirsek özel bir yazılış tipi elde ediyoruz. Bu fonksiyon isminden önce
	//gelen fonksiyon parametreleri önde olacak şekilde yazılınca nesneden sonra . operatörünü kullanarak sanki o nesneye ait bir fonksiyonmuş
	//gibi bir özellik kazanıyor. Böylece okuması daha kolay bir yazım şekli elde ediyoruz.
	fmt.Println(c.area())

	playlist := Playlist{1, "ali baba okula gider", "http://d.marmaragazetesi.com/gallery/461_1.jpg", "http://techslides.com/demos/sample-videos/small.webm"}
	fmt.Println(playlist.CreatePlaylistJson())
}
