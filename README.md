🕵️ Dark Web Scraper

Bu proje, dark web üzerindeki belirli siteleri otomatik olarak ziyaret eden ve içeriklerini analiz amacıyla kaydeden bir web scraping aracıdır.
Araç, verilen .onion adreslerine bağlanarak sayfanın ekran görüntüsünü (screenshot) alır ve HTML içeriğini indirerek yerel olarak saklar.

Bu sayede araştırmacılar ve siber güvenlik uzmanları, dark web sitelerini offline olarak inceleyebilir ve analiz edebilir.

🚀 Özellikler

🧅 .onion adreslerini otomatik ziyaret etme

📸 Web sayfasının ekran görüntüsünü alma

📄 HTML içeriğini indirme ve kaydetme

🔁 Birden fazla siteyi otomatik tarama

💾 Analiz için verileri yerel olarak saklama

🧰 Kullanılan Teknolojiler

Python

Tor Network

Requests / Selenium / Playwright (kullandığına göre değiştirebilirsin)

HTML Parsing

Headless Browser Automation

⚙️ Kurulum

Projeyi çalıştırmadan önce Tor servisinin aktif olması gerekir.

1️⃣ Depoyu klonla
git clone https://github.com/kullaniciadi/darkweb-scraper.git
cd darkweb-scraper
2️⃣ Gerekli paketleri yükle
pip install -r requirements.txt
3️⃣ Tor servisini başlat
service tor start
4️⃣ Scripti çalıştır
python scraper.py
📂 Çıktılar

Script çalıştırıldığında aşağıdaki veriler oluşturulur:

output/
 ├── screenshots/
 │    └── site1.png
 │
 ├── html/
 │    └── site1.html
 │
 └── logs/
      └── scan.log

screenshots/ → Sayfaların ekran görüntüleri

html/ → İndirilen HTML içerikleri

logs/ → Tarama kayıtları

⚠️ Uyarı

Bu araç yalnızca eğitim, araştırma ve siber güvenlik analizleri için geliştirilmiştir.
Dark web üzerindeki içerikleri tararken yerel yasalara ve etik kurallara uyulması kullanıcının sorumluluğundadır.
