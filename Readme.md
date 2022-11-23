# Gerando Certificado Raiz

```sh
openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out certs/ca.crt \
  -keyout certs/ca.key \
  -addext "subjectAltName = DNS:localhost" \
  -subj "/C=BR/ST=Bahia/L=Salvador/O=MyOrg/OU=Your Unit/CN=localhost"
```

# Gerando Certificado para o Servidor

```sh
openssl genrsa \
  -out certs/server.key 2048
```

```sh
openssl req \
  -new \
  -key certs/server.key \
  -subj '/CN=server' \
  -addext "subjectAltName = DNS:localhost" \
  -out certs/server.csr
```

```sh
openssl x509 \
  -req \
  -in certs/server.csr \
  -CA certs/ca.crt \
  -CAkey certs/ca.key \
  -CAcreateserial \
  -days 365 \
  -out certs/server.crt \
  -extfile <(echo subjectAltName = DNS:localhost)
```

# Gerando Certificado para o cliente

```sh
openssl genrsa \
  -out certs/client.key 2048
```

```sh
openssl req \
  -new \
  -key certs/client.key \
  -subj '/CN=client' \
  -addext "subjectAltName = DNS:localhost" \
    -out certs/client.csr
```

```sh
openssl x509 \
  -req \
  -in certs/client.csr \
  -CA certs/ca.crt \
  -CAkey certs/ca.key \
  -CAcreateserial \
  -days 365 \
  -out certs/client.crt \
  -extfile <(echo subjectAltName = DNS:localhost)
```
