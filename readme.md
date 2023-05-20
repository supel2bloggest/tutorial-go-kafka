# GOKAFKA Tutorial

ทดลอง tutorial kafka ตาม codebangkok

## [Reference](https://www.youtube.com/watch?v=RjtIdUOpH04)

### คำสั่งเบื้องต้น Basic Command

สร้าง topic

`kafka-topics --bootstrap-server=localhost:9092 --topic=bondhello --create`

สร้าง consumer (listener) subscribe topic

`kafka-console-consumer --bootstrap-server=localhost:9092 --topic=bondhello`

สร้้าง producer (sender) run service เพื่อ input message to topic

`kafka-console-producer --bootstrap-server=localhost:9092 --topic=bondhello`
