package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	kAyarla()

	hedeflerim := hyukle("targets.yaml")
	if len(hedeflerim) == 0 {
		return
	}

	fmt.Println("--- Dark Web Görsel Tarama Başlatıldı ---")

	allocCtx, cancel := tarayıcıAyar()
	defer cancel()

	sonuclar := make(map[string]string)

	for _, url := range hedeflerim {
		durum := SiteTaraa(allocCtx, url)
		sonuclar[url] = durum
	}
	Raporla(sonuclar)
}

func kAyarla() {
	os.MkdirAll("outputs", 0755)
	os.MkdirAll("logs", 0755)
}

func hyukle(dosyaYolu string) []string {
	content, err := os.ReadFile(dosyaYolu)
	if err != nil {
		fmt.Println("[HATA] Hedef dosyası okunamadı:", err)
		return nil
	}

	var tListe []string
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			tListe = append(tListe, line)
		}
	}
	return tListe
}

func tarayıcıAyar() (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ProxyServer("socks5://127.0.0.1:9150"),
		chromedp.Flag("ignore-certificate-errors", true),
	)
	return chromedp.NewExecAllocator(context.Background(), opts...)
}

func SiteTaraa(ctx context.Context, url string) string {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	fmt.Printf("[INFO] İşleniyor: %s\n", url)

	taskCtx, cancel := chromedp.NewContext(ctx)
	taskCtx, cancel = context.WithTimeout(taskCtx, 90*time.Second)
	defer cancel()

	var resimVerisi []byte
	var htmlIcerigi string

	err := chromedp.Run(taskCtx,
		chromedp.Navigate(url),
		chromedp.Sleep(5*time.Second),
		chromedp.FullScreenshot(&resimVerisi, 90),
		chromedp.OuterHTML("html", &htmlIcerigi),
	)

	if err != nil {
		log.Printf("[ERR] %s hatası: %v", url, err)
		fmt.Printf(" :( :'(   -> Başarısız (Timeout veya Kapalı)\n")
		return "BAŞARISIZ"
	}

	baseName := strings.NewReplacer("http://", "", "https://", "", ".", "_", "/", "_").Replace(url)

	os.WriteFile(filepath.Join("outputs", baseName+".png"), resimVerisi, 0644)
	os.WriteFile(filepath.Join("outputs", baseName+".html"), []byte(htmlIcerigi), 0644)

	fmt.Printf(" :) :) -> BAŞARILI: Görsel ve HTML kaydedildi.\n")
	return "BAŞARILI"
}

func Raporla(sonuclar map[string]string) {
	f, err := os.Create("tarama_raporu.txt")
	if err != nil {
		fmt.Println("Rapor oluşturulamadı:", err)
		return
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("--- TARAMA RAPORU (%s) ---\n\n", time.Now().Format("2006-01-02 15:04:05")))

	for url, durum := range sonuclar {
		f.WriteString(fmt.Sprintf("[%s] %s\n", durum, url))
	}
	fmt.Println("\n[+] Rapor 'tarama_raporu.txt' dosyasına kaydedildi.")
}
