IMAGE := quay.io/mangirdas/go-test-app

clean:
	rm -rf http

http: clean
	go build ./cmd/$@

image:
	podman build . -f images/Dockerfile -t ${IMAGE}

run:
	podman run -p:8080:8080 ${IMAGE}

push:
	podman push ${IMAGE}
