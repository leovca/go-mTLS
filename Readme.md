# Gerando Certificado Raiz

```sh
openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out cert.pem \
  -keyout key.pem \
  -subj "/C=BR/ST=Bahia/L=Salvador/O=MyOrg/OU=Your Unit/CN=localhost"
```
