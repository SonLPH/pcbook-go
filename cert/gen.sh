# 1. Generate CA's private key and self-signed certificate
rm *.pem
openssl req -x509 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=VN/ST=HCM/L=HCM/O=FPT/OU=FPT/CN=*.fpt.com/emailAddress=iwinter6963@gmail.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text
# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -keyout server-key.pem -out server-req.pem -subj "/C=VN/ST=HCM/L=HCM/O=HOME/OU=HOME/CN=*.fpt.com/emailAddress=lphoangson1708@gmail.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
