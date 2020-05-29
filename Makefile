build:
	./make

run: build
	./start

stop:
	./stopper

restart: stop build run

clean:
	./cleaner

migrate:
	sql-migrate up
	sql-migrate up -config=dbconfig-campaignership.yml
	sql-migrate up -config=dbconfig-katresnan.yml
