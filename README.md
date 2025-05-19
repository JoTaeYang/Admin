## 게임 도메인의 전체적인 시스템 구성 프로젝트
운영툴부터 만들자라는 생각으로 git project 이름이 Admin으로 시작.
게임 개발을 할 때, 필요한 전체적인 서버를 구성하는 게 목표.

환경 : Go, MySQL, Redis

구현 예정
1. 로그 => Kafka 를 세팅 후 ElasticSearch에 데이터를 쌓음
2. Prometheus + Grafana로 모니터
3. Kubernetes로 관리
4. model + repository => generic 하게 만들기
5. model의 Record 기록하기[origin, change]

왜 ORM 안 썼는가?
1. 생 쿼리 써봐야 ORM이 얼마나 좋은지 알지 않겠는가?
2. 나중에 쿼리 튜닝을 해야할 경우가 생겼을 때, ORM으로는 어렵다고 들었다.



## 고민사항
DB에서 데이터를 읽어오는 Selector안에 Repository가 있다.
Repository는 Updater도 쓴다.

외부에서 주입하는 것을 어느 정도 유지하면서,
두 Util Struct가 같은 Repository를 공유했으면 좋겠다.

어떻게 구현할 것인가?(https://www.notion.so/Repository-struct-1f8c159c51ad80b69d15cd650d5bb921)


