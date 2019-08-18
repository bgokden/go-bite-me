# go-bite-me
Simple Byte Array to Byte Array Map implementation in Golang

It is a very simple implementation.

I frequently need a temporary key value store 
where I need to push data then iterate over it only once then delete whole data.

Unfortunately, go slices are not hashable so I couldn't use the default map implementation.
I don't want get/delete so I don't care about key retrievel performance.
