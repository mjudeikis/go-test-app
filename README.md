# GO Test App for random tests


Deploy in OpenShift:
```
oc new-app --name client quay.io/mangirdas/go-test-app
oc new-app --name server quay.io/mangirdas/go-test-app
```

Or `oc create -f deploy.yaml`
