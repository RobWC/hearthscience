GOBINDATA="go-bindata"

bindata:
	$(GOBINDATA) -pkg="hearthscience" templates/

clean:
	-rm -f bindata.go
