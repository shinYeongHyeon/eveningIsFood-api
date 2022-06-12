# Evening Is Food ! (API)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/shinYeongHyeon/eveningIsFood-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/shinYeongHyeon/eveningIsFood-api)](https://goreportcard.com/report/github.com/shinYeongHyeon/eveningIsFood-api)
[![ReleaseBadge](https://img.shields.io/github/v/release/shinYeongHyeon/eveningIsFood-api)](https://pkg.go.dev/github.com/shinYeongHyeon/eveningIsFood-api)

This API is used eveningIsFood-web

## Run
```markdown
$ docker-compose --f docker-compose-db.yml up {{-d}}   
$ docker-compose up {{--force-recreate}} {{-d}}
```

### Option description
> -f: 기본 docker-compose 파일 말고 실행하기 위함  
> -d: 백그라운드에서 실행  
> --force-recreate : 새 이미지 생성해서 실행

## Documentation
[documentation](localhost:9999/docs/)

### Documentation UP-TO-DATE
```shell
$ swag init
```

## Test
```shell
$ TESTING=true go test ./...
```

---

#### Appendix
- [Commit rules](./commitRules.md)