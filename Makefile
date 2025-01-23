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

.PHONY: consul
consul:
	@open "http://localhost:8500/ui/"

.PHONY: run_user
run_user:
	@cd app/user && go run . && cd ../.. \

.PHONY: run_product
run_product:
	cd app/product && go run . && cd ../.. \

.PHONY: run_cart
run_cart:
	cd app/cart && go run . && cd ../.. \

.PHONY: run_order
run_order:
	cd app/order && go run . && cd ../.. \

.PHONY: run_payment
run_payment:
	cd app/payment && go run . && cd ../.. \

.PHONY: run_checkout
run_checkout:
	cd app/checkout && go run . && cd ../.. \

.PHONY: run_frontend
run_frontend:
	cd app/frontend && go run . && cd ../.. \
