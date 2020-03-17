module github.com/cuttle-ai/data-integration-service

go 1.13

replace github.com/cuttle-ai/auth-service => ../auth-service/

replace github.com/cuttle-ai/brain => ../brain/

replace github.com/cuttle-ai/octopus => ../octopus/

replace github.com/cuttle-ai/configs => ../configs/

replace github.com/cuttle-ai/db-toolkit => ../db-toolkit/

require (
	github.com/cuttle-ai/auth-service v0.0.0-00010101000000-000000000000
	github.com/cuttle-ai/configs v0.0.0-20190824112953-7860fdfd0dae
	github.com/cuttle-ai/db-toolkit v0.0.0-00010101000000-000000000000 // indirect
	github.com/cuttle-ai/file-uploader-service v0.0.0-20200307152319-90e778ef6f2a // indirect
	github.com/jinzhu/gorm v1.9.12
)
