# ISG Yönetim Sistemi API

Bu proje, İş Sağlığı ve Güvenliği (ISG) yönetim sistemi için geliştirilmiş bir RESTful API'dir. Go dilinde yazılmıştır ve Gin web framework'ü kullanılmaktadır.

## Özellikler

- Kullanıcı yönetimi (kayıt, giriş, profil)
- Proje yönetimi
- Personel yönetimi
- ISG eğitim takibi
- Sağlık raporu takibi
- JWT tabanlı kimlik doğrulama
- Detaylı loglama sistemi

## Teknolojiler

- Go 1.21+
- Gin Web Framework
- GORM ORM
- PostgreSQL
- JWT Authentication

## Başlangıç

### Gereksinimler

- Go 1.21 veya üzeri
- PostgreSQL
- Git

### Kurulum

1. Projeyi klonlayın:
```bash
git clone https://github.com/yourusername/isg-api.git
cd isg-api
```

2. Bağımlılıkları yükleyin:
```bash
go mod download
```

3. `.env` dosyasını oluşturun:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=isg_db
JWT_SECRET=your_jwt_secret
JWT_EXPIRATION=24h
```

4. Veritabanını oluşturun:
```bash
createdb isg_db
```

5. Uygulamayı başlatın:
```bash
go run cmd/main.go
```

## API Endpoints

### Kullanıcı İşlemleri

- `POST /api/register` - Yeni kullanıcı kaydı
- `POST /api/login` - Kullanıcı girişi
- `GET /api/user/:id` - Kullanıcı profili

### Proje İşlemleri

- `POST /api/project` - Yeni proje oluşturma
- `GET /api/projects` - Tüm projeleri listeleme
- `GET /api/project/:id` - Proje detayları
- `PUT /api/project/:id` - Proje güncelleme
- `DELETE /api/project/:id` - Proje silme
- `GET /api/project/user/:user_id` - Kullanıcının projeleri

### Personel İşlemleri

- `POST /api/personnel` - Yeni personel ekleme
- `PUT /api/personnel/:id` - Personel güncelleme
- `DELETE /api/personnel/:id` - Personel silme

### ISG Eğitim İşlemleri

- `POST /api/isg` - Yeni ISG eğitimi ekleme
- `GET /api/isg` - Tüm ISG eğitimlerini listeleme
- `PUT /api/isg/:id` - ISG eğitimi güncelleme
- `DELETE /api/isg/:id` - ISG eğitimi silme

### Sağlık Raporu İşlemleri

- `POST /api/saglik-raporu` - Yeni sağlık raporu ekleme
- `PUT /api/saglik-raporu/:id` - Sağlık raporu güncelleme
- `DELETE /api/saglik-raporu/:id` - Sağlık raporu silme
- `GET /api/saglik-raporu/personel/:personel_id` - Personelin sağlık raporları

## Veritabanı Yapısı

### Users
- ID (Primary Key)
- Name
- Email (Unique)
- Password
- Role
- CreatedAt

### Projects
- ID (Primary Key)
- UserID (Foreign Key)
- Description
- CreatedAt
- UpdatedAt

### Personel
- ID (Primary Key)
- ProjectID (Foreign Key)
- NameSurname
- TcNo (Unique)
- PhoneNo
- Profession
- BloodType
- IsActive
- Description
- CreatedAt
- UpdatedAt
- DeletedAt

### Isg_Egitim
- ID (Primary Key)
- PersonelID (Foreign Key)
- BaslangicTarihi
- BitisTarihi

### SaglikRaporu
- ID (Primary Key)
- PersonelID (Foreign Key)
- BaslangicTarihi
- BitisTarihi

## Güvenlik

- Tüm endpoint'ler JWT token ile korunmaktadır (login ve register hariç)
- Şifreler bcrypt ile hashlenmektedir
- API anahtarları ve hassas bilgiler .env dosyasında saklanmaktadır

## Loglama

- Tüm HTTP istekleri loglanır
- Hata ve bilgi mesajları ayrı dosyalarda tutulur
- Log dosyaları günlük olarak oluşturulur
- Log formatı: `[SEVİYE] TARİH SAAT DOSYA:Satır | MESAJ`

## Hata Yönetimi

- Tüm hatalar uygun HTTP durum kodları ile döndürülür
- Hata mesajları Türkçe olarak verilir
- Hatalar loglanır ve izlenebilir

## Geliştirme

### Kod Standartları

- Go fmt kullanılır
- Açıklayıcı değişken ve fonksiyon isimleri
- Türkçe hata mesajları
- İngilizce kod yorumları

### Test

```bash
go test ./...
```



## İletişim

Sorularınız veya önerileriniz için:
- Email: gokberkkozak@gmail.com
- GitHub Issues: [https://github.com/zegasega/isg-api/issues ](https://github.com/zegasega/Nano-Isg-API/issues)
