# API

- https://github.com/cloud-hypervisor/cloud-hypervisor/blob/main/docs/api.md#rest-api

```
$ brew install openapi-generator

$ openapi-generator generate -i https://raw.githubusercontent.com/cloud-hypervisor/cloud-hypervisor/master/vmm/src/api/openapi/cloud-hypervisor.yaml -g go -o /tmp/test/
```
