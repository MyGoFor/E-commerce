export MOD=github.com/MyGoFor/E-commerce

.PHONY: cwgo_hertz_home
cwgo_hertz_home:
	@cd app/frontend && \
	cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module ${MOD}/app/frontend -I ../../idl && \
	cd ../..

.PHONY: cwgo_hertz_auth_page
cwgo_hertz_auth_page:
	@cd app/frontend && \
	cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module ${MOD}/app/frontend -I ../../idl && \
	cd ../..

.PHONY: cwgo_hertz_product_page
cwgo_hertz_product_page:
	@cd app/frontend && \
	cwgo server --type HTTP --idl ../../idl/frontend/product_page.proto --service frontend -module ${MOD}/app/frontend -I ../../idl && \
	cd ../..

.PHONY: cwgo_hertz_category_page
cwgo_hertz_category_page:
	@cd app/frontend && \
	cwgo server --type HTTP --idl ../../idl/frontend/category_page.proto --service frontend -module ${MOD}/app/frontend -I ../../idl && \
	cd ../..

.PHONY: cwgo_hertz_cart_page
cwgo_hertz_cart_page:
	@cd app/frontend && \
	cwgo server --type HTTP --idl ../../idl/frontend/cart_page.proto --service frontend -module ${MOD}/app/frontend -I ../../idl && \
	cd ../..

.PHONY: cwgo_hertz_checkout_page
cwgo_hertz_checkout_page:
	@cd app/frontend && \
	cwgo server --type HTTP --idl ../../idl/frontend/checkout_page.proto --service frontend -module ${MOD}/app/frontend -I ../../idl && \
	cd ../..

.PHONY: cwgo_hertz_order_page
cwgo_hertz_order_page:
	@cd app/frontend && \
	cwgo server --type HTTP --idl ../../idl/frontend/order_page.proto --service frontend -module ${MOD}/app/frontend -I ../../idl && \
	cd ../..

.PHONY: docker
docker:
	@sudo docker compose up -d

.PHONY: cwgo_kitex_client_user
cwgo_kitex_client_user:
	@cd rpc_gen && \
	cwgo client --type RPC --idl ../idl/user.proto --service user -module ${MOD}/rpc_gen -I ../idl && \
	cd ..

.PHONY: cwgo_kitex_server_user
cwgo_kitex_server_user:
	@cd app/user && \
	cwgo server --type RPC --idl ../../idl/user.proto --service user --pass "-use ${MOD}/rpc_gen" -module ${MOD}/app/user -I ../../idl && \
	cd ../..

.PHONY: cwgo_kitex_client_product
cwgo_kitex_client_product:
	@cd rpc_gen && \
	cwgo client --type RPC --idl ../idl/product.proto --service product -module ${MOD}/rpc_gen -I ../idl && \
	cd ..

.PHONY: cwgo_kitex_server_product
cwgo_kitex_server_product:
	@cd app/product && \
	cwgo server --type RPC --idl ../../idl/product.proto --service product --pass "-use ${MOD}/rpc_gen" -module ${MOD}/app/product -I ../../idl && \
	cd ../..

.PHONY: cwgo_kitex_client_cart
cwgo_kitex_client_cart:
	@cd rpc_gen && \
	cwgo client --type RPC --idl ../idl/cart.proto --service cart -module ${MOD}/rpc_gen -I ../idl && \
	cd ..

.PHONY: cwgo_kitex_server_cart
cwgo_kitex_server_cart:
	@cd app/cart && \
	cwgo server --type RPC --idl ../../idl/cart.proto --service cart --pass "-use ${MOD}/rpc_gen" -module ${MOD}/app/cart -I ../../idl && \
	cd ../..

.PHONY: cwgo_kitex_client_order
cwgo_kitex_client_order:
	@cd rpc_gen && \
	cwgo client --type RPC --idl ../idl/order.proto --service order -module ${MOD}/rpc_gen -I ../idl && \
	cd ..

.PHONY: cwgo_kitex_server_order
cwgo_kitex_server_order:
	@cd app/order && \
	cwgo server --type RPC --idl ../../idl/order.proto --service order --pass "-use ${MOD}/rpc_gen" -module ${MOD}/app/order -I ../../idl && \
	cd ../..

.PHONY: cwgo_kitex_client_payment
cwgo_kitex_client_payment:
	@cd rpc_gen && \
	cwgo client --type RPC --idl ../idl/payment.proto --service payment -module ${MOD}/rpc_gen -I ../idl && \
	cd ..

.PHONY: cwgo_kitex_server_payment
cwgo_kitex_server_payment:
	@cd app/payment && \
	cwgo server --type RPC --idl ../../idl/payment.proto --service payment --pass "-use ${MOD}/rpc_gen" -module ${MOD}/app/payment -I ../../idl && \
	cd ../..

.PHONY: cwgo_kitex_client_checkout
cwgo_kitex_client_checkout:
	@cd rpc_gen && \
	cwgo client --type RPC --idl ../idl/checkout.proto --service checkout -module ${MOD}/rpc_gen -I ../idl && \
	cd ..

.PHONY: cwgo_kitex_server_checkout
cwgo_kitex_server_checkout:
	@cd app/checkout && \
	cwgo server --type RPC --idl ../../idl/checkout.proto --service checkout --pass "-use ${MOD}/rpc_gen" -module ${MOD}/app/checkout -I ../../idl && \
	cd ../..

.PHONY: cwgo_kitex_client_email
cwgo_kitex_client_email:
	@cd rpc_gen && \
	cwgo client --type RPC --idl ../idl/email.proto --service email -module ${MOD}/rpc_gen -I ../idl && \
	cd ..

.PHONY: cwgo_kitex_server_email
cwgo_kitex_server_email:
	@cd app/email && \
	cwgo server --type RPC --idl ../../idl/email.proto --service email --pass "-use ${MOD}/rpc_gen" -module ${MOD}/app/email -I ../../idl && \
	cd ../..

.PHONY: consul
consul:
	@open "http://localhost:8500/ui/"

.PHONY: user_run
user_run:
	@cd app/user && go run . && cd ../.. \

.PHONY: product_run
product_run:
	@cd app/product && go run . && cd ../.. \

.PHONY: cart_run
cart_run:
	@cd app/cart && go run . && cd ../.. \

.PHONY: order_run
order_run:
	@cd app/order && go run . && cd ../.. \

.PHONY: payment_run
payment_run:
	@cd app/payment && go run . && cd ../.. \

.PHONY: checkout_run
checkout_run:
	@cd app/checkout && go run . && cd ../.. \

.PHONY: email_run
email_run:
	@cd app/email && go run . && cd ../.. \

.PHONY: frontend_run
frontend_run:
	@cd app/frontend && go run . && cd ../.. \

.PHONY: grafana
grafana:
	@open "http://localhost:3000"

.PHONY: jaeger
jaeger:
	@open "http://localhost:16686/search"

# make build-frontend v=v1.1.1
.PHONY: build-frontend
build-frontend:
	sudo docker build -f ./deploy/Dockerfile.frontend -t frontend:${v} .

# make build-svc v=v1.1.1 svc=product
.PHONY: build-svc
build-svc:
	sudo docker build -f ./deploy/Dockerfile.svc -t ${svc}:${v} --build-arg SVC=${svc} .

.PHONY: build-all
build-all:
	sudo docker build -f ./deploy/Dockerfile.frontend -t frontend:${v} .
	sudo docker build -f ./deploy/Dockerfile.svc -t cart:${v} --build-arg SVC=cart .
	sudo docker build -f ./deploy/Dockerfile.svc -t checkout:${v} --build-arg SVC=checkout .
	sudo docker build -f ./deploy/Dockerfile.svc -t email:${v} --build-arg SVC=email .
	sudo docker build -f ./deploy/Dockerfile.svc -t order:${v} --build-arg SVC=order .
	sudo docker build -f ./deploy/Dockerfile.svc -t payment:${v} --build-arg SVC=payment .
	sudo docker build -f ./deploy/Dockerfile.svc -t product:${v} --build-arg SVC=product .
	sudo docker build -f ./deploy/Dockerfile.svc -t user:${v} --build-arg SVC=user .

.PHONY: dockerfile-frontend
dockerfile-frontend:
	@sudo docker run -v ./app/frontend/conf:/opt/E-commerce/frontend/conf --network e-commerce_default --env-file ./app/frontend/.env -p 8080:8080 7fd2ba529bb5

.PHONY: dockerfile-product
dockerfile-product:
	@sudo docker run -v ./app/product/conf:/opt/E-commerce/product/conf --network e-commerce_default --env-file ./app/product/.env 40dbbb3deba7