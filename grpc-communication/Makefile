
.PHONY: clean certs all
all: clean

clean:
	-rm -f server.grpc.key
	-rm -f server.grpc.crt

# source: https://github.com/denji/golang-tls
certs:
# Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
	openssl ecparam -genkey -name secp384r1 -out server.grpc.key
# Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
	openssl req -new -x509 -sha256 \
	-key server.grpc.key -out server.grpc.crt -days 3650 \
	-subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=server.grpc"
