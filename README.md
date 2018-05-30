# WebServer with terraform

This is nothing more than the basic example copied from the terratest blog post, I wrote it to run this test by myself.


# Getting to run this

Install terratest (should have Go installed)

```
go get github.com/gruntwork-io/terratest/modules/terraform
```

Run the tests

```
cd tests
go test -v 
```