[기능]

1. gitlab WEBHOOK 연동하여 channel에 정보 알림.

- 유저가 스윗봇을 통해서 repository에 대한 스윗 채널을 선택할 수 있음
    - !setting repository-url channel-id
    - 설정하지 않을 경우 [Gitlab] 이라는 채널에 기본적으로 알림가도록
    - 어차피 gitlab 단에서 웹훅을 컨트롤하는 거기 때문에 api서버는 정보를 받으면 채널에 정보만 넘기면 된다 (심플)

2. 연차 명령어

- 연차 등록, 연차 연동 

3. 아지트 연동? 