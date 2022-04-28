if [ $1 = "test" ]; then
	go test -v ./...
elif [ $1 = "dev" ]; then
	gowatch
else 
	echo "command not found"
fi
