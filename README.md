WB L0

localhost:8186/ - форма для поиска заказа.

Результаты нагрузочного тестирования сервиса:
Инструмент тестирования vegeta (срипт scripts/loadtest/test.go)
Роут /order тестировался методом post с передачей в теле uid необходимого заказа
Тестирование проводилось на одной с сервисом машине.

rps: 1000 / 15сек
Время мин / сред / макс (µs - микросек) : 1.6213ms 3.189008ms 92.0096ms
Status Codes: map[200:15000]
Кол-во запросов: 15000

rps: 2000 / 30сек
Время мин / сред / макс (µs - микросек) : 998µs 9.285614ms 365.6712ms
Status Codes: map[200:59999]
Кол-во запросов: 59999     

rps: 3000 / 30сек
Время мин / сред / макс (µs - микросек) : 1.6955ms 1.134462122s 2.8787689s
Status Codes: map[200:89997]
Кол-во запросов: 89997      

rps: 4000 / 30сек
Время мин / сред / макс (µs - микросек) : 3.1424372s 3.556777628s 10.0731895s
Status Codes: map[0:34883 200:85115]
Кол-во запросов: 119998

В случае, если заменить парсинг, формирование, отдачу html файла и 
выдавать данные в json - скорость значительно вырастет:

rps: 4000 / 30сек
Время мин / сред / макс (µs - микросек) : 531.6µs 1.56961ms 58.583ms
Status Codes: map[200:119997]
Кол-во запросов: 119997  