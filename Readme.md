# Gerando Certificado Raiz

```sh
openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out certs/cert.pem \
  -keyout certs/key.pem \
  -addext "subjectAltName = DNS:localhost" \
  -subj "/C=BR/ST=Bahia/L=Salvador/O=MyOrg/OU=Your Unit/CN=localhost"
```
