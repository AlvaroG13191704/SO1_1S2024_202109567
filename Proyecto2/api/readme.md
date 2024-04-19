## Como hacer deploy a Artifact Registry

Build
```bash
docker buildx build --platform linux/amd64 -t api:0.1 .
```

Tag
```bash
docker tag api:0.1 us-central1-docker.pkg.dev/sopes1-417522/so1p2/api:0.1
```

Push
```bash
docker push us-central1-docker.pkg.dev/sopes1-417522/so1p2/api:0.1
```