package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	//Kullanıcıdan url bekleme kısmı

	urlalici := flag.String("url", "", "Hedeflenen web sitesinin url adresini giriniz:")
	flag.Parse()

	//Url girilmediyse kapansın

	if *urlalici == "" {
		fmt.Println("Lütfen bir url adresi giriniz :(")
		os.Exit(2)
	}

	hedefurl := *urlalici

	//Dosya isimlendirme işlemleri

	temizisim := strings.ReplaceAll(hedefurl, "https://", "")
	temizisim = strings.ReplaceAll(temizisim, "http://", "")
	temizisim = strings.ReplaceAll(temizisim, "www.", "")
	temizisim = strings.ReplaceAll(temizisim, "/", "_")
	temizisim = strings.ReplaceAll(temizisim, ":", "")

	//Tarayıcı işlemleri

	//Context oluşturma

	ctx, iptal := chromedp.NewContext(context.Background())
	defer iptal()

	//30 sn içinde bir cevap alınmassa network hatası vb durumları için

	ctx, iptal = context.WithTimeout(ctx, 30*time.Second)
	defer iptal()

	var durumkodu int64
	var html_icerigi string
	var ekran_goruntusu []byte
	var linkler []string

	//Tarayıcıdan gelen cevapları dinlemek için

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		// Gelen olay bir "Cevap Alındı" (ResponseReceived) olayı mı?
		if ev, ok := ev.(*network.EventResponseReceived); ok {
			// Bu cevap bizim ana sitemizden mi geliyor?
			if ev.Response.URL == hedefurl || ev.Response.URL == hedefurl+"/" {
				durumkodu = ev.Response.Status // Kodu yakaladık!
			}
		}
	})

	//Tarayıcının hedef adrese gitmesi ve istenen işlemleri gerçekleştirmesi

	fmt.Printf("Hedef: %s üzerinde işlem yapılıyor...\n", hedefurl)
	hata := chromedp.Run(ctx,
		chromedp.Navigate(hedefurl),
		chromedp.Sleep(2*time.Second),
		chromedp.OuterHTML("html", &html_icerigi),
		chromedp.FullScreenshot(&ekran_goruntusu, 100),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(a => a.href)`, &linkler),
	)

	//Durum kodlarının yansıtılması

	if hata != nil {
		log.Fatalf("Hata: Sİteye hiçbir şekilde erişilemedi: %v", hata)
	}

	if durumkodu == 404 {
		fmt.Printf("Hata: Sayfa bulunamadı (HTTP 404)\n")
		return
	} else if durumkodu != 200 {
		fmt.Printf("Site farklı bir durum kodu dönderdi: %d\n", durumkodu)
	} else {
		fmt.Printf("Siteye erişildi (HTTP 200).\n")
	}

	//İstediğim 200 kodunu alırsam işlemlere devam edeceğim

	if durumkodu == 200 {

		//Çıktıları koyabiliceğim bir klasör oluşturma

		ciktiklasoru := "outputs"

		if hata := os.MkdirAll(ciktiklasoru, 0755); hata != nil {
			log.Fatalf("Klasör oluşturulamadı: %v", hata)
		}

		//Dosya isimlerini oluşturma

		htmldosyaadi := fmt.Sprintf("html_%s.html", temizisim)
		pngdosyaadi := fmt.Sprintf("ss_%s.png", temizisim)
		linklerdosyaadi := fmt.Sprintf("linkler_%s.txt", temizisim)

		//Dosya yollarını belirginleştirme

		htmldosyayolu := filepath.Join(ciktiklasoru, htmldosyaadi)
		pngdosyayolu := filepath.Join(ciktiklasoru, pngdosyaadi)
		linklerdosyayolu := filepath.Join(ciktiklasoru, linklerdosyaadi)

		//Html içeriğini dosyaya yaz

		if hata := os.WriteFile(htmldosyayolu, []byte(html_icerigi), 0644); hata != nil {
			log.Printf("Html kaydedilirken hata oluştu : %v\n", hata)
		} else {
			fmt.Printf("Html başarıyla kaydedildi: %v\n", htmldosyayolu)
		}

		//Ekran görüntüsünü dosyaya kaydet

		if hata := os.WriteFile(pngdosyayolu, ekran_goruntusu, 0644); hata != nil {
			log.Printf("Ekran görüntüsü kaydedilirken hata oluştu: %v\n", hata)
		} else {
			fmt.Printf("Ekran görüntüsü başarıyla kaydedildi: %v\n", pngdosyayolu)
		}

		//Linkleri dosyaya yaz

		link_listesi := strings.Join(linkler, "\n")
		if hata := os.WriteFile(linklerdosyayolu, []byte(link_listesi), 0644); hata != nil {
			log.Printf("Linkler kaydedilemedi: %v\n", hata)
		} else {
			fmt.Printf("Linkler kaydedildi (%v adet): %v\n", len(linkler), linklerdosyayolu)
		}
	}

}
