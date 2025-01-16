.PHONY: cwgo_hertz_home
cwgo_hertz_home:
	@cd app/frontend && \
	cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/MyGoFor/E-commerce/app/frontend -I ../../idl && \
	cd ../..