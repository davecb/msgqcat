
install:
	go install
	scp ~/go/bin/msgqcat ${TARGET}
