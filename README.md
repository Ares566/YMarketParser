# YMarketParser
Парсер Яндекс Маркета

- [x] Парсинг страницы товара: категория, название, метки, рейтинг, количество отзывов, цены, дисконт
- [x] Реформатирование кода
- [x] Парсинг корзинки на предмет остатков на складе
- [ ] Добавить проверки на "существование" элементов на странице

### Запуск парсера
Пока "топорно": указываем URL в main.go, затем
```
$ go run ./cmd/app/main.go

```
### Пример парсинга товара с остатками на складе
![Пример парсинга товара](https://downloader.disk.yandex.ru/preview/584de1cb4b8ccc1d15672b172a891204e1857252905a5743fe46bb9a267e795f/60d5adbf/ZLCWNzh5m5rW5fbi4_0Cnc9xv-pRX-FbeEGu0nwwWyIL6GamYXedoENs0XghAx6-UWC3X_5i6LmepCDDV_3bVg%3D%3D?uid=0&filename=2021-06-25_09-19-18.png&disposition=inline&hash=&limit=0&content_type=image%2Fpng&owner_uid=0&tknv=v2&size=2048x2048)
