PROJECT_NAME = "university-ds-a"

#Local settings
BINARY_NAME = ${PROJECT_NAME}
BINARIES = "./bin"
MAIN_DIR = "cmd/${BINARY_NAME}"

#GitHub Info
GIT_LOCAL_NAME = "rodziievskyi-maksym"
GIT_LOCAL_EMAIL = "rodziyevskydev@gmail.com"
GITHUB = "github.com/${GIT_LOCAL_NAME}/${PROJECT_NAME}"

init:
	@echo "::> Creating a module root..."
	@go mod init ${GITHUB}
	@mkdir "cmd" && mkdir "cmd/"${BINARY_NAME}
	@touch ${MAIN_DIR}/main.go
	@echo "package main\n\nimport \"fmt\"\n\nfunc main(){\n\tfmt.Println(\"${BINARY_NAME}\")\n}" > ${MAIN_DIR}/main.go
	@touch VERSION && echo 0.0.1 > VERSION
	@git add ${MAIN_DIR}/main.go go.mod VERSION
	@echo "::> Finished!"

build:
	@echo "::> Building..."
	@go build -o ${BINARIES}/${BINARY_NAME} ${MAIN_DIR}
	@echo "::> Finished!"

run:
	@go build -o ${BINARIES}/${BINARY_NAME} ${MAIN_DIR}
	@${BINARIES}/${BINARY_NAME}

clean:
	@echo "::> Cleaning..."
	@go clean
	@rm -rf ${BINARIES}
	@go mod tidy
	@echo "::> Finished"

local-git:
	@git config --local user.name ${GIT_LOCAL_NAME}
	@git config --local user.email ${GIT_LOCAL_EMAIL}
	@git config --local --list

git-init:
	@echo "::> Git initialization begin..."
	@git init
	@git config --local user.name ${GIT_LOCAL_NAME}
	@git config --local user.email ${GIT_LOCAL_EMAIL}
	@touch .gitignore
	@echo ".idea" > .gitignore
	@echo "bin" > .gitignore
	@touch README.md
	@git add README.md
	@git commit -m "first commit"
	@git branch -M main
	@git remote add origin https://${GITHUB}
	@git push -u origin main
	@echo "::> Finished"

test:
	go test -v -cover ./...

.PNONY: init build run clean local-git git-init postgres create-db drop-db migrate-up migrate-down sqlc test mock migrate-down-last migrate-up-last