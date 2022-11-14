- GOROOT: Go를 설치한 위치를 지정하는 환경변수. 기본 경로는 /usr/local/go
- GOPATH: Go 프로젝트를 실행하거나 개발할 때 사용할 워킹 디렉토리를 가르키는 환경 변수. 기본 경로는 /Users/{username}/go
    - 하단에는 각각 bin, pkg, src라는 폴더가 생김
        - bin: Go 프로젝트의 실행 파일
        - pkg: 컴파일된 패키지의 오브젝트
        - src: Go 프로젝트의 소스 코드
  
- 실행 명령어는 go run {sourceName}
- 패키지 설치 명령어는 go install {packageUrl}
    - 이 때 설치된 패키지의 실행 파일이 bin 폴더에 생김

- Go는 main 파일을 기반으로 프로젝트가 실행됨 main으로 지정할 go 파일 상단에 <code>package main</code>을 적음으로 해당 파일을 main 파일로 지정할 수 있음
