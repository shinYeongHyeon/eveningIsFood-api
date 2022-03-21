# Evening Is Food ! (API)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/shinYeongHyeon/eveningIsFood-api)

## 커밋 규칙
```markdown
{type}: ({이슈번호}) {title}

{body} (생략가능)

See also : {이슈번호} (참고 이슈, 생략가능)
```

### For example,
```markdown
FEAT: (EIF-T-1234) 유저 이름 변경

유저 이름변경 Application 생성

See also: (EIF-I-1233)
```

## 커밋 TYPE 유형
```markdown
FEAT: 새로운 기능의 추가
FIX: 버그 수정
DOCS: 문서 수정
STYLE: 스타일 관련 기능(코드 포맷팅, 세미콜론 누락, 코드 자체의 변경이 없는 경우)
REFACTOR: 코드 리펙토링
TEST: 테스트 코트, 리펙토링 테스트 코드 추가
CHORE: 빌드 업무 수정, 패키지 매니저 수정(ex .gitignore 수정 같은 경우)
```

---

## 서버 실행
```markdown
$ docker-compose -f docker-compose-db.yml up {{-d}}
$ docker-compose up {{-d}}
```

### 옵션 설명
> -f: 기본 docker-compose 파일 말고 실행하기 위함
> -d: 백그라운드에서 실행

### Formatting
```shell
gofmt -w src/*
```

### Test
```shell
go test ./...
```
