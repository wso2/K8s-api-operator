# Working with HTTP/HTTPS Private Registry

You can use a HTTP, insecure registry or HTTPS, secure registry to push the built micro-gateway.

## Deploy a Registry Server

1.  Create a password file with one entry for the user `TEST_USER`, with password `TEST_PASSWORD`
    ```sh
    >> mkdir auth
    >> docker run \
      --entrypoint htpasswd \
      registry:2 -Bbn TEST_USER TEST_PASSWORD > auth/htpasswd
    ```
1. Add following to your hosts file (Linux/Mac: `/etc/hosts`, Windows: `c:\windows\system32\drivers\etc\hosts`).
    ```sh
    <MY_IP_ADDRESS>       myregistrydomain.com
    ```

### Deploy a HTTP registry

1. Start the registry with basic authentication.
    ```sh
    >> docker run -d \
      -p 5001:5000 \
      --restart=always \
      --name registry-http \
      -v "$(pwd)"/auth:/auth \
      -e "REGISTRY_AUTH=htpasswd" \
      -e "REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm" \
      -e REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd \
      registry:2
    ```
   
1. Add insecure registry.
   Edit `docker/daemon.json` file `<HOME_DIR>/.docker/daemon.json` to add insecure registry.
   ```json
   { 
      "debug":true,
      "experimental":false,
      "insecure-registries":[ 
         "myregistrydomain.com:5001"
      ]
   }
   ```
   If you are using `Minikube` start it with the flag `--insecure-registry="myregistrydomain.com:5001"`.

### Deploy a HTTPS registry

1. Create self-signed certificates
    ```sh
    >> mkdir -p certs
    >> openssl req \
      -newkey rsa:4096 -nodes -sha256 -keyout certs/domain.key \
      -x509 -days 365 -out certs/domain.crt
    ```
   Be sure to use the name `myregistrydomain.com` as a CN.

1. Start the registry with basic authentication.
    ```sh
    >> docker run -d \
      -p 5002:5000 \
      --restart=always \
      --name registry-auth \
      -v "$(pwd)"/auth:/auth \
      -e "REGISTRY_AUTH=htpasswd" \
      -e "REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm" \
      -e REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd \
      -v "$(pwd)"/certs:/certs \
      -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/domain.crt \
      -e REGISTRY_HTTP_TLS_KEY=/certs/domain.key \
      registry:2
    ```
1. Instruct every Docker daemon to trust that certificate. The way to do this depends on your OS.
   
   - Linux: Copy the domain.crt file to `/etc/docker/certs.d/myregistrydomain.com:5002/ca.crt` on every Docker host. You do not need to restart Docker.
   
   - Windows Server:
       1. Open Windows Explorer, right-click the domain.crt file, and choose Install certificate. When prompted, select the following options:
          - Store location: local machine
          - Place all certificates in the following store: selected
       1. Click Browser and select Trusted Root Certificate Authorities.
       
       1. Click Finish. Restart Docker.
   
   - Docker Desktop for Mac: Follow the instructions on [Adding custom CA certificates](https://docs.docker.com/docker-for-mac/faqs/#how-do-i-add-custom-ca-certificates). Restart Docker.
   
   - Docker Desktop for Windows: Follow the instructions on [Adding custom CA certificates](https://docs.docker.com/docker-for-windows/faqs/#how-do-i-add-custom-ca-certificates). Restart Docker.

Follow the documentation https://docs.docker.com/registry/deploying for more information on deploying a HTTP/HTTPS registry.

## Install API Operator

- Execute the following command to install API Operator interactively and configure repository to push the microgateway image.
- Select registry type.
  - Select "HTTP Private Registry" as the repository type for HTTP registry.
  - Select "Docker Hub" as the repository type for HTTPS registry.
- Enter the file path of the downloaded service account key JSON File.
- Confirm configuration are correct with entering "Y"

### HTTP registry example
```sh
>> apictl install api-operator
Choose registry type:
1: Docker Hub (Or others, quay.io, HTTPS registry)
2: Amazon ECR
3: GCR
4: HTTP Private Registry
Choose a number: 1: 4
Enter private registry (10.100.5.225:5000/jennifer): myregistrydomain.com:5001/my-repository
Enter username: TEST_USER
Enter password:

Repository: myregistrydomain.com:5001/my-repository
Username  : TEST_USER
Confirm configurations: Y:
```

Output:
```sh
[Installing OLM]
customresourcedefinition.apiextensions.k8s.io/clusterserviceversions.operators.coreos.com created
...

[Installing API Operator]
subscription.operators.coreos.com/my-api-operator created
[Setting configs]
namespace/wso2-system created
...

[Setting to K8s Mode]
```

### HTTPS registry example
```sh
>> apictl install api-operator
Choose registry type:
1: Docker Hub (Or others, quay.io, HTTPS registry)
2: Amazon ECR
3: GCR
4: HTTP Private Registry
Choose a number: 1: 1
Enter repository name (docker.io/john | quay.io/mark | 10.100.5.225:5000/jennifer): myregistrydomain.com:5002/my-repository
Enter username: TEST_USER
Enter password:

Repository: myregistrydomain.com:5002/my-repository
Username  : TEST_USER
Confirm configurations: Y:
```

Output:
```sh
[Installing OLM]
customresourcedefinition.apiextensions.k8s.io/clusterserviceversions.operators.coreos.com created
...

[Installing API Operator]
subscription.operators.coreos.com/my-api-operator created
[Setting configs]
namespace/wso2-system created
...

[Setting to K8s Mode]
```