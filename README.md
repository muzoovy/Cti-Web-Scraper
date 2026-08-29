```text
# Go CTI Web Scraper

**Go tabanlı otomatik keşif ve bilgi toplama aracı.**

Bu proje, hedef web sitelerinden manuel etkileşime gerek kalmadan HTML, ekran görüntüsü ve bağlantı haritasını çıkarmak amacıyla **Golang** dilini ve **Chromedp** teknolojisi kullanılarak geliştirilmiştir.

##  Temel Özellikler

* **Headless Tarayıcı Otomasyonu:** Modern ve JavaScript ağırlıklı siteleri arka planda sorunsuz işler.
* **Ekran görüntüsü alma:** Hedef sitenin anlık durumunu yüksek çözünürlüklü `.png` formatında kaydeder.
* **Kaynak Kod Analizi:** Sitenin DOM yapısını tarayarak HTML kaynak kodunu indirir.
* **Bağlantı haritasını çıkarma:** Sayfadaki tüm dış ve iç bağlantıları (`href`) ayrıştırarak listeler.
* **Durum kodu yönetimi:** Ağ hatalarını ve 404 (Sayfa Bulunamadı) gibi durumları tespit eder, boş yere veri çekmeye çalışmaz.
* **Otomatik Dosyalama:** URL içindeki geçersiz karakterleri temizler ve tüm çıktıları `outputs/` klasöründe düzenli bir şekilde saklar.

## Kurulum

### Gereksinimler
* Bilgisayarınızda [Go (Golang)](https://go.dev/dl/) kurulu olmalıdır.
* Google Chrome veya Chromium tarayıcısı yüklü olmalıdır.

```

### Projeyi İndirme

```bash
git clone [https://github.com/Muzoovy4606/Cti-Web-Scraper.git](https://github.com/Muzoovy4606/Cti-Web-Scraper.git)
cd Cti-Web-Scraper

```



### Bağımlılıkları Yükleme

```bash
go mod tidy

```

## Kullanım

Aracı çalıştırmak için `-url` parametresi ile hedef siteyi belirtmeniz yeterlidir.

```bash
go run cti-web-scraper.go -url [https://hedef-site.com]

```

### Derleme (Opsiyonel)

Aracı taşınabilir bir `.exe` dosyasına dönüştürmek isterseniz:

```bash
go build -o scraper
./scraper -url [https://hedef-site.com]

```

## 📂 Çıktı Yapısı

Program çalıştığında ana dizinde `outputs/` adında bir klasör oluşturur ve tüm verileri buraya kaydeder.

```text
cti-web-scraper/
├── cti-web-scraper.go
├── outputs/
│   ├── html_hedef_site_com.html      # Kaynak Kod Dosyası
│   ├── ss_hedef_site_com.png         # Ekran Görüntüsü
│   └── linkler_hedef_site_com.txt    # Link Listesi
└── README.md

```

<p align="center">
SiberVatan CTI çalışması kapsamında <a href="https://www.google.com/search?q=https://github.com/Muzoovy4606">Muzoovy</a> tarafından geliştirilmiştir.
</p>

```

```
