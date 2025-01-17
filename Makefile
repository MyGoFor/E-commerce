.PHONY: cwgo_hertz_home
cwgo_hertz_home:
	@cd app/frontend && \
	cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/MyGoFor/E-commerce/app/frontend -I ../../idl && \
	cd ../..

.PHONY: cwgo_hertz_auth_page
cwgo_hertz_auth_page:
	@cd app/frontend && \
	cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module github.com/MyGoFor/E-commerce/app/frontend -I ../../idl && \
	cd ../..

.PHONY: docker
docker:
	@sudo docker compose up -d