Необходимо создать утилиту командной строки, которая конвертирует одну валюту в другую используя
https://coinmarketcap.com/api/v1/#section/Introduction как источник данных.

Пример:
$> ./app 123.45 USD BTC

Программа должна вывести результат в консоль.

### Run
```bash
make build
./app 123.45 USD BTC
```