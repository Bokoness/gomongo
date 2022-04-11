if [ $1 = "test" ]; then
	go test -v ./...
elif [ $1 = "run" ]; then
	gowatch
else 
	echo "command not found"
fi
