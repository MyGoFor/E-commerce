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

.PHONY: consul
consul:
	@open "http://localhost:8500/ui/"
