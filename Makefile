
install:
	go install
	scp ~/go/bin/msgqcat ${DAVE}@bridge2:bin/msgqcat
