

export PATH := ./bin:$(PATH)

./bin/bats:
	git clone https://github.com/sstephenson/bats.git ./tmp/bats
	./tmp/bats/install.sh .

acceptance-test:
	@bats --tap acceptance.bats