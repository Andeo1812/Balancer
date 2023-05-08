# Balancer

Данный проект создан для ознакомления с работой балансировщиков

## Окружение

* Nginx выступает в роли балансировщика
* Сервис-монолит на Go
* Метрики собираются Prometheus
* Мониторинг на Grafana
* [Фреймворк]((https://github.com/wg/wrk)) нагрузочного тестирования
* Docker, docker-compose

## Сценарий

1. Поднимаем Nginx, Prometheus, Grafana
2. Запускаем сервис на 3 "подах"
3. Запускаем нагрузочное тестирование с одной частотой
4. Смотрим на дашбордах распределение загрузки под действием нагрузки
5. "Тушим" 1 "под"
6. Смотрим новое распределение нагрузки
7. Включаем снова под
8. Смотрим новое распределение нагрузки

## Команды для запуска

### Запуск кластера сервиса:
Важно предварительно установить переменную POD = UUID "пода"  

```bash
make run-cluster-app POD=N
```

### Запуск мониторинга:
Важно установить права доступа для директорий которые монтирует docker-compose для grafana

```bash
make run-monitoring
```

### Запуск nginx:

```bash
make run-nginx
```

### Запуск нагрузки:
Запуск атакера на nginx в 4 треда и 100 конекшенов в течении 10 минут

```bash
make bench
```

### Остановки всей инфры машины:

```bash
make stop
```

## Результат

![image](https://user-images.githubusercontent.com/88785411/236829951-2ce84096-3043-4e2f-b399-56d86988838c.png)

* Запуск всех подов с нагрузкой 
* Отключение 3 пода
* Включение 3 пода 