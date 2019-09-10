#!/bin/bash
cd "$GOPATH/src/domain.com"

echo "==== Release domain.com ===="
echo "  Deleting the local binary if it exists..."
rm domain.com
echo "  Done!"

echo "  Deleting existing code..."
ssh root@sub.domain.com "rm -rf /root/go/src/domain.com"
echo "  Code deleted successfully!"

echo "  Uploading code..."
rsync -avr --exclude '.git/*' --exclude 'tmp/*' --exclude 'images/*' ./
root@sub.domain.com:/root/go/src/domain.com/
echo "  Code uploaded successfully!"

echo "  Go getting deps..."
ssh root@sub.domain.com "export GOPATH=/root/go; /usr/local/go/bin/go get golang.org/x/crypto/bcrypt"
ssh root@sub.domain.com "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/gorilla/mux"
ssh root@sub.domain.com "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/gorilla/schema"
ssh root@sub.domain.com "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/lib/pq"
ssh root@sub.domain.com "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/jinzhu/gorm"
ssh root@sub.domain.com "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/gorilla/csrf"

echo "  Building the code on remote server..."
ssh root@sub.domain.com 'export GOPATH=/root/go; cd /root/app;
/usr/local/go/bin/go build -o ./server $GOPATH/src/domain.com/*.go'
echo "  Code built successfully!"

echo "  Moving assets..."
ssh root@sub.domain.com "cd /root/app; cp -R /root/go/src/domain.com/assets ."
echo "  Assets moved successfully!"

echo "  Moving views..."
ssh root@sub.domain.com "cd /root/app; cp -R /root/go/src/domain.com/views ."
echo "  Views moved successfully!"

echo "  Moving Caddyfile..."
ssh root@sub.domain.com "cd /root/app; cp /root/go/src/domain.com/Caddyfile ."
echo "  Caddyfile moved successfully!"

echo "  Restarting the server..."
ssh root@sub.domain.com "sudo service domain.com restart"
echo "  Server restarted successfully!"

echo "  Restarting Caddy server..."
ssh root@sub.domain.com "sudo service caddy restart"
echo "  Caddy restarted successfully!"

echo "==== Done releasing domain.com ===="
