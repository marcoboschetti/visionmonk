curl 'http://localhost:8080/api/p/shop_product' \
  -H 'Connection: keep-alive' \
  -H 'sec-ch-ua: " Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36' \
  -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' \
  -H 'Accept: */*' \
  -H 'x-auth-token: c403ad1c-6b03-481c-bb82-84cc1cd6e300' \
  -H 'X-Requested-With: XMLHttpRequest' \
  -H 'sec-ch-ua-platform: "Linux"' \
  -H 'Origin: http://localhost:8080' \
  -H 'Sec-Fetch-Site: same-origin' \
  -H 'Sec-Fetch-Mode: cors' \
  -H 'Sec-Fetch-Dest: empty' \
  -H 'Referer: http://localhost:8080/' \
  -H 'Accept-Language: es,en;q=0.9,en-US;q=0.8' \
  -H 'Cookie: _ga=GA1.1.1360784377.1613625444' \
  --data-raw '{"sku":"RJ1323","inventory":123,"price_cts":12300,"catalog_product":{"title":"AA","description":"BB","barcode":"GG","image_base_64":"","category":"CC","brand":"DD","color":"EE","size":"FF"}}' 