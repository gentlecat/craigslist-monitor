fmt-js :
	cd frontend && yarn run prettier

fmt-go :
	go fmt -x ./...

webpack : fmt-js
	cd frontend && yarn run bundle

test-go : fmt-go
	go test ./... -bench .

test-js : fmt-js
	cd frontend && yarn run lint

rebuild : fmt-go
	go build -a -o ./build/monitor ./cmd/monitor/monitor.go
	go build -a -o ./build/browser ./cmd/browser/browser.go

build-monitor : fmt-go
	go build -o ./build/monitor ./cmd/monitor/monitor.go

build-browser : fmt-go
	go build -o ./build/browser ./cmd/browser/browser.go

run : build-browser build-monitor test-js
	# Using `concurrently` from npm to run both front-end builds and back-end at
	# the same time.
	cd frontend && \
	npx concurrently "../build/browser" "../build/monitor -u 'https://vancouver.craigslist.org/search/apa?availabilityMode=0&bundleDuplicates=1&hasPic=1&housing_type=1&housing_type=2&housing_type=4&housing_type=5&housing_type=8&laundry=1&maxSqft=1600&max_bedrooms=1&max_price=2200&minSqft=600&min_price=1600&parking=1&parking=2&postal=V6E0A8&search_distance=10'" "yarn run bundle" \
		--names browser,monitor,webpack, \
		--prefix-colors green.bold,blue.bold,red.bold

test-run-monitor : build-monitor
	./build/monitor -u "https://vancouver.craigslist.org/search/apa?availabilityMode=0&bundleDuplicates=1&hasPic=1&housing_type=1&housing_type=2&housing_type=4&housing_type=5&housing_type=8&laundry=1&maxSqft=1600&max_bedrooms=1&max_price=2200&minSqft=600&min_price=1600&parking=1&parking=2&postal=V6E0A8&search_distance=10"

run-browser : build-browser
	./build/browser
