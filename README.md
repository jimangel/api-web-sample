# Build the application

```
VERSION=0.11.2
OS=Darwin     # or Linux
ARCH=x86_64  # or arm64, i386, s390x
curl -OL https://github.com/google/ko/releases/download/v${VERSION}/ko_${VERSION}_${OS}_${ARCH}.tar.gz
tar xzf ko_${VERSION}_${OS}_${ARCH}.tar.gz ko
chmod +x ./ko
sudo mv ko /usr/local/bin/
rm -rf ko_${VERSION}_${OS}_${ARCH}.tar.gz


export PROJECT_ID=<PROJECT_ID>
export KO_DOCKER_REPO="gcr.io/$PROJECT_ID/api-demo"

cd ./api && ko build --platform=all --bare .

cd ../web

export KO_DOCKER_REPO="gcr.io/$PROJECT_ID/web-demo"

ko build --platform=all --bare .
```

To make GCR images publically available:

https://gist.github.com/jimangel/6c46f20b8f156d45d2a66175d8bbe9ab
